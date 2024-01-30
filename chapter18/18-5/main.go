package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	fmt.Println(now.Format(time.RFC822))
	// 27 Aug 18 11:31 JST
	fmt.Println(now.Format("2006/01/02 03:04:05 MST"))
	// 2017/08/27 11:31:53 JST

}
