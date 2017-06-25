package iofuncs

import (
	"io"
)

func Echo(conn io.ReadWriteCloser) {
	defer conn.Close()
	io.Copy(conn, conn)
}
