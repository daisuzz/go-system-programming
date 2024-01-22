package main

import (
	"crypto/rand"
	"io"
	"os"
)

func main() {
	randReader := rand.Reader
	newFile, err := os.Create("new.txt")
	defer newFile.Close()
	if err != nil {
		panic(err)
	}
	io.CopyN(newFile, randReader, 1024)
}
