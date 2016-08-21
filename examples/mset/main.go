package main

import (
	"fmt"

	"github.com/gitchander/cairo"
	"github.com/gitchander/cairo/color"
)

func main() {
	if err := ExampleMSet(); err != nil {
		fmt.Println(err.Error())
	}
}

func ExampleMSet() error {

	surface, err := cairo.NewSurface(cairo.FORMAT_ARGB32, 512, 512)
	if err != nil {
		return err
	}
	defer surface.Destroy()

	var (
		width  = surface.GetWidth()
		height = surface.GetHeight()
		stride = surface.GetStride()
	)

	n := surface.GetDataLength()
	bs := make([]byte, n)

	//color.SurfaceFillRGB(surface, color.NewRGB(1, 1, 1))

	if err = surface.GetData(bs); err != nil {
		return err
	}

	renderMSet(bs, width, height, stride, color.NewRGB(0, 0, 0))

	if err = surface.SetData(bs); err != nil {
		return err
	}

	if err = surface.WriteToPNG("fractal.png"); err != nil {
		return err
	}

	return nil
}

func renderMSet(bs []byte, width, height, stride int, c color.RGB) {

	var (
		dx = 4.0 / float64(width)
		dy = 4.0 / float64(height)
	)

	var clBackground, clForeground, clResult color.RGBA
	cR, cG, cB := c.GetRGB()

	coder := color.NewCoderBGRA32()

	n := 200

	y := -2.0
	for pY := 0; pY < height; pY++ {
		x := -2.0
		for pX := 0; pX < width; pX++ {

			cA := calcAlphaSubpixel3x3(x, y, dx, dy, n)

			clForeground = color.NewRGBA(cR, cG, cB, cA)

			i := pX * 4
			clBackground, _ = coder.Decode(bs[i:])
			clResult = color.Over(clForeground, clBackground)
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
