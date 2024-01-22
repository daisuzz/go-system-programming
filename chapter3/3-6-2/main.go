package main

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

var source = `１行目
２行目
３行目`

func main() {
	reader := bufio.NewReader(strings.NewReader(source))
	for {
		// 改行文字まで読み込む
		line, err := reader.ReadString('\n')
		fmt.Printf("%#v\n", line)
		// 終端文字の場合はループを抜ける
		if err == io.EOF {
			break
		}
	}
}
