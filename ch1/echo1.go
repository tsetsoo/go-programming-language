package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	var s, sep string
	start := time.Now()
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	duration := time.Since(start)
	fmt.Println(duration.Nanoseconds())
	fmt.Println(s)
}
