package models

import "gopkg.in/gographics/imagick.v2/imagick"

// Эта структура будет хранить объект с искаженным слоем (*imagick.MagickWand) и оригинальный слой (models.l).
// Это нужно для передачи в основной тред информации о позиции и координатах смещения оригинального слоя.
// И чтобы не выдумывать новые аттрибуты, было решено для упрощения передавать весь объект models.Layer.
type PositionMagicWand struct {
	Layer     Layer
	MagicWand *imagick.MagickWand
}
