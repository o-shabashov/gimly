package models

import (
    "github.com/ant0ine/go-json-rest/rest"
)

type PostData struct {
    Format string   `json:"format"`
    Width  uint     `json:"width"`
    Height uint     `json:"height"`
    Layers [] Layer `json:"layers"`
}

func (p *PostData) Validate(r *rest.Request) error {
    // Конвертация JSON в структуру
    if err := r.DecodeJsonPayload(p); err != nil {
        return err
    }

    return nil
}
