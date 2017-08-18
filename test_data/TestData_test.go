package test_data

import (
    "testing"
    . "github.com/smartystreets/goconvey/convey"
    "encoding/json"
    "gimly/models"
)

func TestValidTestdata(t *testing.T) {

    Convey("Sure that all test data is valid JSON", t, func() {
        r := models.PostData{}
        err := json.Unmarshal([]byte(Request), &r)
        So(err, ShouldBeNil)

        l := models.Layer{}
        err = json.Unmarshal([]byte(BackgroundLayer), &l)
        So(err, ShouldBeNil)

        l = models.Layer{}
        err = json.Unmarshal([]byte(MainLayer), &l)
        So(err, ShouldBeNil)

        l = models.Layer{}
        err = json.Unmarshal([]byte(OverlayLayer), &l)
        So(err, ShouldBeNil)

        l = models.Layer{}
        err = json.Unmarshal([]byte(BuildLayer), &l)
        So(err, ShouldBeNil)
    })
}