package models

import (
    "gopkg.in/gographics/imagick.v2/imagick"
    "gimly/dvm"
)

const MULTIPLIER  = 2

// TODO правильный тип искажения, на основе запроса
func ProcessDistort(layer Layer, baseImage *imagick.MagickWand) (*imagick.MagickWand, error) {
    err := baseImage.DistortImage(imagick.DISTORTION_POLYNOMIAL, layer.DistortionMatrix, false)

    return baseImage, err
}

func Polynomial(layer Layer, baseImage *imagick.MagickWand) (*imagick.MagickWand, error) {
    err := baseImage.DistortImage(imagick.DISTORTION_POLYNOMIAL, layer.DistortionMatrix, false)

    return baseImage, err
}

func PartialDistort(layer Layer, baseImage *imagick.MagickWand) (*imagick.MagickWand, error){
    var resultImage *imagick.MagickWand
    var err error
    width := baseImage.GetImageWidth() * MULTIPLIER
    height := baseImage.GetImageHeight() * MULTIPLIER

    sampleImage := imagick.NewMagickWand()
    pw := imagick.NewPixelWand()
    pw.SetColor("none")

    sampleImage.NewImage(height, width, pw)
    sampleImage.SetImageVirtualPixelMethod(imagick.VIRTUAL_PIXEL_TRANSPARENT)

    // Copying data from the memory of sampleImage onto the data of resultImage. They continue to remain distinct areas
    // of memory, so updates will not propagate.
    // See https://stackoverflow.com/questions/21011023/copy-pointer-values-a-b-in-golang
    *resultImage = *sampleImage

    matrix := dvm.DistortionVectorMatrix{}
    matrix.SetFromDistortionMatrix(layer.DistortionMatrix)

    matrixParts := dvm.SplitMatrix(matrix.VectorMatrix, 2, 2)

    return resultImage, err
}