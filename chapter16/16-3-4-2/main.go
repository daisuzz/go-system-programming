package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	// Poolを作成, Newで新規作成時のコードを実装
	var count int
	pool := sync.Pool{
		New: func() interface{} {
			count++
			return fmt.Sprintf("created: %d", count)
		},
	}

	// GCを呼ぶと追加された要素が消える
	pool.Put("removed: 1")
	pool.Put("removed: 2")
	runtime.GC()
	fmt.Println(pool.Get())
}
