package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"fmt"
	"math"
	"math/rand"
	"time"
	"log"
	"strconv"
	"net/http"
)

var palette = []color.Color{color.Black, color.RGBA{0,150,200,1}}

const (
	whiteIndex = 0
	blackIndex = 1
)

func main() {
	
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000",nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	
	r.ParseForm()
	form := r.Form
	t := form["q"]
	if len(t)==0 {
		fmt.Println("error")
		return
	}
	fmt.Println(t[0])
	a, err:=strconv.Atoi(t[0])
	if err != nil {
		fmt.Println(err)
	}
	lissahous(w, a)
	fmt.Println(a)
}


func lissahous(out io.Writer, cycles int) {
	const (
		res = 0.001
		size = 300
		nframes = 164
		delay = 6
	)
	rand.Seed(time.Now().UTC().UnixNano())
	freq := rand.Float64()
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0,0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < math.Pi*2*float64(cycles); t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size + int(x*size + 0.5), size+int(y*size+0.5), blackIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim)
}