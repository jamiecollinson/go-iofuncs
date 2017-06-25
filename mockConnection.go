package iofuncs

import (
	"io"
)

// mockConnection is a mock for an io.ReadWriteCloser
type mockConnection struct {
	input  []byte
	output []byte
	closed bool
}

func (c *mockConnection) Read(p []byte) (n int, err error) {
	n = copy(p, c.input)
	c.input = c.input[n:]
	if len(c.input) == 0 {
		err = io.EOF
	}
	return
}

func (c *mockConnection) Write(p []byte) (n int, err error) {
	c.output = p
	return len(p), nil
}

func (c *mockConnection) Close() error {
	c.closed = true
	return nil
}

func createMockConnection() *mockConnection {
	return &mockConnection{}
}
