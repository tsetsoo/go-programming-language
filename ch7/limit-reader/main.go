package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	limitReader := LimitReader(os.Stdin, 15)
	reader := bufio.NewReader(limitReader)
	read, err := reader.ReadString('\n')

	if err == io.EOF {
		fmt.Println("too much")
	} else {
		fmt.Println("its fine", read)
	}
}

func LimitReader(r io.Reader, limit int64) io.Reader {
	counter := int64(0)
	fn := func(p []byte) (n int, err error) {
		readBytes, err := r.Read(p)
		if err != nil {
			return readBytes, err
		}
		totalBytes := counter + int64(readBytes)
		if totalBytes > limit {
			return int(totalBytes - limit), io.EOF
		}
		counter = totalBytes
		return readBytes, nil
	}

	return ReaderFunc(fn)
}

type ReaderFunc func(p []byte) (n int, err error)

func (f ReaderFunc) Read(p []byte) (n int, err error) {
	return f(p)
}
