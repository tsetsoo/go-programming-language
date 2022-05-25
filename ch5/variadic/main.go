package main

import (
	"fmt"
	"math"
	"net/http"
	"os"

	"golang.org/x/net/html"
)

func min(vals ...int) int {
	min := math.MaxInt
	for _, val := range vals {
		if val < min {
			min = val
		}
	}
	return min
}

func max(vals ...int) int {
	max := math.MinInt
	for _, val := range vals {
		if val > max {
			max = val
		}
	}
	return max
}

func stringJoin(delim string, vals ...string) string {
	if len(vals) == 0 {
		return ""
	}
	result := ""
	for i := 0; i < len(vals)-1; i++ {
		result += fmt.Sprintf("%s%s", vals[i], delim)
	}
	result += vals[len(vals)-1]
	return result
}

func breadthFirst(f func(item *html.Node) []*html.Node, worklist []*html.Node, addFunc func(item *html.Node)) {
	for len(worklist) > 0 {
		items := worklist
		worklist = nil
		for _, item := range items {
			addFunc(item)
			worklist = append(worklist, f(item)...)
		}
	}
}

func crawl(doc *html.Node) []*html.Node {
	children := make([]*html.Node, 0)
	for c := doc.FirstChild; c != nil; c = c.NextSibling {
		children = append(children, c)
	}
	return children
}

func ElementsByTagName(doc *html.Node, names ...string) []*html.Node {
	toReturn := make([]*html.Node, 0)
	if len(names) == 0 {
		return toReturn
	}
	namesSet := make(map[string]struct{}, len(names))
	for _, name := range names {
		namesSet[name] = struct{}{}
	}

	addFunc := func(nodeToCheck *html.Node) {
		if nodeToCheck.Type != html.TextNode {
			_, ok := namesSet[nodeToCheck.Data]
			if ok {
				toReturn = append(toReturn, nodeToCheck)
			}
		}
	}

	breadthFirst(crawl, []*html.Node{doc}, addFunc)

	return toReturn
}

func main() {
	// fmt.Println(stringJoin(",", "kek", "lelemale", "opa"))
	// fmt.Println(stringJoin(",", "alonekek"))
	// fmt.Println(stringJoin(","))
	url := os.Args[1]
	resp, err := http.Get(url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
		os.Exit(1)
	}
	defer resp.Body.Close()

	doc, err := html.Parse(resp.Body)
	fmt.Println(ElementsByTagName(doc, "h1"))
	fmt.Println(ElementsByTagName(doc, "h2"))
	fmt.Println(ElementsByTagName(doc, "h1", "h2"))
}
