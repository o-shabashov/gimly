package models

import (
    "encoding/json"
    "github.com/xeipuuv/gojsonschema"
    "fmt"
    "github.com/ant0ine/go-json-rest/rest"
)

type Request struct {
    Format string   `json:"format"`
    Height uint     `json:"height"`
    Layers [] Layer `json:"layers"`
    Width  uint     `json:"width"`
}

func ReadRequest(r *rest.Request) (request Request) {
    decoder := json.NewDecoder(r.Body)
    err := decoder.Decode(&request)
    if err != nil {
        panic(err) // TODO сделать rest.Error()
    }

    // TODO validate
    return
}

func validateJson() {
    schemaLoader := gojsonschema.NewReferenceLoader("file:///home/me/schema.json")
    documentLoader := gojsonschema.NewReferenceLoader("file:///home/me/document.json")

    result, err := gojsonschema.Validate(schemaLoader, documentLoader)
    if err != nil {
        panic(err.Error())
    }

    if result.Valid() {
        fmt.Printf("The document is valid\n")
    } else {
        fmt.Printf("The document is not valid. see errors :\n")
        for _, desc := range result.Errors() {
            fmt.Printf("- %s\n", desc)
        }
    }
}
