package models

import (
    "io/ioutil"
    "net/http"
)

// Получение картинки дизайна по HTTP
func GetImageBlob(path string) ([]byte, error) {
    response, err := http.Get(path)

    defer response.Body.Close()
    data, err := ioutil.ReadAll(response.Body)

    return data, err
}
