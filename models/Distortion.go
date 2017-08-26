package models

import (
    "gopkg.in/gographics/imagick.v2/imagick"
    "gimly/dvm"
)

// Каждая часть изображения будет расширена в ширину и высоту на это значение в процентах
const PART_SCALE = 1

// Минимальный размер расширения части изображения в пикселях
const MIN_PART_SCALE_SIZE = 3

// Множитель, число на которое умножается все изображение, для улучшения качества искажения
const MULTIPLIER = 2

// TODO правильный тип искажения, на основе запроса
func ProcessDistort(layer Layer, baseImage *imagick.MagickWand) (*imagick.MagickWand, error) {
    err := baseImage.DistortImage(imagick.DISTORTION_POLYNOMIAL, layer.DistortionMatrix, false)

    return baseImage, err
}

func Polynomial(layer Layer, baseImage *imagick.MagickWand) (*imagick.MagickWand, error) {
    err := baseImage.DistortImage(imagick.DISTORTION_POLYNOMIAL, layer.DistortionMatrix, false)

    return baseImage, err
}

func PartialDistort(layer Layer, baseImage *imagick.MagickWand) (*imagick.MagickWand, error) {
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

    for _, matrixPart := range matrixParts {
        matrix := dvm.DistortionVectorMatrix{}
        matrix.VectorMatrix = matrixPart
        matrix.Clone()

        matrix.Multiply(MULTIPLIER)

        partSumWidth := matrix.GetWidth() * PART_SCALE / 100
        partSumHeight := matrix.GetHeight() * PART_SCALE / 100

        if partSumWidth < MIN_PART_SCALE_SIZE {
            partSumWidth = MIN_PART_SCALE_SIZE
        }
        if partSumHeight < MIN_PART_SCALE_SIZE {
            partSumHeight = MIN_PART_SCALE_SIZE
        }

        var imagePart *imagick.MagickWand
        *imagePart = *baseImage

        imagePart.ScaleImage(width, height)
        imagePart.CropImage(
            uint(matrix.GetWidth()+partSumWidth),
            uint(matrix.GetHeight()+partSumHeight),
            int(matrix.GetLeft()),
            int(matrix.GetTop()),
        )

        var fullImagePart *imagick.MagickWand
        *fullImagePart = *sampleImage
        err = fullImagePart.CompositeImage(imagePart, imagick.COMPOSITE_OP_OVER, int(matrix.GetLeft()), int(matrix.GetTop()))
        err = fullImagePart.DistortImage(imagick.DISTORTION_BILINEAR, matrix.GetDistortionMatrix(), false)

        err = resultImage.CompositeImage(fullImagePart, imagick.COMPOSITE_OP_OVER, 0, 0)
    }

    // Просто пиздец как тупо, но так написано в оригинальном генераторе
    if (MULTIPLIER != 1) {
        resultImage.ScaleImage(width, height)
    }

    return resultImage, err
}
