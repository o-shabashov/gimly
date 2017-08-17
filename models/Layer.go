package models

import (
    "gopkg.in/gographics/imagick.v2/imagick"
    "net/http"
    "io/ioutil"
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
    BackgroundColor  string    `json:"background_color"`
    BackgroundPath   string    `json:"background_path"`
    BackgroundLayout string    `json:"background_layout"`

    // Вообще этих полей нет в JSON схеме. Но они добавляются в процессе конвертации PostData.ConvertPositioning()
    OverlayWidth  float64   `json:"overlay_width"`
    OverlayHeight float64   `json:"overlay_height"`
    OverlayLeft   float64   `json:"overlay_left"`
    OverlayTop    float64   `json:"overlay_top"`
}

const DISTORT_POLYNOMIAL string = "polynomial"
const NUMB_COORDINATES_POINT int = 4

func Build(channel chan PositionMagicWand, errors chan error, layer Layer) {
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

func ProcessBackground(layer Layer, baseImage *imagick.MagickWand) (*imagick.MagickWand, error) {
    background := imagick.NewMagickWand()
    x, y := 0, 0

    data, err := getImageBlob(layer.BackgroundPath)
    if err != nil {
        return baseImage, err
    }

    background.ReadImageBlob(data)

    switch layer.BackgroundLayout {
    case "scale":
        background.ScaleImage(baseImage.GetImageWidth(), baseImage.GetImageHeight())
        break

    case "tile":
        tmp := imagick.NewMagickWand()
        pw := imagick.NewPixelWand()
        pw.SetColor("none")
        tmp.NewImage(uint(layer.Width), uint(layer.Height), pw)
        background = tmp.TextureImage(background)
        break

    case "center":
        x = (int(baseImage.GetImageWidth()) - int(background.GetImageWidth())) / 2
        x = (int(baseImage.GetImageHeight()) - int(background.GetImageHeight())) / 2

        break
    }

    baseImage.CompositeImage(background, imagick.COMPOSITE_OP_OVER, x, y)

    return baseImage, err
}

func ProcessMain(layer Layer, baseImage *imagick.MagickWand) (*imagick.MagickWand, error) {
    main := imagick.NewMagickWand()
    pw := imagick.NewPixelWand()
    pw.SetColor("none")

    main.NewImage(uint(layer.DesignWidth), uint(layer.DesignHeight), pw)

    data, err := getImageBlob(layer.Path)
    if err != nil {
        return baseImage, err
    }
    main.ReadImageBlob(data)
    main.ResizeImage(uint(layer.DesignWidth), uint(layer.DesignHeight), imagick.FILTER_CATROM, 1)

    baseImage.CompositeImage(main, imagick.COMPOSITE_OP_OVER, int(layer.DesignLeft), int(layer.DesignTop))

    return baseImage, err
}

// TODO правильный тип искажения, на основе запроса
func ProcessDistort(layer Layer, baseImage *imagick.MagickWand) (*imagick.MagickWand) {
    baseImage.DistortImage(imagick.DISTORTION_POLYNOMIAL, layer.DistortionMatrix, false)

    return baseImage
}

func ProcessOverlay(baseImage *imagick.MagickWand, layer Layer) (*imagick.MagickWand, error) {
    overlay := imagick.NewMagickWand()

    data, err := getImageBlob(layer.OverlayPath)
    if err != nil {
        return overlay, err
    }

    overlay.ReadImageBlob(data)
    overlay.SetImageVirtualPixelMethod(imagick.VIRTUAL_PIXEL_TRANSPARENT)

    if layer.OverlayWidth != 0 && layer.OverlayHeight != 0 {
        overlay.ScaleImage(uint(layer.OverlayWidth), uint(layer.OverlayHeight))
    }

    baseImage.CompositeImage(overlay, imagick.COMPOSITE_OP_DST_OUT, int(layer.OverlayLeft), int(layer.OverlayTop))

    return baseImage, err
}

// Получение картинки дизайна по HTTP
func getImageBlob(path string) ([]byte, error) {
    response, err := http.Get(path)

    defer response.Body.Close()
    data, err := ioutil.ReadAll(response.Body)

    return data, err
}
