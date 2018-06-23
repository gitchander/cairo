package main

import (
	"log"

	"github.com/gitchander/cairo"
	"github.com/gitchander/cairo/colorf"
)

func main() {

	surface, err := cairo.NewSurface(cairo.FORMAT_ARGB32, 512, 512)
	if err != nil {
		log.Fatal(err)
	}
	defer surface.Destroy()

	var (
		width  = surface.GetWidth()
		height = surface.GetHeight()
		stride = surface.GetStride()
	)

	n := surface.GetDataLength()
	bs := make([]byte, n)

	if err = surface.GetData(bs); err != nil {
		log.Fatal(err)
	}

	renderMSet(bs, width, height, stride, colorf.RGB{})

	if err = surface.SetData(bs); err != nil {
		log.Fatal(err)
	}

	if err = surface.WriteToPNG("fractal.png"); err != nil {
		log.Fatal(err)
	}
}

func renderMSet(bs []byte, width, height, stride int, c colorf.RGB) {

	var (
		dx = 4.0 / float64(width)
		dy = 4.0 / float64(height)
	)

	var clBackground, clForeground, clResult colorf.RGBA
	cR, cG, cB := c.GetRGB()

	coder := colorf.NewCoderBGRA32()

	n := 200

	y := -2.0
	for pY := 0; pY < height; pY++ {
		x := -2.0
		for pX := 0; pX < width; pX++ {

			cA := calcAlphaSubpixel3x3(x, y, dx, dy, n)

			clForeground = colorf.RGBA{
				R: cR,
				G: cG,
				B: cB,
				A: cA,
			}

			i := pX * 4
			clBackground, _ = coder.Decode(bs[i:])
			clResult = clForeground.Over(clBackground)
			coder.Encode(bs[i:], clResult)

			x += dx
		}
		bs = bs[stride:]
		y += dy
	}
}

var subpixelShifts3x3 = []float64{
	-1.0 / 3.0,
	0.0,
	+1.0 / 3.0,
}

func calcAlphaSubpixel3x3(x0, y0 float64, dx, dy float64, n int) float64 {

	shift := subpixelShifts3x3
	m := len(shift)

	count := 0
	for iX := 0; iX < m; iX++ {
		for iY := 0; iY < m; iY++ {

			x := x0 + dx*shift[iX]
			y := y0 + dy*shift[iY]

			i := MandelbrotSet(x, y, n)
			if i >= n {
				count++
			}
		}
	}

	alpha := float64(count) / float64(m*m)

	return alpha
}

var subpixelShifts4x4 = []float64{
	-3.0 / 8.0,
	-1.0 / 8.0,
	+1.0 / 8.0,
	+3.0 / 8.0,
}

func calcAlphaSubpixel4x4(x0, y0 float64, dx, dy float64, n int) float64 {

	shift := subpixelShifts4x4
	m := len(shift)

	count := 0
	for iX := 0; iX < m; iX++ {
		for iY := 0; iY < m; iY++ {

			x := x0 + dx*shift[iX]
			y := y0 + dy*shift[iY]

			i := MandelbrotSet(x, y, n)
			if i >= n {
				count++
			}
		}
	}

	alpha := float64(count) / float64(m*m)

	return alpha
}
