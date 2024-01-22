package main

import (
	"fmt"
	"time"
)

func main() {
	c := time.After(10 * time.Second)
	<-c
	fmt.Println("timer finished")
}
