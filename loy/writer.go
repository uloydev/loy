package loy

import "bytes"

type BytesWriter struct {
	buf *bytes.Buffer
}

func (w *BytesWriter) Write(p []byte) (n int, err error) {
	return w.buf.Write(p)
}
