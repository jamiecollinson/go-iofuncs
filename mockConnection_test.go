package iofuncs

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestMockReader(t *testing.T) {
	Convey("A mockReader Read()s until empty", t, func() {
		mockReader := createMockReader()
		mockReader.setInput("ab")

		result := make([]byte, 1)

		mockReader.Read(result)
		So(string(result), ShouldEqual, "a")

		mockReader.Read(result)
		So(string(result), ShouldEqual, "b")

		n, err := mockReader.Read(result)
		So(n, ShouldBeZeroValue)
		So(err, ShouldBeError)
	})
}

func TestMockWriter(t *testing.T) {
	Convey("A mockWriter Write()s to output", t, func() {
		mockWriter := createMockWriter()
		mockWriter.Write([]byte("test"))
		So(string(mockWriter.output), ShouldEqual, "test")
	})
}

func TestMockCloser(t *testing.T) {
	Convey("A mockCloser can record a Close()", t, func() {
		mockClose := createMockCloser()

		So(mockClose.closed, ShouldBeFalse)
		mockClose.Close()
		So(mockClose.closed, ShouldBeTrue)
	})
}
