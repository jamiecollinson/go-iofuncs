package iofuncs

import (
	"io"
)

// Forward forwards one connection to another
func Forward(conn1 io.ReadWriteCloser, conn2 io.ReadWriteCloser) {
	go func() {
		defer conn1.Close()
		defer conn2.Close()

		io.Copy(conn1, conn2)
	}()
	go func() {
		defer conn1.Close()
		defer conn2.Close()

		io.Copy(conn2, conn1)
	}()
}
