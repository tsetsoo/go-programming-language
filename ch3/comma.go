package main

import (
	"bytes"
	"fmt"
)

func main() {
	fmt.Println(comma("1113431443"))
}

func comma(s string) string {
	var buf bytes.Buffer
	n := len(s)
	fmt.Println(n)
	for i := 0; i < len(s); i++ {
		n--
		buf.WriteByte(s[i])
		if n != 0 && n%3 == 0 {
			buf.WriteString(",")
		}
	}
	return buf.String()
}
