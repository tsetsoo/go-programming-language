package main

import (
	"flag"
	"fmt"

	"tsvetelinpantev.com/go-programming-language/ch2/conv"
)

type celsiusFlag struct {
	conv.Celsius
}

func (f *celsiusFlag) Set(s string) error {
	var unit string
	var value float64
	fmt.Sscanf(s, "%f%s", &value, &unit)
	switch unit {
	case "C", "°C":
		f.Celsius = conv.Celsius(value)
		return nil
	case "F", "°F":
		f.Celsius = conv.FToC(conv.Fahrenheit(value))
		return nil
	case "K", "°K":
		f.Celsius = conv.KToC(conv.Kelvin(value))
		return nil
	}
	return fmt.Errorf("invalid temperature %q", s)
}

func CelsiusFlag(name string, value conv.Celsius, usage string) *conv.Celsius {
	f := celsiusFlag{value}
	flag.CommandLine.Var(&f, name, usage)
	return &f.Celsius
}

var temp = CelsiusFlag("temp", 20.0, "the temperature")

func main() {
	flag.Parse()
	fmt.Println(*temp)
}
