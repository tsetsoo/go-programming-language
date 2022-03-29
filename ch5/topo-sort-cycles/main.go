package main

import (
	"fmt"
)

var prereqs = map[string][]string{
	"algorithms":     {"data structures"},
	"calculus":       {"linear algebra"},
	"linear algebra": {"calculus"},

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

func main() {
	for i, course := range topoSort(prereqs) {
		fmt.Printf("%d:\t%s\n", i+1, course)
	}
}

func topoSort(m map[string][]string) []string {
	var order []string
	seen := make(map[string]bool)
	var visitAll func(items string, currentCycleDetection map[string]bool)

	visitAll = func(item string, currentCycleDetection map[string]bool) {
		if currentCycleDetection[item] {
			fmt.Printf("cycle detected with: %s\n", item)
			return
		}
		if seen[item] {
			return
		}

		seen[item] = true
		currentCycleDetection[item] = true

		items, ok := m[item]
		if ok {
			for _, itemToIter := range items {
				visitAll(itemToIter, currentCycleDetection)
			}
		}
		delete(currentCycleDetection, item)
		order = append(order, item)
	}

	for key := range m {
		currentCycleDetection := make(map[string]bool)
		visitAll(key, currentCycleDetection)
	}

	return order
}
