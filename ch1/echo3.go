package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	toPrint := strings.Join(os.Args[1:], " ")
	duration := time.Since(start)
	fmt.Println(duration.Nanoseconds())
	fmt.Println(toPrint)
	// fmt.Println(os.Args[1:])
	// fmt.Println(os.Args[0])
}
