package main

import (
	"io"
	"os"
)

func main() {
	oldFile, err := os.Open("old.txt")
	if err != nil {
		panic(err)
	}
	defer oldFile.Close()
	newFile, err := os.Create("new.txt")
	if err != nil {
		panic(err)
	}
	defer newFile.Close()
	if _, err := io.Copy(newFile, oldFile); err != nil {
		return
	}
}
