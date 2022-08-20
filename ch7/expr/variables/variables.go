package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"tsvetelinpantev.com/go-programming-language/ch7/expr/eval"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	expr, err := eval.Parse(scanner.Text())
	if err != nil {
		fmt.Printf("Error while parsing expression: %s, %v\n", scanner.Text(), err)
		return
	}
	variables := make(map[eval.Var]bool)
	if err := expr.Check(variables); err != nil {
		fmt.Printf("Error while verifing expression: %s, %v\n", scanner.Text(), err)
		return
	}
	fmt.Println(variables)
	envs := eval.Env{}
	for v := range variables {
		fmt.Println("Enter value for: " + v)
		scanner.Scan()
		floatValue, err := strconv.ParseFloat(scanner.Text(), 64)
		if err != nil {
			fmt.Printf("Could not parse: '%s' as float\n", scanner.Text())
		}
		envs[eval.Var(v)] = floatValue
	}

	fmt.Println(expr.Eval(envs))
}
