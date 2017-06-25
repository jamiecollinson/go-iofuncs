package iofuncs

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestMockConnection(t *testing.T) {
	Convey("Given a mockConnection", t, func() {

		mockConn := createMockConnection()

		Convey("I can Read() from mockInput until it is empty", func() {
			mockConn.input = []byte("ab")
			result := make([]byte, 1)
			mockConn.Read(result)
			So(string(result), ShouldEqual, "a")
			mockConn.Read(result)
			So(string(result), ShouldEqual, "b")
			n, err := mockConn.Read(result)
			So(n, ShouldBeZeroValue)
			So(err, ShouldBeError)
		})

		Convey("I can Write() to mockOutput", func() {
			mockConn.Write([]byte("test output"))
			So(string(mockConn.output), ShouldEqual, "test output")
		})

		Convey("I can Close() and trigger closeSpy()", func() {
			So(mockConn.closed, ShouldBeFalse)
			mockConn.Close()
			So(mockConn.closed, ShouldBeTrue)
		})

	})
}
