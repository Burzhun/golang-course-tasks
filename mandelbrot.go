package main

import (
	"image"
	"image/color"
	"image/png"
	"math/cmplx"
	"log"
	"net/http"
)
	
func main() {
	
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000",nil))
	
}

func handler(w http.ResponseWriter, r *http.Request) {
	const (
		xmin, ymin, xmax, ymax = -2, -2, 2, 2
		width, height = 3024, 3024
	)
	img := image.NewRGBA(image.Rect(0,0,width, height))
	for py := 0; py < height; py++ {
		y := float64(py) / height * (ymax - ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin
			z := complex(x, y)
			img.Set(px, py, mandelbrot(z))
		}
	}
	png.Encode(w, img)
}


func mandelbrot(z complex128) color.Color {
	const iterations = 200
	const contrast = 15
	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {	
			return color.Gray{255 - contrast*n}
		}
	}
	return color.Black
}
	