package models

import "gopkg.in/gographics/imagick.v3/imagick"

// Эта структура будет хранить искаженный слой и его порядок, для пердачи по каналу из горутины в основной трэд
type PositionMagicWand struct{
    Position  int
    MagicWand *imagick.MagickWand
}
