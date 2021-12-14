package main

import (
	"image"
	"image/color"
	"image/png"
	"io"
	"log"
	"math/cmplx"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		mandelbrotImage(w)
	})
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func mandelbrotImage(w io.Writer) {
	const (
		xmin, ymin, xmax, ymax = -2, -2, +2, +2
		width, height          = 8192, 8192
		// supersampling          = 2
	)
	// var colors [width * supersampling][height * supersampling]color.Color

	// for py := 0; py < height*supersampling; py++ {
	// 	y := float64(py)/height*(ymax-ymin) + ymin
	// 	for px := 0; px < width*supersampling; px++ {
	// 		x := float64(px)/width*(xmax-xmin) + xmin
	// 		z := complex(x, y)
	// 		colors[px][py] = mandelbrot(z)
	// 	}
	// }
	// for i := 0; i < width; i++ {
	// 	for j := 0; j < height; j++ {
	// 		colorI, colorJ := i*supersampling, j*supersampling
	// 		pr1, pg1, pb1, a1 := colors[colorI][colorJ].RGBA()
	// 		pr2, pg2, pb2, a2 := colors[colorI][colorJ+1].RGBA()
	// 		pr3, pg3, pb3, a3 := colors[colorI+1][colorJ].RGBA()
	// 		pr4, pg4, pb4, a4 := colors[colorI+1][colorJ+1].RGBA()

	// 		superSampledColor := color.RGBA{uint8((pr1 + pr2 + pr3 + pr4) / (supersampling * supersampling)), uint8((pg1 + pg2 + pg3 + pg4) / (supersampling * supersampling)), uint8((pb1 + pb2 + pb3 + pb4) / (supersampling * supersampling)), uint8((a1 + a2 + a3 + a4) / (supersampling * supersampling))}
	// 		img.Set(j, i, superSampledColor)

	// 	}
	// }
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin
			z := complex(x, y)
			// Image point (px, py) represents complex value z.
			img.Set(px, py, mandelbrot(z))
		}
	}
	png.Encode(w, img)
}

func mandelbrot(z complex128) color.Color {
	const iterations = 37
	const contrast = 7
	for n := uint8(0); n < iterations; n++ {
		z -= (z*z*z*z - 1) / (4 * z * z * z)
		if cmplx.Abs(z*z*z*z-1) < 1e-6 {
			// return color.RGBA{contrast*n, 0, 0, 10}
			return color.Gray{255 - n*contrast}
		}
	}
	return color.Gray{0}
}

// func newton(z complex128) color.Color {

// }
