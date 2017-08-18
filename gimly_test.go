package main

import (
    "testing"
    . "github.com/smartystreets/goconvey/convey"
    "gimly/models"
    "github.com/xeipuuv/gojsonschema"
    "encoding/json"
    "gopkg.in/gographics/imagick.v2/imagick"
    "gimly/gimly_test"
)

const testDesignURL = "http://catalog.fm.vsemayki.ru/20283848485940a9c5b6b982.28126856"

func TestPostDataStruct(t *testing.T) {

    Convey("Validate JSON by given schema", t, func() {
        schemaLoader := gojsonschema.NewStringLoader(gimly_test.JsonSchema)
        documentLoader := gojsonschema.NewStringLoader(gimly_test.Request)

        result, err := gojsonschema.Validate(schemaLoader, documentLoader)

        So(err, ShouldBeNil)
        So(result.Valid(), ShouldBeTrue)
    })

    Convey("Created PostData struct equal JSON", t, func() {
        p := models.PostData{}
        err := json.Unmarshal([]byte(gimly_test.Request), &p)
        So(err, ShouldBeNil)

        Convey("Converting layer coordinates", func() {
            p.ConvertPositioning()
            cpd := models.PostData{}
            err := json.Unmarshal([]byte(gimly_test.ConvertedPostData), &cpd)

            So(err, ShouldBeNil)
            So(p, ShouldResemble, cpd)
        })
    })

}

func TestLayer(t *testing.T) {
    imagick.Initialize()
    defer imagick.Terminate()

    Convey("Get image blob by HTTP", t, func() {
        data, err := models.GetImageBlob(testDesignURL)

        So(err, ShouldBeNil)
        So(data, ShouldHaveSameTypeAs, []byte{})
        So(data, ShouldNotBeEmpty)
    })

    Convey("Process background layer", t, func() {
        baseImage := imagick.NewMagickWand()
        pw := imagick.NewPixelWand()
        pw.SetColor("none")

        layer := models.Layer{}
        err := json.Unmarshal([]byte(gimly_test.BackgroundLayer), &layer)
        So(err, ShouldBeNil)

        baseImage.NewImage(uint(layer.DesignWidth), uint(layer.DesignHeight), pw)

        result, err := models.ProcessBackground(layer, baseImage)
        So(err, ShouldBeNil)

        So(result, ShouldHaveSameTypeAs, baseImage)
        So(result.GetImageBlob(), ShouldHaveSameTypeAs, []byte{})
    })

    Convey("Process main layer", t, func() {
        baseImage := imagick.NewMagickWand()
        pw := imagick.NewPixelWand()
        pw.SetColor("none")

        layer := models.Layer{}
        err := json.Unmarshal([]byte(gimly_test.MainLayer), &layer)
        So(err, ShouldBeNil)

        baseImage.NewImage(uint(layer.DesignWidth), uint(layer.DesignHeight), pw)

        result, err := models.ProcessMain(layer, baseImage)
        So(err, ShouldBeNil)

        So(result, ShouldHaveSameTypeAs, baseImage)
        So(result.GetImageBlob(), ShouldHaveSameTypeAs, []byte{})
    })
}
