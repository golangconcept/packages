package main

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestMathOperations(t *testing.T) {
	Convey("Given two numbers", t, func() {
		a := 2
		b := 3

		Convey("When we add them", func() {
			result := a + b

			Convey("Then the result should be 5", func() {
				So(result, ShouldEqual, 5)
			})
		})

		Convey("When we multiply them", func() {
			result := a * b

			Convey("Then the result should be 6", func() {
				So(result, ShouldEqual, 6)
			})
		})
	})
}
