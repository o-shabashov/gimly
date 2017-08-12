package models

import (
    "github.com/ant0ine/go-json-rest/rest"
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

func (p *PostData) ConvertPositioning() {
    for index := range p.Layers {
        cl := p.Layers[index]

        cl.Left   = p.Width * cl.Left / 100
        cl.Width  = p.Width * cl.Width / 100
        cl.Top    = p.Height * cl.Top / 100
        cl.Height = p.Height * cl.Height / 100

        cl.DesignLeft   = cl.Width * cl.DesignLeft / 100
        cl.DesignWidth  = cl.Width * cl.DesignWidth / 100
        cl.DesignTop    = cl.Height * cl.DesignTop / 100
        cl.DesignHeight = cl.Height * cl.DesignHeight / 100

        cl.OverlayWidth  = p.Width
        cl.OverlayHeight = p.Height
        cl.OverlayLeft   = -cl.Height
        cl.OverlayTop    = -cl.Top

        // TODO описание
        if len(cl.DistortionMatrix) != 0 {
            for index := range cl.DistortionMatrix {
                // Это X координата
                if index%2 == 0 {
                    cl.DistortionMatrix[index] = cl.DistortionMatrix[index] * p.Width / 100
                } else {
                    // Это Y координата
                    cl.DistortionMatrix[index] = cl.DistortionMatrix[index] * p.Height / 100
                }
            }
        }
    }
}
