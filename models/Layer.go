package models

import (
    "gopkg.in/gographics/imagick.v3/imagick"
    "net/http"
    "io/ioutil"
)

type Layer struct {
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
}

const DistortPolynomial string = "polynomial"
const NumbCoordinatesPoint int = 4

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

    mw.CompositeImage(overlay, imagick.COMPOSITE_OP_OVER, false, int(layer.Left), int(layer.Top))

    return mw, err
}

func DistortLayer(channel chan PositionMagicWand, errors chan error, layer Layer) {
    mw := imagick.NewMagickWand()
    pmw := PositionMagicWand{Position: layer.Position}

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

    // Пересчитать матрицу
    if layer.DistortionType == DistortPolynomial {
        layer.RecalculateMatrix()
    }

    // Само искажение, самая долгая операция
    if len(layer.DistortionMatrix) != 0 {
        // TODO правильный тип искажения, на основе запроса
        mw.DistortImage(imagick.DISTORTION_POLYNOMIAL, layer.DistortionMatrix, false)
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

    // Отдаём в канал структуры с позицией и зображением
    channel <- pmw
}

func (l *Layer) RecalculateMatrix() {

    if l.DistortionOrder == 0 {
        numbPoints := len(l.DistortionMatrix) / NumbCoordinatesPoint

        if l.NumbPointsSide == 0 || l.NumbPointsSide == 2 {
            l.DistortionOrder = 1.5
        } else if l.NumbPointsSide == 3 && numbPoints <= 15 {
            l.DistortionOrder = 2
        } else if l.NumbPointsSide == 3 && numbPoints > 15 || l.NumbPointsSide == 4 {
            l.DistortionOrder = 3
        } else {
            l.DistortionOrder = 4
        }

        l.DistortionMatrix = append([]float64{l.DistortionOrder}, l.DistortionMatrix...)
    }
}