package dvm

import (
    "testing"
    . "github.com/smartystreets/goconvey/convey"
    "gimly/test_data"
    "encoding/json"
    "gimly/models"
)

func TestSetFromDistortionMatrix(t *testing.T) {

    Convey("Is matrix chunks are divided to arrays by 2*DIMENSION in each?", t, func() {
        d := DistortionVectorMatrix{}
        layer := models.Layer{}
        err := json.Unmarshal([]byte(test_data.OverlayLayer), &layer)
        So(err, ShouldBeNil)

        result := d.SetFromDistortionMatrix(layer.DistortionMatrix)
        So(result, ShouldResemble, []DistortionVector{
            {Start: Point{Left: 0.09732, Top: 1.78571}, End: Point{Left: 0, Top: 9.06593}},
            {Start: Point{Left: 16.74777, Top: 1.78571}, End: Point{Left: 16.74777, Top: 1.78571}},
        })
    })
}
