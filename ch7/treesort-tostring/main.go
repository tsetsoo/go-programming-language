package main

import (
	"fmt"
	"strings"
)

type tree struct {
	value       int
	left, right *tree
}

func main() {
	tree := &tree{
		value: 2,
		left: &tree{
			value: 0,
			left: &tree{
				value: -1,
			},
			right: &tree{
				value: 1,
			},
		},
		right: &tree{
			value: 4,
			left: &tree{
				value: 3,
			},
			right: &tree{
				value: 5,
			},
		},
	}
	fmt.Println(tree)
}

func (t *tree) String() string {
	var sb strings.Builder
	addValue(t, &sb)
	return sb.String()
}

func addValue(tr *tree, sb *strings.Builder) {
	if tr == nil {
		return
	}
	addValue(tr.left, sb)
	sb.WriteString(fmt.Sprintf("%d ", tr.value))
	addValue(tr.right, sb)
}
