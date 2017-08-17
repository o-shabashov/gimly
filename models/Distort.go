package models

import (
    "gopkg.in/gographics/imagick.v2/imagick"
)

// TODO правильный тип искажения, на основе запроса
func ProcessDistort(layer Layer, baseImage *imagick.MagickWand) (*imagick.MagickWand) {
    baseImage.DistortImage(imagick.DISTORTION_POLYNOMIAL, layer.DistortionMatrix, false)

    return baseImage
}
