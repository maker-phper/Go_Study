// Mandelbrot emits a PNG image of the Mandelbrot fractal
package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"math"
	"math/cmplx"
	"os"
)

type Func func(complex128) complex128

var colorPool = []color.RGBA{
	{170, 57, 57, 255},
	{170, 108, 57, 255},
	{34, 102, 102, 255},
	{45, 136, 45, 255},
}

var chosenColors = map[complex128]color.RGBA{}

func main() {
	const (
		xmin, ymin, xmax, ymax = -2, -2, +2, +2
		width, height          = 1024, 1024
	)

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin
			z := complex(x, y)
			// Image point (px, py) represents complex value z.
			img.Set(px, py, z4(z))
		}
	}

	f1, err := os.Create("mandelbrot.png") //file in build config Working directory
	if err != nil {
		fmt.Println(err)
	}
	defer f1.Close()

	//png.Encode(os.Stdout, img) // NOTE: ignoring errors
	png.Encode(f1, img) // NOTE: ignoring errors
}

// f(x) = x^4 - 1
//
// z' = z - f(z)/f'(z)
//    = z - (z^4 - 1) / (4 * z^3)
//    = z - (z - 1/z^3) / 4
func z4(z complex128) color.Color {
	f := func(z complex128) complex128 {
		return z*z*z*z - 1
	}
	fPrime := func(z complex128) complex128 {
		return (z - 1/(z*z*z)) / 4
	}
	return newton(z, f, fPrime)
}

func mandelbrot(z complex128) color.Color {
	const iterations = 200
	const contrast = 15

	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			switch {
			case n > 50:
				return color.RGBA{100, 0, 0, 255}
			default:
				// logarithmic blue gradient to show small differences on the
				// periphery of the fractal.
				logScale := math.Log(float64(n)) / math.Log(float64(iterations))
				return color.RGBA{0, 0, 255 - uint8(logScale*255), 255}
			}
		}
	}
	return color.Black
}

func newton(z complex128, f Func, fPrime Func) color.Color {
	const iterations = 37
	const contrast = 7
	for i := uint8(0); i < iterations; i++ {
		z -= fPrime(z)
		if cmplx.Abs(f(z)) < 1e-6 {
			root := complex(round(real(z), 4), round(imag(z), 4))
			c, ok := chosenColors[root]
			if !ok {
				if len(colorPool) == 0 {
					panic("no colors left")
				}
				c = colorPool[0]
				colorPool = colorPool[1:]
				chosenColors[root] = c
			}
			// Convert to YCbCr to make producing different shades easier.
			y, cb, cr := color.RGBToYCbCr(c.R, c.G, c.B)
			scale := math.Log(float64(i)) / math.Log(iterations)
			y -= uint8(float64(y) * scale)
			return color.YCbCr{y, cb, cr}
		}
	}
	return color.Black
}

func round(f float64, digits int) float64 {
	if math.Abs(f) < 0.5 {
		return 0
	}
	pow := math.Pow10(digits)
	return math.Trunc(f*pow+math.Copysign(0.5, f)) / pow
}
