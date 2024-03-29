package main

import (
	"fmt"
	"time"
)

func main() {
	tasks := []string{
		"cmake ..",
		"cmake . --build Release",
		"cpack",
	}
	for _, task := range tasks {
		go func() {
			// goroutineが起動する時にはループが周りきって
			// 全部のtaskが最後のタスクになってしまう
			fmt.Println(task)
		}()
	}
	time.Sleep(time.Second)
}
