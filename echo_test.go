package iofuncs

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestEchoSpec(t *testing.T) {

	mockConn := createMockConnection()
	mockConn.input = []byte("test input")

	Convey("Given I echo a connection", t, func() {

		go Echo(mockConn)

		Convey("It should echo input", func() {
			So(string(mockConn.output), ShouldEqual, "test input")
		})

	})

}
