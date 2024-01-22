package main

import (
	"encoding/csv"
	"io"
	"os"
)

func main() {
	file, err := os.Create("test.txt")
	if err != nil {
		panic(err)
	}
	multiwriter := io.MultiWriter(file, os.Stdout)
	writer := csv.NewWriter(multiwriter)
	writer.Write([]string{"1", "2", "3"})
	writer.Flush()
}
