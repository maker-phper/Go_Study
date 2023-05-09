package main

import (
	"fmt"
	"image"
	"image/color"
	"image/gif"
	"io"
	"log"
	"math"
	"math/rand"
	"net/http"
	"strconv"
)

var palette = []color.Color{color.White, color.Black}

const (
	whiteIndex = 0
	blackIndex = 1
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if err := r.ParseForm(); err != nil {
			log.Print(err)
		}

		var cycles int
		var err error

		cycleNumber, ok := r.Form["cycles"]
		if ok {
			cycles, err = strconv.Atoi(cycleNumber[0])
			if err != nil {
				fmt.Fprintf(w, "conver type fail %v", err)
			}
		}
		lissajous(w, cycles)
	})
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}
func lissajous(out io.Writer, cycleNumber int) {
	const (
		res     = 0.001 //angular resolution
		size    = 100   //image canvas covers [-size .. +size]
		nframes = 64    //number of animation frames
		delay   = 8     //delay between frames in 10ms units
	)

	var cycles int //number of complete x oscillator revolutions
	if cycleNumber != 0 {
		cycles = cycleNumber
	} else {
		cycles = 5
	}

	freq := rand.Float64() * 3.0 //relative frequency of y oscillator
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0 //phase difference
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < float64(cycles)*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), blackIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim) // Note: ignoring encode errors
}
