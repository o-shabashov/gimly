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
    DistortionOrder  string    `json:"distortion_order"`
    DistortionType   string    `json:"distortion_type"`
    Height           float64   `json:"height"`
    Left             float64   `json:"left"`
    NumbPointsSide   int64     `json:"numb_points_side"`
    OverlayPath      string    `json:"overlay_path"`
    Path             string    `json:"path"`
    Position         int64     `json:"position"`
    Top              float64   `json:"top"`
    Type             string    `json:"type"`
    Width            float64   `json:"width"`
}

// TODO получать оверлей картинку в отдельном треде, пока идёт искажение основного слоя
func MaskLayer(mw *imagick.MagickWand, layer Layer) (*imagick.MagickWand, error) {
    overlay := imagick.NewMagickWand()

    // Получаем оверлей по HTTP от файлменеджера
    response, err := http.Get(layer.OverlayPath)
    if err != nil {
        return mw, err
    }

    defer response.Body.Close()
    data,err := ioutil.ReadAll(response.Body)
    if err != nil {
        return mw, err
    }

    overlay.ReadImageBlob(data)
    overlay.SetImageVirtualPixelMethod(imagick.VIRTUAL_PIXEL_TRANSPARENT)

    mw.CompositeImage(overlay, imagick.COMPOSITE_OP_OVER, false, int(layer.Left), int(layer.Top))

    return mw, err
}

func DistortLayer(channel chan *imagick.MagickWand, errors chan error, layer Layer) {
    mw := imagick.NewMagickWand()

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

    // Само искажение, самая долгая операция
    mw.DistortImage(imagick.DISTORTION_POLYNOMIAL, layer.DistortionMatrix, false)

    // Накладываем маску на слой
    mw, err = MaskLayer(mw, layer)
    if err != nil {
        errors <- err
        return
    }

    // Отдаём в канал изображение
    channel <- mw
}
