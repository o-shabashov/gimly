package models

type PostData struct {
    Format string   `json:"format"`
    Height uint     `json:"height"`
    Layers [] Layer `json:"layers"`
    Width  uint     `json:"width"`
}
