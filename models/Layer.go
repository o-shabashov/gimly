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
func MaskLayer(mw *imagick.MagickWand, layer Layer) (*imagick.MagickWand) {
    overlay := imagick.NewMagickWand()

    // Получаем оверлей по HTTP от файлменеджера
    response, _ := http.Get(layer.OverlayPath)
    defer response.Body.Close()
    data, _ := ioutil.ReadAll(response.Body)

    overlay.ReadImageBlob(data)
    overlay.SetImageVirtualPixelMethod(imagick.VIRTUAL_PIXEL_TRANSPARENT)

    mw.CompositeImage(overlay, imagick.COMPOSITE_OP_OVER, false, 0, 0) // TODO top, left

    return mw
}

func DistortLayer(channel chan *imagick.MagickWand, layer Layer) {
    mw := imagick.NewMagickWand()

    // Получаем слой по HTTP от файлменеджера
    response, _ := http.Get(layer.Path)
    defer response.Body.Close()
    data, _ := ioutil.ReadAll(response.Body)

    mw.ReadImageBlob(data)
    mw.SetImageVirtualPixelMethod(imagick.VIRTUAL_PIXEL_TRANSPARENT)

    // Само искажение, самая долгая операция
    mw.DistortImage(imagick.DISTORTION_POLYNOMIAL, layer.DistortionMatrix, false)

    // Отдаём в канал изображение после наложения маски
    channel <- MaskLayer(mw, layer)
}
