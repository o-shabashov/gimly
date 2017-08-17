package models

import (
    "gopkg.in/gographics/imagick.v2/imagick"
)

func ProcessMain(layer Layer, baseImage *imagick.MagickWand) (*imagick.MagickWand, error) {
    main := imagick.NewMagickWand()
    pw := imagick.NewPixelWand()
    pw.SetColor("none")

    main.NewImage(uint(layer.DesignWidth), uint(layer.DesignHeight), pw)

    data, err := GetImageBlob(layer.Path)
    if err != nil {
        return baseImage, err
    }
    main.ReadImageBlob(data)
    main.ResizeImage(uint(layer.DesignWidth), uint(layer.DesignHeight), imagick.FILTER_CATROM, 1)

    baseImage.CompositeImage(main, imagick.COMPOSITE_OP_OVER, int(layer.DesignLeft), int(layer.DesignTop))

    return baseImage, err
}