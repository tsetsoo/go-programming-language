// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 203.

// The surface program plots the 3-D surface of a user-provided function.
package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"tsvetelinpantev.com/go-programming-language/ch7/expr/eval"
)

func main() {
	http.HandleFunc("/calc", calc)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func calc(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	expr, err := parseAndCheck(r.Form.Get("expr"))
	if err != nil {
		http.Error(w, "bad expr: "+err.Error(), http.StatusBadRequest)
		return
	}
	env, err := parseEnv(r.Form.Get("env"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	fmt.Fprintln(w, expr.Eval(env))
}

func parseAndCheck(s string) (eval.Expr, error) {
	if s == "" {
		return nil, fmt.Errorf("empty expression")
	}
	expr, err := eval.Parse(s)
	if err != nil {
		return nil, err
	}
	vars := make(map[eval.Var]bool)
	if err := expr.Check(vars); err != nil {
		return nil, err
	}
	return expr, nil
}

func parseEnv(s string) (eval.Env, error) {
	env := eval.Env{}
	assignments := strings.Fields(s)
	for _, a := range assignments {
		fields := strings.Split(a, "=")
		if len(fields) != 2 {
			return env, fmt.Errorf("bad assignment: %s\n", a)
		}
		ident, valStr := fields[0], fields[1]
		val, err := strconv.ParseFloat(valStr, 64)
		if err != nil {
			return env, fmt.Errorf("bad value for %s: %s\n", ident, err)
		}
		env[eval.Var(ident)] = val
	}
	return env, nil
}
