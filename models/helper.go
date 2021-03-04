package models

import (
	"errors"
	"io/ioutil"
	"net/http"

	"gopkg.in/gographics/imagick.v2/imagick"
)

// Получение картинки дизайна по HTTP
func GetImageBlob(path string) ([]byte, error) {
	response, err := http.Get(path)

	defer response.Body.Close()
	data, err := ioutil.ReadAll(response.Body)

	return data, err
}

// Возвращает массив массивов, в каждом по chunkSize элементов.
func ArrayChunk(data []float64, chunkSize int) (result [][]float64, err error) {
	for i := 0; i < len(data); i += chunkSize {
		end := i + chunkSize

		// Если элементов не хватило на массив, то он будет создан частично.
		// Например, если элементов всего 5, а chunkSize = 3, то результат будет [[0,1,2] [3,4]]
		if end > len(data) {
			end = len(data)
		}

		if len(data[i:end]) < chunkSize {
			return result, errors.New("result array is smaller than ChunkSize")
		}

		result = append(result, data[i:end])
	}

	return
}

// Создаём новое изображение (холст)
func InitImage(width uint, height uint) (image *imagick.MagickWand) {
	pw := imagick.NewPixelWand()
	pw.SetColor("none")

	image = imagick.NewMagickWand()
	image.NewImage(width, height, pw)
	image.SetImageMatte(true)
	image.SetImageMatteColor(pw)
	image.SetImageVirtualPixelMethod(imagick.VIRTUAL_PIXEL_TRANSPARENT)

	return
}
