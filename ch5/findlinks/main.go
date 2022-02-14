package main

import (
	"fmt"
	"net/http"
	"os"

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
	for _, link := range visitLoop(doc) {
		fmt.Println(link)
	}
}

func visitRecursive(links []string, n *html.Node) []string {
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				links = append(links, a.Val)
			}
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		links = visitRecursive(links, c)
	}
	return links
}

func visitLoop(n *html.Node) []string {
	nodes := make([]*html.Node, 0)
	nodes = append(nodes, n)
	links := make([]string, 0)

	for i := 0; ; i++ {
		if i == len(nodes) {
			break
		}
		currentNode := nodes[i]
		if currentNode.Type == html.ElementNode && (currentNode.Data == "a" || currentNode.Data == "script" || currentNode.Data == "style" || currentNode.Data == "img") {
			for _, a := range currentNode.Attr {
				if a.Key == "href" || a.Key == "src" {
					links = append(links, a.Val)
				}
			}
		}
		for c := currentNode.FirstChild; c != nil; c = c.NextSibling {
			nodes = append(nodes, c)
		}
	}
	return links
}
