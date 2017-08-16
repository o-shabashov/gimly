package models

import (
    "github.com/ant0ine/go-json-rest/rest"
    "io/ioutil"
    "encoding/json"
)

type PostData struct {
    Format string   `json:"format"`
    Width  float64  `json:"width"`
    Height float64  `json:"height"`
    Layers [] Layer `json:"layers"`
}

func (p *PostData) Validate(r *rest.Request) error {
    // Конвертация JSON в структуру
    if err := r.DecodeJsonPayload(p); err != nil {
        return err
    }

    return nil
}

// Перевод процентного позиционирования в пиксельное TODO отдавать ошибки
func (p *PostData) ConvertPositioning() {
    for i := range p.Layers {
        p.Layers[i].Left   = p.Width * p.Layers[i].Left / 100
        p.Layers[i].Width  = p.Width * p.Layers[i].Width / 100
        p.Layers[i].Top    = p.Height * p.Layers[i].Top / 100
        p.Layers[i].Height = p.Height * p.Layers[i].Height / 100

        p.Layers[i].DesignLeft   = p.Layers[i].Width * p.Layers[i].DesignLeft / 100
        p.Layers[i].DesignWidth  = p.Layers[i].Width * p.Layers[i].DesignWidth / 100
        p.Layers[i].DesignTop    = p.Layers[i].Height * p.Layers[i].DesignTop / 100
        p.Layers[i].DesignHeight = p.Layers[i].Height * p.Layers[i].DesignHeight / 100

        p.Layers[i].OverlayWidth  = p.Width
        p.Layers[i].OverlayHeight = p.Height
        p.Layers[i].OverlayLeft   = -p.Layers[i].Left
        p.Layers[i].OverlayTop    = -p.Layers[i].Top

        // Матрица искажений имеет вид массива парных координат - [X координата, Y координата, X, Y,...]
        // Здесь идёт пересчёт этих координат из процентов в пиксели.
        // Каждый X умножается на ширину слоя и делится на 100
        // Каждый Y умножается на высоту слоя и делится на 100
        if len(p.Layers[i].DistortionMatrix) != 0 {
            for j := range p.Layers[i].DistortionMatrix {
                if j%2 == 0 {
                    // Это X координата
                    p.Layers[i].DistortionMatrix[j] = p.Layers[i].DistortionMatrix[j] * p.Layers[i].Width / 100
                } else {
                    // Это Y координата
                    p.Layers[i].DistortionMatrix[j] = p.Layers[i].DistortionMatrix[j] * p.Layers[i].Height / 100
                }
            }
        }

        // Пересчитать матрицу для полиномиального искажения. В начало массива искажений добавляется коэффициент
        // искажений, зависящий от количества точек на стороне.
        if p.Layers[i].DistortionType == DISTORT_POLYNOMIAL {

            if  p.Layers[i].DistortionOrder == 0 {
                numbPoints := len(p.Layers[i].DistortionMatrix) / NUMB_COORDINATES_POINT

                if p.Layers[i].NumbPointsSide == 0 || p.Layers[i].NumbPointsSide == 2 {
                    p.Layers[i].DistortionOrder = 1.5
                } else if p.Layers[i].NumbPointsSide == 3 && numbPoints <= 15 {
                    p.Layers[i].DistortionOrder = 2
                } else if p.Layers[i].NumbPointsSide == 3 && numbPoints > 15 || p.Layers[i].NumbPointsSide == 4 {
                    p.Layers[i].DistortionOrder = 3
                } else {
                    p.Layers[i].DistortionOrder = 4
                }
            }

            p.Layers[i].DistortionMatrix = append([]float64{p.Layers[i].DistortionOrder}, p.Layers[i].DistortionMatrix...)
        }
    }
    d, _ := json.Marshal(p)
    ioutil.WriteFile("Converted config WO matrix.json", []byte(d), 0644)
}
