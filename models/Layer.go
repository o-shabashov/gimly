package models

import (
    "gopkg.in/gographics/imagick.v2/imagick"
    "net/http"
    "io/ioutil"
    "fmt"
)

type Layer struct {
    Top              float64   `json:"top"`
    Left             float64   `json:"left"`
    Type             string    `json:"type"`
    Width            float64   `json:"width"`
    Height           float64   `json:"height"`
    Position         int       `json:"position"`
    DesignTop        float64   `json:"design_top"`
    DesignLeft       float64   `json:"design_left"`
    DesignWidth      float64   `json:"design_width"`
    DesignHeight     float64   `json:"design_height"`
    DistortionType   string    `json:"distortion_type"`
    DistortionOrder  float64   `json:"distortion_order"`
    NumbPointsSide   int64     `json:"numb_points_side"`
    DistortionMatrix []float64 `json:"distortion_matrix"`
    OverlayPath      string    `json:"overlay_path"`
    Path             string    `json:"path"`

    // Вообще этих полей нет в JSON схеме. Но они добавляются в процессе конвертации PostData.ConvertPositioning()
    OverlayWidth  float64   `json:"overlay_width"`
    OverlayHeight float64   `json:"overlay_height"`
    OverlayLeft   float64   `json:"overlay_left"`
    OverlayTop    float64   `json:"overlay_top"`

    BackgroundColor  string    `json:"background_color"`
}

const DISTORT_POLYNOMIAL string = "polynomial"
const NUMB_COORDINATES_POINT int = 4

// TODO получать оверлей картинку в отдельном треде, пока идёт искажение основного слоя
func MaskLayer(mw *imagick.MagickWand, layer Layer) (*imagick.MagickWand, error) {
    overlay := imagick.NewMagickWand()

    // Получаем оверлей по HTTP от файлменеджера
    response, err := http.Get(layer.OverlayPath)
    if err != nil {
        return mw, err
    }

    defer response.Body.Close()
    data, err := ioutil.ReadAll(response.Body)
    if err != nil {
        return mw, err
    }

    overlay.ReadImageBlob(data)
    overlay.SetImageVirtualPixelMethod(imagick.VIRTUAL_PIXEL_TRANSPARENT)

    if layer.OverlayWidth != 0 && layer.OverlayHeight != 0 {
        overlay.ScaleImage(uint(layer.OverlayWidth), uint(layer.OverlayHeight))
    }

    mw.CompositeImage(overlay, imagick.COMPOSITE_OP_DST_OUT, int(layer.OverlayLeft), int(layer.OverlayTop))

    mw.WriteImage(fmt.Sprintf("masked_%v.png", layer.Position))
    return mw, err
}

func Build(channel chan PositionMagicWand, errors chan error, layer Layer) {
    // Основное изображение, на него будем наносить все данные
    mw := imagick.NewMagickWand()
    mw.SetSize(uint(layer.DesignWidth), uint(layer.DesignHeight))

    // Бэкграунд изображение
    bmw := imagick.NewMagickWand()
    pw := imagick.NewPixelWand()

    if layer.BackgroundColor != "" {
        pw.SetColor("#" + layer.BackgroundColor)
    } else {
        pw.SetColor("none")
    }

    bmw.NewImage(uint(layer.DesignWidth), uint(layer.DesignHeight), pw)
    mw.CompositeImage(bmw, imagick.COMPOSITE_OP_OVER, int(layer.Left), int(layer.Top))

    pmw := PositionMagicWand{Layer: layer}

    // Получаем слой по HTTP от файлменеджера
    response, err := http.Get(layer.Path)
    if err != nil {
        errors <- err
        return
    }

    defer response.Body.Close()
    data, err := ioutil.ReadAll(response.Body)
    if err != nil {
        errors <- err
        return
    }

    mw.ReadImageBlob(data)
    mw.SetImageVirtualPixelMethod(imagick.VIRTUAL_PIXEL_TRANSPARENT)

    // TODO вынести в отдельные методы, на основе типа слоёв, ресайз основного слоя
    // Изменяем размер
    mw.ResizeImage(uint(layer.DesignWidth), uint(layer.DesignHeight), imagick.FILTER_CATROM, 1)

    // Само искажение, самая долгая операция
    if len(layer.DistortionMatrix) != 0 {
        // TODO правильный тип искажения, на основе запроса
        mw.DistortImage(imagick.DISTORTION_POLYNOMIAL, layer.DistortionMatrix, false)
        mw.WriteImage(fmt.Sprintf("distorted_%v.png", layer.Position))
    }

    // Накладываем маску на слой
    if layer.OverlayPath != "" {
        mw, err = MaskLayer(mw, layer)
        if err != nil {
            errors <- err
            return
        }
    }

    pmw.ImageDataBytes = mw.GetImageBlob()
    mw.Destroy()

    // Отдаём в канал структуры с позицией и изображением
    channel <- pmw
}
