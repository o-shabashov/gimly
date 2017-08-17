package models

import (
    "gopkg.in/gographics/imagick.v2/imagick"
    "net/http"
    "io/ioutil"
)

const DISTORT_POLYNOMIAL string = "polynomial"
const NUMB_COORDINATES_POINT int = 4

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
    BackgroundColor  string    `json:"background_color"`
    BackgroundPath   string    `json:"background_path"`
    BackgroundLayout string    `json:"background_layout"`

    // Вообще этих полей нет в JSON схеме. Но они добавляются в процессе конвертации PostData.ConvertPositioning()
    OverlayWidth  float64   `json:"overlay_width"`
    OverlayHeight float64   `json:"overlay_height"`
    OverlayLeft   float64   `json:"overlay_left"`
    OverlayTop    float64   `json:"overlay_top"`
}

func (layer Layer) Build(channel chan PositionMagicWand, errors chan error) {
    // Основное изображение, на него будем наносить все данные
    baseImage := imagick.NewMagickWand()
    pw := imagick.NewPixelWand()
    pw.SetColor("none")
    baseImage.NewImage(uint(layer.DesignWidth), uint(layer.DesignHeight), pw)

    var err error
    pmw := PositionMagicWand{Layer: layer}

    // Обрабатываем фон
    if layer.BackgroundPath != "" {
        baseImage, err = ProcessBackground(layer, baseImage)
        if err != nil {
            errors <- err
            return
        }
    }

    // Обрабатываем основной слой
    if layer.Path != "" {
        baseImage, err = ProcessMain(layer, baseImage)
        if err != nil {
            errors <- err
            return
        }
    }

    // Искажение основного слоя, самая долгая операция
    if len(layer.DistortionMatrix) != 0 {
        baseImage = ProcessDistort(layer, baseImage)
    }

    // Накладываем слой наложения
    if layer.OverlayPath != "" {
        baseImage, err = ProcessOverlay(baseImage, layer)
        if err != nil {
            errors <- err
            return
        }
    }
    pmw.MagicWand = baseImage

    // Отдаём в канал структуры с позицией и изображением
    channel <- pmw
}

// Получение картинки дизайна по HTTP
func GetImageBlob(path string) ([]byte, error) {
    response, err := http.Get(path)

    defer response.Body.Close()
    data, err := ioutil.ReadAll(response.Body)

    return data, err
}
