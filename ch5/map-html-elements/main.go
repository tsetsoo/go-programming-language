package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"golang.org/x/net/html"
)

func main() {
	url := os.Args[1]
	resp, err := http.Get(url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
		os.Exit(1)
	}
	defer resp.Body.Close()

	doc, err := html.Parse(resp.Body)
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks1: %v\n", err)
		os.Exit(1)
	}
	fmt.Println("Element Name   Element Count")
	elementsMap := make(map[string]int)
	for elementName, elementCount := range visit(elementsMap, doc) {
		fmt.Println(elementName, strings.Repeat(" ", 13-len(elementName)), elementCount)
	}
}

func visit(elementsMap map[string]int, n *html.Node) map[string]int {
	if n.Type == html.ElementNode {
		elementsMap[n.Data] = elementsMap[n.Data] + 1
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		elementsMap = visit(elementsMap, c)
	}
	return elementsMap
}
