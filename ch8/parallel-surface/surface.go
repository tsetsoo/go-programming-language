package main

import (
	"math"
)

const (
	width, height = 600, 320
	cells         = 100
	xyrange       = 30.0
	xyscale       = width / 2 / xyrange
	zscale        = height * 0.4
	angle         = math.Pi / 6
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle)

func surface() {
	//fmt.Printf("<svg xmlns='http://www.w3.org/2000/svg' style='stroke: grey; fill: white; stroke-width: 0.7' width='%d' height='%d'>", width, height)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay := corner(i+1, j)
			bx, by := corner(i, j)
			cx, cy := corner(i, j+1)
			dx, dy := corner(i+1, j+1)
			if noInfinities([]float64{ax, ay, bx, by, cx, cy, dx, dy}) {
				//fmt.Printf("<polygon points='%g,%g %g,%g %g,%g %g,%g'/>\n", ax, ay, bx, by, cx, cy, dx, dy)
			}
		}
	}
	// fmt.Println("</svg>")S
}

func parallelSurface() {
	waitChan := make(chan struct{}, cells)

	//fmt.Printf("<svg xmlns='http://www.w3.org/2000/svg' style='stroke: grey; fill: white; stroke-width: 0.7' width='%d' height='%d'>", width, height)
	for i := 0; i < cells; i++ {
		waitChan <- struct{}{}
		go func(i int) {
			for j := 0; j < cells; j++ {
				ax, ay := corner(i+1, j)
				bx, by := corner(i, j)
				cx, cy := corner(i, j+1)
				dx, dy := corner(i+1, j+1)
				if noInfinities([]float64{ax, ay, bx, by, cx, cy, dx, dy}) {
					//	fmt.Printf("<polygon points='%g,%g %g,%g %g,%g %g,%g'/>\n", ax, ay, bx, by, cx, cy, dx, dy)
				}
			}
			<-waitChan
		}(i)
	}
	//fmt.Println("</svg>")
}

func parallelJSurface() {
	waitChan := make(chan struct{}, cells)

	//fmt.Printf("<svg xmlns='http://www.w3.org/2000/svg' style='stroke: grey; fill: white; stroke-width: 0.7' width='%d' height='%d'>", width, height)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			waitChan <- struct{}{}
			go func(j int) {
				ax, ay := corner(i+1, j)
				bx, by := corner(i, j)
				cx, cy := corner(i, j+1)
				dx, dy := corner(i+1, j+1)
				if noInfinities([]float64{ax, ay, bx, by, cx, cy, dx, dy}) {
					//	fmt.Printf("<polygon points='%g,%g %g,%g %g,%g %g,%g'/>\n", ax, ay, bx, by, cx, cy, dx, dy)
				}
				<-waitChan
			}(i)
		}
	}
	//fmt.Println("</svg>")
}

func corner(i, j int) (sx float64, sy float64) {
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	z := f(x, y)

	sx = width/2 + (x-y)*cos30*xyscale
	sy = height/2 + (x+y)*sin30*xyscale - z*zscale
	return
}

func noInfinities(numbers []float64) bool {
	for _, number := range numbers {
		if math.IsInf(number, 0) || math.IsNaN(number) {
			return false
		}
	}
	return true
}

func f(x, y float64) float64 {
	// r := math.Hypot(x, y)
	// return math.Sin(r) / r
	// return -(y+47)*math.Sin(math.Sqrt(math.Abs(y+x/2+47))) - x*math.Sin(math.Sqrt(math.Abs(x-(y+47))))
	return x*x - y*y
}
