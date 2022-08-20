package main

import (
	"encoding/xml"
	"fmt"
	"io"
	"os"
	"strings"
)

type Node interface {
	String() string
}

type CharData string

type Element struct {
	Type     xml.Name
	Attr     []xml.Attr
	Children []Node
}

func main() {
	dec := xml.NewDecoder(strings.NewReader(`<?xml version="1.0" encoding="UTF-8" standalone="no"?><root testAttr="testValue">  <result>    <child>data1</child>    <child>A1343358848.646</child>    <child>      <internal>       <data>one</data>       <data>two</data>       <unique>Z1343358848.646</unique></internal>    </child>  </result></root>`))
	var element Element
	var stack []*Element = make([]*Element, 0)
	for {
		tok, err := dec.Token()
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Fprintf(os.Stderr, "xml-tree: %v\n", err)
			os.Exit(1)
		}
		switch tok := tok.(type) {
		case xml.StartElement:
			elementToAdd := Element{Type: tok.Name, Attr: tok.Attr, Children: make([]Node, 0)}
			if len(stack) == 0 {
				element = elementToAdd
				stack = append(stack, &element)
			} else {
				stack[len(stack)-1].Children = append(stack[len(stack)-1].Children, &elementToAdd)
				stack = append(stack, &elementToAdd)
			}
		case xml.EndElement:
			stack = stack[:len(stack)-1] // pop
		case xml.CharData:
			if len(stack) == 0 {
				continue
			}
			elementToAdd := CharData(tok)
			stack[len(stack)-1].Children = append(stack[len(stack)-1].Children, &elementToAdd)
		}
	}
	fmt.Println(element)
}

func (e Element) String() string {
	toReturn := fmt.Sprintf("<%s>\n", e.Type.Local)
	for _, child := range e.Children {
		toReturn += fmt.Sprintf("%s", child.String())
	}

	return toReturn
}

func (c CharData) String() string {
	return string(c)
}
