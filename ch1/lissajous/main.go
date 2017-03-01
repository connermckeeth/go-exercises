// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// Run with "web" command-line argument for web server.
// See page 13.

// Lissajous generates GIF animations of random Lissajous figures.
package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
	"os"
)

var palette = []color.Color{color.Black, color.RGBA{0, 255, 0, 255}}

const (
	blackIndex = 0 // first color in palette
	greenIndex = 1 // next color in palette
)

func main() {
	lissajous(os.Stdout)
}

func lissajous(out io.Writer) {
	const (
		cycles  = 3     // number of complete x oscillator revolutions
		res     = 0.0001 // angular resolution
		size    = 120   // image canvas covers [-size..+size]
		nframes = 256    // number of animation frames
		delay   = 5     // delay between frames in 10ms units
	)
	freq := rand.Float64() * 3.0 // relative frequency of y oscillator
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0 // phase difference
	var (
		rl uint8 = 0
		gl uint8 = 255
		bl uint8 = 0
		al uint8 = 255
	)
	var (
		rb uint8 = 0
		gb uint8 = 0
		bb uint8 = 0
		ab uint8 = 0
	)
	// var pickedIndex uint8
	var lineColor = color.RGBA{rl,gl,bl,al}
	var bgColor = color.RGBA{rb,gb,bb,ab}
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		//if i%2 == 0 {
			lineColor = myChange(lineColor)
		/*} else {
			bgColor = myChange(bgColor)
		}*/
		palette = []color.Color{bgColor, lineColor}
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), greenIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim) // NOTE: ignoring encoding errors
}

func myChange(color color.RGBA) color.RGBA {
	var r float64 = float64(color.R)
	var g float64 = float64(color.G)
	var b float64 = float64(color.B)
	var a float64 = float64(color.A)	

	r = (r+g-b)/3
	color.R = uint8(r)
	g = (g+b-r)/3
	color.G = uint8(g)
	b = (b+r-g)/3
	color.B = uint8(b)
	color.A = uint8(a)
	return color
}
