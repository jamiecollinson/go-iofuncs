package iofuncs

import (
	"io"
)

// mockReader mocks io.Reader
type mockReader struct {
	input []byte
}

func (c *mockReader) Read(p []byte) (n int, err error) {
	n = copy(p, c.input)
	c.input = c.input[n:]
	if len(c.input) == 0 {
		err = io.EOF
	}
	return
}

func (c *mockReader) setInput(s string) { c.input = []byte(s) }

func createMockReader() *mockReader {
	return &mockReader{}
}

// mockWriter mocks io.Writer
type mockWriter struct {
	output []byte
}

func (c *mockWriter) Write(p []byte) (n int, err error) {
	c.output = p
	return len(p), nil
}

func createMockWriter() *mockWriter {
	return &mockWriter{}
}

// mockCloser mocks io.Closer
type mockCloser struct {
	closed bool
}

func (c *mockCloser) Close() (err error) {
	c.closed = true
	return
}

func createMockCloser() *mockCloser {
	return &mockCloser{}
}

// mockConnection is a mock for an io.ReadWriteCloser
type mockConnection struct {
	mockReader
	mockWriter
	mockCloser
}

func createMockConnection() *mockConnection {
	return &mockConnection{}
}
