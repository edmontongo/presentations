package march_test

import (
	"testing"
	. "github.com/edmontongo/presentations/2014-03/asm"
	. "github.com/smartystreets/goconvey/convey"
)

type Test struct {
	a, b, res int64
}

func TestEdmontongo(t *testing.T) {
	Convey("it should add", t, func() {
		tests := []Test{
			{10, 5, 15},
			{10, 0, 10},
			{10, -15, -5},
			{10, 25, 35},
		}
		for _, ex := range tests {
			So(Add(ex.a, ex.b), ShouldEqual, ex.res)
		}
	})

	Convey("it should sub", t, func() {
		tests := []Test{
			{10, 5, 5},
			{10, 0, 10},
			{10, -15, 25},
			{10, 25, -15},
		}
		for _, ex := range tests {
			So(Sub(ex.a, ex.b), ShouldEqual, ex.res)
		}
	})

	Convey("it should multiply", t, func() {
		tests := []Test{
			{10, 5, 50},
			{10, 0, 0},
			{10, -15, -150},
			{10, 25, 250},
		}
		for _, ex := range tests {
			So(Mul(ex.a, ex.b), ShouldEqual, ex.res)
		}
	})
}
