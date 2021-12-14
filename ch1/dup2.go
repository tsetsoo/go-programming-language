package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]map[string]int)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts, "")
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts, arg)
			f.Close()
		}
	}
	for lines, mapp := range counts {
		filecount := len(mapp)
		if filecount == 1 {
			total := 0
			for _, count := range mapp {
				total += count
			}
			if total < 1 {
				continue
			}
		}
		fmt.Printf("%s appeared in %d files\n", lines, len(mapp))
		for name, count := range mapp {
			fmt.Printf("\t%d times in %s\n", count, name)
		}
	}
}
func countLines(f *os.File, counts map[string]map[string]int, fileName string) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		if counts[input.Text()] == nil {
			counts[input.Text()] = make(map[string]int)
		}
		counts[input.Text()][fileName]++
	}
}
