package iofuncs

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestMultiWriter(t *testing.T) {
	Convey("I can add a Reader in the read position and it will be read from", t, func() {
		mockReader := createMockConnection()
		mockReader.input = []byte("test")

		multiWriter := CreateMultiWriter(mockReader)

		result := make([]byte, 4)
		multiWriter.Read(result)
		So(string(result), ShouldEqual, "test")
	})

	Convey("I can add Writers and they will all be written to", t, func() {
		mockReader := createMockConnection()
		mockWriter1 := createMockConnection()
		mockWriter2 := createMockConnection()

		multiWriter := CreateMultiWriter(mockReader, mockWriter1, mockWriter2)
		multiWriter.Write([]byte("test"))

		So(string(mockWriter1.output), ShouldEqual, "test")
		So(string(mockWriter2.output), ShouldEqual, "test")
	})

	Convey("Closing the MultiWriter closes the reader and all writers", t, func() {
		mockReader := createMockConnection()
		mockWriter1 := createMockConnection()
		mockWriter2 := createMockConnection()

		multiWriter := CreateMultiWriter(mockReader, mockWriter1, mockWriter2)
		multiWriter.Close()

		So(mockReader.closed, ShouldBeTrue)
		So(mockWriter1.closed, ShouldBeTrue)
		So(mockWriter2.closed, ShouldBeTrue)
	})
}
