package models

import (
    "encoding/json"
    "testing"

    "gimly/test_data"

    . "github.com/smartystreets/goconvey/convey"
    "github.com/xeipuuv/gojsonschema"
    "gopkg.in/gographics/imagick.v2/imagick"
)

const testDesignURL = "http://catalog.fm.vsemayki.ru/20283848485940a9c5b6b982.28126856"

func TestPostDataStruct(t *testing.T) {

    Convey("Validate JSON by given schema", t, func() {
        schemaLoader := gojsonschema.NewStringLoader(test_data.JsonSchema)
        documentLoader := gojsonschema.NewStringLoader(test_data.RequestPolynomialDistort)

        result, err := gojsonschema.Validate(schemaLoader, documentLoader)

        So(err, ShouldBeNil)

        if result.Valid() != true {
            Println(result.Errors())
        }

        So(result.Valid(), ShouldBeTrue)
    })

    Convey("Created PostData struct equal JSON", t, func() {
        p := PostData{}
        err := json.Unmarshal([]byte(test_data.RequestPolynomialDistort), &p)
        So(err, ShouldBeNil)

        Convey("Converting l coordinates", func() {
            p.ConvertPositioning()
            cpd := PostData{}
            err := json.Unmarshal([]byte(test_data.ConvertedPostData), &cpd)

            So(err, ShouldBeNil)
            So(p, ShouldResemble, cpd)
        })
    })

}

func TestLayersProcess(t *testing.T) {
    imagick.Initialize()
    defer imagick.Terminate()

    Convey("Get image blob by HTTP", t, func() {
        data, err := GetImageBlob(testDesignURL)

        So(err, ShouldBeNil)
        So(data, ShouldHaveSameTypeAs, []byte{})
        So(data, ShouldNotBeEmpty)
    })

    Convey("Process background l", t, func() {
        baseImage := imagick.NewMagickWand()
        pw := imagick.NewPixelWand()
        pw.SetColor("none")

        layer := Layer{}
        err := json.Unmarshal([]byte(test_data.BackgroundLayer), &layer)
        So(err, ShouldBeNil)

        baseImage.NewImage(uint(layer.DesignWidth), uint(layer.DesignHeight), pw)

        result, err := layer.ProcessBackground(baseImage)
        So(err, ShouldBeNil)

        So(result, ShouldHaveSameTypeAs, baseImage)
        So(result.GetImageBlob(), ShouldHaveSameTypeAs, []byte{})
    })

    Convey("Process main l", t, func() {
        baseImage := imagick.NewMagickWand()
        pw := imagick.NewPixelWand()
        pw.SetColor("none")

        layer := Layer{}
        err := json.Unmarshal([]byte(test_data.MainLayer), &layer)
        So(err, ShouldBeNil)

        baseImage.NewImage(uint(layer.DesignWidth), uint(layer.DesignHeight), pw)

        result, err := layer.ProcessMain(baseImage)
        So(err, ShouldBeNil)

        So(result, ShouldHaveSameTypeAs, baseImage)
        So(result.GetImageBlob(), ShouldHaveSameTypeAs, []byte{})
    })

    Convey("Process overlay l", t, func() {
        baseImage := imagick.NewMagickWand()
        pw := imagick.NewPixelWand()
        pw.SetColor("none")

        layer := Layer{}
        err := json.Unmarshal([]byte(test_data.OverlayLayer), &layer)
        So(err, ShouldBeNil)

        baseImage.NewImage(uint(layer.DesignWidth), uint(layer.DesignHeight), pw)

        result, err := layer.ProcessOverlay(baseImage)
        So(err, ShouldBeNil)

        So(result, ShouldHaveSameTypeAs, baseImage)
        So(result.GetImageBlob(), ShouldHaveSameTypeAs, []byte{})
    })

    Convey("Process distortion by main l", t, func() {
        baseImage := imagick.NewMagickWand()
        pw := imagick.NewPixelWand()
        pw.SetColor("none")

        layer := Layer{}
        err := json.Unmarshal([]byte(test_data.MainLayer), &layer)
        So(err, ShouldBeNil)

        baseImage.NewImage(uint(layer.DesignWidth), uint(layer.DesignHeight), pw)

        result, err := layer.ProcessDistort(baseImage)
        So(err, ShouldBeNil)

        So(result, ShouldHaveSameTypeAs, baseImage)
        So(result.GetImageBlob(), ShouldHaveSameTypeAs, []byte{})
    })

    Convey("Main Layer.Build() method should work fine", t, func() {
        p := PostData{}
        err := json.Unmarshal([]byte(test_data.RequestPolynomialDistort), &p)
        So(err, ShouldBeNil)

        p.ConvertPositioning()

        baseImage := imagick.NewMagickWand()
        pw := imagick.NewPixelWand()
        pw.SetColor("none")

        layer := Layer{}
        err = json.Unmarshal([]byte(test_data.BuildLayer), &layer)
        So(err, ShouldBeNil)

        baseImage.NewImage(uint(layer.DesignWidth), uint(layer.DesignHeight), pw)

        mapPositionMw := make(map[int]PositionMagicWand)
        channel := make(chan PositionMagicWand, 1)
        errors := make(chan error)

        go layer.Build(channel, errors)
        select {
        case pmw := <-channel:
            mapPositionMw[pmw.Layer.Position] = pmw
        case err := <-errors:
            Println(err.Error())
        }

        So(err, ShouldBeNil)
        So(mapPositionMw[0], ShouldHaveSameTypeAs, PositionMagicWand{})
    })
}
