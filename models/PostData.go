package models

import (
    "github.com/ant0ine/go-json-rest/rest"
    "fmt"
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
    array := []float64{1, 2, 3, 4, 5}
    fmt.Println(array)

    for index, element := range array {

        // Это X координата
        if index%2 == 0 {
            array[index] = array[index] * 2
        } else {
            // Это Y координата
            array[index] = array[index] * 3
        }
        fmt.Println(index, element)
    }
    fmt.Println(array)
    return nil
}
