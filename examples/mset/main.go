package main

import (
	"image/color"
	"log"

	"github.com/gitchander/cairo"
	"github.com/gitchander/cairo/colorf"
)

func main() {

	surface, err := cairo.NewSurface(cairo.FORMAT_ARGB32, 512, 512)
	checkError(err)
	defer surface.Destroy()

	var (
		width  = surface.GetWidth()
		height = surface.GetHeight()
		stride = surface.GetStride()
	)

	n := surface.GetDataLength()
	bs := make([]byte, n)

	err = surface.GetData(bs)
	checkError(err)

	renderMSet(bs, width, height, stride, color.Black)

	err = surface.SetData(bs)
	checkError(err)

	err = surface.WriteToPNG("fractal.png")
	checkError(err)
}

func renderMSet(bs []byte, width, height, stride int, c color.Color) {

	var (
		dx = 4.0 / float64(width)
		dy = 4.0 / float64(height)
	)

	cf := colorf.RGBModel.Convert(c).(colorf.RGB)

	coder := colorf.CoderBGRA32

	n := 200

	y := -2.0
	for pY := 0; pY < height; pY++ {
		x := -2.0
		for pX := 0; pX < width; pX++ {

			cA := calcAlphaSubpixel3x3(x, y, dx, dy, n)

			clForeground := colorf.RGBA{
				R: cf.R,
				G: cf.G,
				B: cf.B,
				A: cA,
			}

			i := pX * 4
			clBackground, err := coder.Decode(bs[i:])
			checkError(err)

			clBackgroundf := colorf.RGBAModel.Convert(clBackground).(colorf.RGBA)

			clResult := clForeground.Over(clBackgroundf)
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
			if i == -1 {
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
			if i == -1 {
				count++
			}
		}
	}

	alpha := float64(count) / float64(m*m)

	return alpha
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
