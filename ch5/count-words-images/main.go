package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"golang.org/x/net/html"
)

func main() {
	words, images, err := CountWordsAndImages(os.Args[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "count words and images: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("words", words, "images", images)

}

func CountWordsAndImages(url string) (words, images int, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	doc, err := html.Parse(resp.Body)
	if err != nil {
		err = fmt.Errorf("parsing HTML: %s", err)
		return
	}
	words, images = countWordsAndImages(doc)
	return
}

func countWordsAndImages(root *html.Node) (words, images int) {
	nodes := make([]*html.Node, 0)
	nodes = append(nodes, root)

	for i := 0; ; i++ {
		if i == len(nodes) {
			break
		}
		currentNode := nodes[i]
		if currentNode.Type == html.ElementNode && currentNode.Data == "img" {
			images++
		} else if currentNode.Type == html.TextNode {
			words += len(strings.Fields(currentNode.Data))
		}
		for c := currentNode.FirstChild; c != nil; c = c.NextSibling {
			nodes = append(nodes, c)
		}
	}
	return
}
