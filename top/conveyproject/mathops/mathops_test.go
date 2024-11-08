package mathops

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestAdd(t *testing.T) {
	Convey("Given two integers", t, func() {
		a := 3
		b := 4

		Convey("When they are added", func() {
			result := Add(a, b)

			Convey("Then the result should be the sum of the two integers", func() {
				So(result, ShouldEqual, 7)
			})
		})
	})
}

func TestSub(t *testing.T) {
	Convey("Given two integers", t, func() {
		Convey("When they are added", func() {

		})
	})
}
