package main

import (
	"encoding/xml"
	"fmt"
	"io"
	"os"
	"strings"
)

type element struct {
	name       string
	attributes map[string]struct{}
}

type elements []element

func main() {
	dec := xml.NewDecoder(os.Stdin)
	var stack elements // stack of element names
	for {
		tok, err := dec.Token()
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Fprintf(os.Stderr, "xmlselect: %v\n", err)
			os.Exit(1)
		}
		switch tok := tok.(type) {
		case xml.StartElement:
			element := element{name: tok.Name.Local, attributes: make(map[string]struct{})}
			for _, val := range tok.Attr {
				element.attributes[val.Value] = struct{}{}
			}
			stack = append(stack, element) // push
		case xml.EndElement:
			stack = stack[:len(stack)-1] // pop
		case xml.CharData:
			if containsAll(stack, os.Args[1:]) {
				fmt.Printf("%s: %s\n", stack.String(), tok)
			}
		}
	}
}

func (e elements) String() string {
	toReturn := ""
	for _, e := range e {
		v := make([]string, 0, len(e.attributes))

		for value := range e.attributes {
			v = append(v, value)
		}
		toReturn += fmt.Sprintf("%s(attributes: %s) ", e.name, strings.Join(v, ","))
	}
	return toReturn
}

// containsAll reports whether x contains the elements of y, in order.
func containsAll(x elements, y []string) bool {
	for len(y) <= len(x) {
		if len(y) == 0 {
			return true
		}
		elementAtXZero := x[0]
		_, ok := elementAtXZero.attributes[y[0]]
		if elementAtXZero.name == y[0] || ok {
			y = y[1:]
		}
		x = x[1:]
	}
	return false
}

//!-
