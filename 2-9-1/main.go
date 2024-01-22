package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Fprintf(os.Stdout, "output: %d %s %f", 100, "one hundlet", 100.0)
}
