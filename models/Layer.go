package models

import (
	"gopkg.in/gographics/imagick.v2/imagick"
)

const DISTORT_POLYNOMIAL = "polynomial"
const DISTORT_PARTIAL = "partial"
const DISTORT_AFFINE = "affine"
const DISTORT_BILINEAR = "bilinear"
const DISTORT_SHEPARDS = "shepards"

const NUMB_COORDINATES_POINT = 4

// Каждая часть изображения будет расширена в ширину и высоту на это значение в процентах
const PART_SCALE = 1

// Минимальный размер расширения части изображения в пикселях
const MIN_PART_SCALE_SIZE = 3

// Множитель, число на которое умножается все изображение, для улучшения качества искажения
const MULTIPLIER = 2

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
	OverlayWidth  float64 `json:"overlay_width"`
	OverlayHeight float64 `json:"overlay_height"`
	OverlayLeft   float64 `json:"overlay_left"`
	OverlayTop    float64 `json:"overlay_top"`
}

func (l Layer) Build(channel chan PositionMagicWand, errors chan error) {
	// Основное изображение, на него будем наносить все данные
	baseImage := imagick.NewMagickWand()
	pw := imagick.NewPixelWand()
	pw.SetColor("none")
	baseImage.NewImage(uint(l.Width), uint(l.Height), pw)

	// Фон для основного изображения
	bgImg := imagick.NewMagickWand()
	bpw := imagick.NewPixelWand()
	bpw.SetColor("none")

	if l.BackgroundColor != "" {
		bpw.SetColor("#" + l.BackgroundColor)
	}

	bgImg.NewImage(uint(l.DesignWidth), uint(l.DesignHeight), bpw)

	// Добавляем фон на основное изображение
	baseImage.CompositeImage(bgImg, imagick.COMPOSITE_OP_OVER, int(l.DesignLeft), int(l.DesignTop))

	var err error
	pmw := PositionMagicWand{Layer: l}

	// Обрабатываем фон
	if l.BackgroundPath != "" {
		baseImage, err = l.ProcessBackground(baseImage)
		if err != nil {
			errors <- err
			return
		}
	}

	// Обрабатываем основной слой
	if l.Path != "" {
		baseImage, err = l.ProcessMain(baseImage)
		if err != nil {
			errors <- err
			return
		}
	}

	// Искажение основного слоя, самая долгая операция
	if len(l.DistortionMatrix) != 0 {
		baseImage, err = l.ProcessDistort(baseImage)
		if err != nil {
			errors <- err
			return
		}
	}

	// Накладываем слой наложения
	if l.OverlayPath != "" {
		baseImage, err = l.ProcessOverlay(baseImage)
		if err != nil {
			errors <- err
			return
		}
	}
	pmw.MagicWand = baseImage

	// Отдаём в канал структуры с позицией и изображением
	channel <- pmw
}

func (l Layer) ProcessBackground(baseImage *imagick.MagickWand) (*imagick.MagickWand, error) {
	background := imagick.NewMagickWand()
	x, y := 0, 0

	data, err := GetImageBlob(l.BackgroundPath)
	if err != nil {
		return baseImage, err
	}

	err = background.ReadImageBlob(data)
	if err != nil {
		return baseImage, err
	}

	switch l.BackgroundLayout {
	case "scale":
		background.ScaleImage(baseImage.GetImageWidth(), baseImage.GetImageHeight())
		break

	case "tile":
		tmp := imagick.NewMagickWand()
		pw := imagick.NewPixelWand()
		pw.SetColor("none")
		tmp.NewImage(uint(l.Width), uint(l.Height), pw)
		background = tmp.TextureImage(background)
		break

	case "center":
		x = (int(baseImage.GetImageWidth()) - int(background.GetImageWidth())) / 2
		x = (int(baseImage.GetImageHeight()) - int(background.GetImageHeight())) / 2
		break
	}

	err = baseImage.CompositeImage(background, imagick.COMPOSITE_OP_OVER, x, y)

	return baseImage, err
}

func (l Layer) ProcessMain(baseImage *imagick.MagickWand) (*imagick.MagickWand, error) {
	main := imagick.NewMagickWand()
	pw := imagick.NewPixelWand()
	pw.SetColor("none")

	main.NewImage(uint(l.DesignWidth), uint(l.DesignHeight), pw)

	data, err := GetImageBlob(l.Path)
	if err != nil {
		return baseImage, err
	}
	main.ReadImageBlob(data)
	main.ResizeImage(uint(l.DesignWidth), uint(l.DesignHeight), imagick.FILTER_CATROM, 1)

	err = baseImage.CompositeImage(main, imagick.COMPOSITE_OP_OVER, int(l.DesignLeft), int(l.DesignTop))

	return baseImage, err
}

