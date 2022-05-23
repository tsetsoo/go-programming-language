// Findlinks4 crawls the web, starting with the URLs on the command line.
package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"unicode/utf8"

	"golang.org/x/net/html"
)

func breadthFirst(f func(item string) []string, worklist []string) {
	seen := make(map[string]bool)
	for len(worklist) > 0 {
		items := worklist
		worklist = nil
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				worklist = append(worklist, f(item)...)
			}
		}
	}
}

func crawl(url string) []string {
	fmt.Println(url)
	list, err := extract(url)
	if err != nil {
		log.Print(err)
	}
	return list
}

func extract(url string) ([]string, error) {
	resp, err := http.Get(url)

	if err != nil {
		return nil, err
	}
	originalHost := resp.Request.URL.Host
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("getting %s: %s", url, resp.Status)
	}

	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		return nil, fmt.Errorf("parsing %s as HTML: %v", url, err)
	}

	var links []string
	visitNode := func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, a := range n.Attr {
				if a.Key != "href" {
					continue
				}
				link, err := resp.Request.URL.Parse(a.Val)
				if err != nil {
					continue // ignore bad URLs
				}
				if link.Host != originalHost {
					continue // skip links outside of domain
				}
				linkString := link.String()
				linkParts := strings.Split(linkString, "#")
				if len(linkParts) > 1 {
					linkString = strings.Join(linkParts[:len(linkParts)-1], "#") // remove anchor tag
				}
				j := len(linkString)
				_, size := utf8.DecodeLastRuneInString(linkString[:j])
				j -= size

				lastByRune := linkString[j:]
				if lastByRune == "/" {
					linkString = linkString[:j]
				}
				fileExtension := filepath.Ext(linkString)
				if len(fileExtension) > 0 && fileExtension != ".html" {
					continue //skip downloadable links
				}
				links = append(links, linkString)
			}
		}
	}
	forEachNode(doc, visitNode, nil)
	fmt.Println(doc.Data)
	return links, nil
}

func forEachNode(n *html.Node, pre, post func(n *html.Node)) {
	if pre != nil {
		pre(n)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, post)
	}
	if post != nil {
		post(n)
	}
}

func main() {
	// Crawl the web breadth-first,
	// starting from the command-line arguments.
	breadthFirst(crawl, os.Args[1:])
}
