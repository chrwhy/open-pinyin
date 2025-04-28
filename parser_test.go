package main

import (
	"github.com/chrwhy/open-pinyin/parser"
	"testing"

	c "github.com/smartystreets/goconvey/convey"
)

func TestSpec(t *testing.T) {

	rawPinyin := "chenhairong"
	expect := [][]string{{"chen", "hair", "rong"}}
	c.Convey("The value should be greater by one", t, func() {
		c.So(parser.Parse(rawPinyin), c.ShouldEqual, expect)
	})

}
