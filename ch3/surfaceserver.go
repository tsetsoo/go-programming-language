package main

import (
	"fmt"
	"io"
	"log"
	"math"
	"net/http"
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

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "image/svg+xml")
		svgFunction(w)
	})
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func svgFunction(w io.Writer) {
	fmt.Fprintf(w, "<svg xmlns='http://www.w3.org/2000/svg' style='stroke: grey; fill: white; stroke-width: 0.7' width='%d' height='%d'>", width, height)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay, _ := corner(i+1, j)
			bx, by, color := corner(i, j)
			cx, cy, _ := corner(i, j+1)
			dx, dy, _ := corner(i+1, j+1)
			if noInfinities([]float64{ax, ay, bx, by, cx, cy, dx, dy}) {
				fmt.Fprintf(w, "<polygon points='%g,%g %g,%g %g,%g %g,%g' style='fill: %s'/>\n", ax, ay, bx, by, cx, cy, dx, dy, color)
			}
		}
	}
	fmt.Fprintln(w, "</svg>")
}

func corner(i, j int) (float64, float64, string) {
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	z := f(x, y)
	color := "blue"
	if z > 0 {
		color = "red"
	}

	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy, color
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
	r := math.Hypot(x, y)
	return math.Sin(r) / r
	// return 0.1 * (math.Sin(x/1.5) + math.Sin(y/1.5)) // egg box
	// return x * x - y * y // saddle
	// return 0.015 * y * math.Sin(x) // moguls
	// return 0.1 * math.Sin(x) * math.Sin(y) // moguls
}
