package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"strings"
)

var csvSource = `長いので省略`

func main() {
	reader := strings.NewReader(csvSource)
	csvReader := csv.NewReader(reader)
	for {
		line, err := csvReader.Read()
		if err == io.EOF {
			break
		}
		fmt.Println(line[2], line[6:9])
	}
}
