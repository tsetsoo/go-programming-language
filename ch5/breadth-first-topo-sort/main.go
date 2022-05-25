// Findlinks4 crawls the web, starting with the URLs on the command line.
package main

import "fmt"

var prereqs = map[string][]string{
	"algorithms": {"data structures"},
	"calculus":   {"linear algebra"},

	"compilers": {
		"data structures",
		"formal languages",
		"computer organization",
	},

	"data structures":       {"discrete math"},
	"databases":             {"data structures"},
	"discrete math":         {"intro to programming"},
	"formal languages":      {"discrete math"},
	"networks":              {"operating systems"},
	"operating systems":     {"data structures", "computer organization"},
	"programming languages": {"data structures", "computer organization"},
}

var seen = make(map[string]bool)

func breadthFirst(f func(item string) []string, worklist []string) {
	levelDeep := 0
	for len(worklist) > 0 {

		items := worklist
		worklist = nil
		for _, item := range items {
			if !seen[item] {
				fmt.Printf("%d. ", levelDeep)

				seen[item] = true
				worklist = append(worklist, f(item)...)
			}
		}
		levelDeep++
	}
}

func crawl(url string) []string {
	fmt.Println(url)
	return prereqs[url]
}

func main() {
	for prereq := range prereqs {
		breadthFirst(crawl, []string{prereq})
	}

}
