package main

import (
	"fmt"
)

func main() {
	fmt.Println(panicRecover())
}

func panicRecover() (returnedWithPanic string) {
	defer func() {
		p := recover()
		if p != nil {
			returnedWithPanic = p.(string)
		}
	}()
	panic("yes")
}
