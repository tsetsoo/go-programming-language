package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
)

func main() {
	writer, count := CountingWriter(os.Stdout)
	writer.Write([]byte("aaa"))
	fmt.Println(*count)
}

type ByteCounter int

func (c *ByteCounter) Write(b []byte) (int, error) {
	*c += ByteCounter(len(b))

	return len(b), nil
}

type WordCounter int

func (c *WordCounter) Write(b []byte) (int, error) {
	scanner := bufio.NewScanner(bytes.NewReader(b))
	scanner.Split(bufio.ScanWords)
	wordsCounted := 0
	for scanner.Scan() {
		wordsCounted += 1
	}
	*c += WordCounter(wordsCounted)
	return wordsCounted, nil
}

type LineCounter int

func (c *LineCounter) Write(b []byte) (int, error) {
	scanner := bufio.NewScanner(bytes.NewReader(b))
	scanner.Split(bufio.ScanLines)
	linesCounted := 0
	for scanner.Scan() {
		linesCounted += 1
	}
	*c += LineCounter(linesCounted)
	return linesCounted, nil
}

func CountingWriter(w io.Writer) (io.Writer, *int64) {
	counter := int64(0)
	fn := func(p []byte) (n int, err error) {
		counter += int64(len(p))
		return w.Write(p)
	}

	return ByteCounterFunc(fn), &counter
}

type ByteCounterFunc func(p []byte) (n int, err error)

func (f ByteCounterFunc) Write(p []byte) (n int, err error) {
	return f(p)
}
