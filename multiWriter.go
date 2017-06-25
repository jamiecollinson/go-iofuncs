package iofuncs

import (
	"io"
)

type MultiWriter struct {
	reader  io.ReadCloser
	writers []io.WriteCloser
}

func (m *MultiWriter) Read(p []byte) (n int, err error) {
	return m.reader.Read(p)
}

func (m *MultiWriter) Write(p []byte) (n int, err error) {
	// TODO: decide on n / err behaviour - we're currently returning last
	for _, w := range m.writers {
		n, err = w.Write(p)
	}
	return
}

func (m *MultiWriter) Close() (err error) {
	// TODO: decide on err behaviour - currently returning last
	err = m.reader.Close()
	for _, w := range m.writers {
		err = w.Close()
	}
	return
}

func CreateMultiWriter(reader io.ReadCloser, writers ...io.WriteCloser) *MultiWriter {
	return &MultiWriter{reader, writers}
}
