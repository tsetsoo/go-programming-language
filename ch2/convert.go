package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"./conv"
)

func main() {
	if len(os.Args) == 1 {
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			convertString(scanner.Text())
		}
	} else {
		for _, arg := range os.Args[1:] {
			convertString(arg)
		}
	}
}

func convertString(s string) {
	t, err := strconv.ParseFloat(s, 64)
	if err != nil {
		fmt.Fprintf(os.Stderr, "convert: %v\n", err)
		os.Exit(1)
	}
	f := conv.Fahrenheit(t)
	c := conv.Celsius(t)
	fmt.Printf("%s = %s, %s = %s\n", f, conv.FToC(f), c, conv.CToF(c))
	feet := conv.Feet(t)
	m := conv.Meter(t)
	fmt.Printf("%s = %s, %s = %s\n", feet, conv.FeetToMeter(feet), m, conv.MeterToFeet(m))
	kg := conv.Kilogram(t)
	p := conv.Pound(t)
	fmt.Printf("%s = %s, %s = %s\n", kg, conv.KgToPound(kg), p, conv.PoundToKg(p))
}
