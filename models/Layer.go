package models

import (
    "gopkg.in/gographics/imagick.v3/imagick"
    "net/http"
    "io/ioutil"
)

type Layer struct {
    BackgroundColor  string    `json:"background_color"`
    DesignHeight     float64   `json:"design_height"`
    DesignLeft       float64   `json:"design_left"`
    DesignTop        float64   `json:"design_top"`
    DesignWidth      float64   `json:"design_width"`
    DistortionMatrix []float64 `json:"distortion_matrix"`
    DistortionOrder  float64   `json:"distortion_order"`
    DistortionType   string    `json:"distortion_type"`
    Height           float64   `json:"height"`
    Left             float64   `json:"left"`
    NumbPointsSide   int64     `json:"numb_points_side"`
    OverlayPath      string    `json:"overlay_path"`
    Path             string    `json:"path"`
    Position         int       `json:"position"`
    Top              float64   `json:"top"`
    Type             string    `json:"type"`
    Width            float64   `json:"width"`

    // Вообще этих полей нет в JSON схеме. Но они добавляются в процессе конвертации PostData.ConvertPositioning()
    OverlayWidth  float64   `json:"overlay_width"`
    OverlayHeight float64   `json:"overlay_height"`
    OverlayTop    float64   `json:"overlay_top"`
    OverlayLeft   float64   `json:"overlay_left"`
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

    mw.ScaleImage(uint(layer.OverlayWidth), uint(layer.OverlayHeight))
    mw.CompositeImage(overlay, imagick.COMPOSITE_OP_DST_OUT, false, int(layer.OverlayTop), int(layer.OverlayLeft))

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
        pw.SetColor(layer.BackgroundColor)
    } else {
        pw.SetColor("none")
    }

    bmw.NewImage(uint(layer.DesignWidth), uint(layer.DesignHeight), pw)

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

    // TODO вынести в отдельные методы, на основе типа слоёв
    // Изменяем размер
    mw.ResizeImage(uint(layer.DesignWidth), uint(layer.DesignHeight), imagick.FILTER_CATROM)

    // Само искажение, самая долгая операция
    if len(layer.DistortionMatrix) != 0 {
        // TODO правильный тип искажения, на основе запроса
        mw.DistortImage(imagick.DISTORTION_POLYNOMIAL, layer.DistortionMatrix, true)
    }

    // Накладываем маску на слой
    if layer.OverlayPath != "" {
        mw, err = MaskLayer(mw, layer)
        if err != nil {
            errors <- err
            return
        }
    }

    pmw.MagicWand = mw

    // Отдаём в канал структуры с позицией и изображением
    channel <- pmw
}
