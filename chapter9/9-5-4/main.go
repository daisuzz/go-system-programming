package main

import (
	"fmt"
	"path/filepath"
)

func main() {
	// パスをそのままクリーンにする
	fmt.Println(filepath.Clean("./path/filepath/../path.go"))
	// path/path.go

	// パスを絶対パスに整形
	abspath, _ := filepath.Abs("path/filepath/path_unix.go")
	fmt.Println(abspath)
	// /Users/daisuzz/go-system-programming/chapter9/9-5-4/path/filepath/path_unix.go

	// パスを相対パスに整形
	relpath, _ := filepath.Rel(
		"/Users/daisuzz/go-system-programming/chapter9/9-5-4",
		" /Users/daisuzz/go-system-programming/chapter9/9-5-4/path/filepath/path.go",
	)
	fmt.Println(relpath)
	// path/filepath/path.go
}
