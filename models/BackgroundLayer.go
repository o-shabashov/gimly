package models

import (
    "gopkg.in/gographics/imagick.v2/imagick"
)

func ProcessBackground(layer Layer, baseImage *imagick.MagickWand) (*imagick.MagickWand, error) {
    background := imagick.NewMagickWand()
    x, y := 0, 0

    data, err := GetImageBlob(layer.BackgroundPath)
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