func (l Layer) ProcessOverlay(baseImage *imagick.MagickWand) (*imagick.MagickWand, error) {
	overlay := imagick.NewMagickWand()

	data, err := GetImageBlob(l.OverlayPath)
	if err != nil {
		return overlay, err
	}

	overlay.ReadImageBlob(data)
	overlay.SetImageVirtualPixelMethod(imagick.VIRTUAL_PIXEL_TRANSPARENT)

	if l.OverlayWidth != 0 && l.OverlayHeight != 0 {
		overlay.ScaleImage(uint(l.OverlayWidth), uint(l.OverlayHeight))
	}

	err = baseImage.CompositeImage(overlay, imagick.COMPOSITE_OP_DST_OUT, int(l.OverlayLeft), int(l.OverlayTop))

	return baseImage, err
}

func (l Layer) ProcessDistort(baseImage *imagick.MagickWand) (bi *imagick.MagickWand, err error) {
	if l.DistortionType == DISTORT_POLYNOMIAL {
		bi, err = l.PolynomialDistort(baseImage)
	}

	if l.DistortionType == DISTORT_PARTIAL {
		bi, err = l.PartialDistort(baseImage)
	}
	if l.DistortionType == DISTORT_AFFINE {
		bi, err = l.AffineDistort(baseImage)
	}
	if l.DistortionType == DISTORT_BILINEAR {
		bi, err = l.BilinearDistort(baseImage)
	}
	if l.DistortionType == DISTORT_SHEPARDS {
		bi, err = l.ShepardsDistort(baseImage)
	}

	return
}

func (l Layer) PolynomialDistort(baseImage *imagick.MagickWand) (*imagick.MagickWand, error) {
	err := baseImage.DistortImage(imagick.DISTORTION_POLYNOMIAL, l.DistortionMatrix, false)

	return baseImage, err
}
func (l Layer) AffineDistort(baseImage *imagick.MagickWand) (*imagick.MagickWand, error) {
	err := baseImage.DistortImage(imagick.DISTORTION_AFFINE, l.DistortionMatrix, false)

	return baseImage, err
}
func (l Layer) BilinearDistort(baseImage *imagick.MagickWand) (*imagick.MagickWand, error) {
	err := baseImage.DistortImage(imagick.DISTORTION_BILINEAR, l.DistortionMatrix, false)

	return baseImage, err
}
func (l Layer) ShepardsDistort(baseImage *imagick.MagickWand) (*imagick.MagickWand, error) {
	err := baseImage.DistortImage(imagick.DISTORTION_SHEPARDS, l.DistortionMatrix, false)

	return baseImage, err
}

func (l Layer) PartialDistort(baseImage *imagick.MagickWand) (*imagick.MagickWand, error) {
	var err error
	width := baseImage.GetImageWidth() * MULTIPLIER
	height := baseImage.GetImageHeight() * MULTIPLIER

	resultImage := InitImage(width, height)

	matrix := DistortionVectorMatrix{}
	matrix.SetFromDistortionMatrix(l.DistortionMatrix)

	matrixParts := SplitMatrix(matrix.VectorMatrix, 2, 2)

	for _, matrixPart := range matrixParts {
		dvm := DistortionVectorMatrix{VectorMatrix: matrixPart}
		dvm.Multiply(MULTIPLIER)

		// Конвертация структуры в массив, понятный ImageMagick
		matrix := dvm.GetDistortionMatrix()

		partSumWidth := dvm.GetWidth() * PART_SCALE / 100
		partSumHeight := dvm.GetHeight() * PART_SCALE / 100

		if partSumWidth < MIN_PART_SCALE_SIZE {
			partSumWidth = MIN_PART_SCALE_SIZE
		}
		if partSumHeight < MIN_PART_SCALE_SIZE {
			partSumHeight = MIN_PART_SCALE_SIZE
		}

		imagePart := InitImage(uint(l.Width), uint(l.Height))
		imagePart.CompositeImage(baseImage, imagick.COMPOSITE_OP_OVER, 0, 0)

		imagePart.ScaleImage(width, height)
		imagePart.CropImage(
			uint(dvm.GetWidth()+partSumWidth),
			uint(dvm.GetHeight()+partSumHeight),
			int(dvm.GetLeft()),
			int(dvm.GetTop()),
		)

		fullImagePart := InitImage(width, height)
		err = fullImagePart.CompositeImage(imagePart, imagick.COMPOSITE_OP_OVER, int(dvm.GetLeft()), int(dvm.GetTop()))
		err = fullImagePart.DistortImage(imagick.DISTORTION_BILINEAR, matrix, false)

		err = resultImage.CompositeImage(fullImagePart, imagick.COMPOSITE_OP_OVER, 0, 0)
	}

	// Возвращаем нормальный размер слою
	if MULTIPLIER != 1 {
		resultImage.ScaleImage(width/MULTIPLIER, height/MULTIPLIER)
	}

	return resultImage, err
}
