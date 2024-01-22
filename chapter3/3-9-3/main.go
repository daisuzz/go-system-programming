package main

import (
	"archive/zip"
	"io"
	"os"
	"strings"
)

func main() {
	newFile, err := os.Create("new.zip")
	if err != nil {
		panic(err)
	}
	defer newFile.Close()

	zipWriter := zip.NewWriter(newFile)
	defer zipWriter.Close()

	writer, err := zipWriter.Create("new.txt")
	if err != nil {
		panic(err)
	}
	// io.WriteStringで実ファイルを作成してzipファイルを作成する方法
	// io.WriteString(writer, "test")

	// strings.Readerからio.Writerを取得する方法
	reader := strings.NewReader("test")
	io.Copy(writer, reader)
}
