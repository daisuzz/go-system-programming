package main

import (
	"io"
	"os"
	"strings"
)

func copyN(dest io.Writer, src io.Reader, length int64) (written int64, err error) {
	limitReader := io.LimitReader(src, length)
	written, err = io.Copy(dest, limitReader)
	if err != nil {
		return 0, err
	}
	if written == length {
		return length, nil
	}
	if written < length && err == nil {
		err = io.EOF
	}
	return
}
func main() {
	reader := strings.NewReader("go,java,kotlin,python")
	copyN(os.Stdout, reader, 7)
	// io.CopyN(os.Stdout, reader, 7)
}
