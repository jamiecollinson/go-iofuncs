package iofuncs

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestForward(t *testing.T) {
	Convey("Given two connections are forwarded", t, func() {

		mockConn1 := createMockConnection()
		mockConn1.input = []byte("input1")

		mockConn2 := createMockConnection()
		mockConn2.input = []byte("input2")

		go Forward(mockConn1, mockConn2)

		Convey("input2 -> output1", func() {
			So(string(mockConn1.output), ShouldEqual, "input2")
		})

		Convey("input1 -> output2", func() {
			So(string(mockConn2.output), ShouldEqual, "input1")
		})

	})
}
