package main

import (
	"fmt"
	"log"
	"os"

	"tsvetelinpantev.com/go-programming-language/ch5/links"
)

var tokens = make(chan struct{}, 20)

func crawl(url string) []string {
	fmt.Println(url)
	tokens <- struct{}{} // acquire a token
	list, err := links.Extract(url)
	<-tokens // release the token

	if err != nil {
		log.Print(err)
	}
	return list
}

type toPass struct {
	links []string
	depth int
}

func main() {
	worklist := make(chan toPass)
	var n int // number of pending sends to worklist

	// Start with the command-line arguments.
	n++
	go func() { worklist <- toPass{os.Args[1:], 0} }()

	// Crawl the web concurrently.
	seen := make(map[string]bool)
	for ; n > 0; n-- {
		list := <-worklist
		if list.depth > 2 {
			continue
		}
		for _, link := range list.links {
			if !seen[link] {
				seen[link] = true
				n++
				go func(link string) {
					worklist <- toPass{crawl(link), list.depth + 1}
				}(link)
			}
		}
	}
}
