package main

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestIntegerStuff(t *testing.T) {
	Convey("#Sum()", t, func() {

		Convey("Plus", func() {
			Convey("sum = int + int", func() {
				So(Add(-2, 3), ShouldEqual, 1)
				So(Add(2, -3), ShouldEqual, -1)
				So(Add(2, 3), ShouldEqual, 5)
				So(Add(-2, -3), ShouldEqual, -5)
			})

		})

		Convey("Minus", func() {

			Convey("sum = int - int", func() {
				So(Minus(3, 1), ShouldEqual, 2)
				So(Minus(-3, -1), ShouldEqual, -2)
				So(Minus(-3, 1), ShouldEqual, -4)
				So(Minus(3, -1), ShouldEqual, 4)
			})

		})

		Convey("Multiplies", func() {

			Convey("sum = int * int", func() {
				So(Multiplies(3, 1), ShouldEqual, 3)
				So(Multiplies(-3, -1), ShouldEqual, 3)
				So(Multiplies(-3, 1), ShouldEqual, -3)
				So(Multiplies(3, -1), ShouldEqual, -3)
			})

		})

		Convey("Divided by", func() {

			Convey("sum = int / int", func() {
				sum, err := DividedBy(3, 1)
				So(sum, ShouldEqual, 3)
				So(err, ShouldEqual, nil)
			})
			Convey("sum = int / 0", func() {
				sum, err := DividedBy(3, 0)
				So(sum, ShouldEqual, 0)
				So(err, ShouldBeError, "** y != 0")
			})

		})

	})

}
