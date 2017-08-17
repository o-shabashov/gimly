package models

import (
    "gopkg.in/gographics/imagick.v2/imagick"
)

func ProcessOverlay(baseImage *imagick.MagickWand, layer Layer) (*imagick.MagickWand, error) {
    overlay := imagick.NewMagickWand()

    data, err := GetImageBlob(layer.OverlayPath)
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