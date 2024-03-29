package main

import (
	"image/color"
	"log"

	"github.com/gitchander/cairo"
	"github.com/gitchander/cairo/colorf"
	. "github.com/gitchander/cairo/examples/mathf"
)

func main() {
	checkError(run())
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func run() error {
	surface, err := cairo.NewSurface(cairo.FORMAT_ARGB32, 512, 512)
	if err != nil {
		return err
	}
	defer surface.Destroy()

	var bg, fg color.Color

	//------------------------------------------------
	// bg = colorf.MustParseColor("#0f07")
	// fg = colorf.MustParseColor("#f007")

	bg = color.NRGBA{G: 255, A: 227}
	fg = colorf.NRGBAf{R: 1.0, G: 0.0, B: 0.0, A: 0.6}

	// bg = colorf.MustParseColor("#ffff")
	// fg = colorf.MustParseColor("#000f")
	//------------------------------------------------

	c, err := cairo.NewCanvas(surface)
	if err != nil {
		return err
	}
	c.FillColor(bg)

	var (
		width  = surface.GetWidth()
		height = surface.GetHeight()
		stride = surface.GetStride()
	)

	n := surface.GetDataLength()
	bs := make([]byte, n)

	err = surface.GetData(bs)
	if err != nil {
		return err
	}

	renderMSet(bs, width, height, stride, fg)

	err = surface.SetData(bs)
	if err != nil {
		return err
	}

	return surface.WriteToPNG("fractal.png")
}

func renderMSet(bs []byte, width, height, stride int, c color.Color) error {

	var (
		dx = 4.0 / float64(width)
		dy = 4.0 / float64(height)
	)

	cf := colorf.NRGBAfModel.Convert(c).(colorf.NRGBAf)

	coder := colorf.CoderBGRA32

	n := 200

	orbit := makeOrbitFunctor(4, n)

	y := -2.0
	for pY := 0; pY < height; pY++ {
		x := -2.0
		for pX := 0; pX < width; pX++ {

			i := pX * 4
			clBackground, err := coder.Decode(bs[i:])
			if err != nil {
				return err
			}

			var (
				//factorA = calcAlphaSubpixel3x3(x, y, dx, dy, orbit)
				factorA = calcAlphaSubpixel4x4(x, y, dx, dy, orbit)
			)
			clForeground := colorf.NRGBAf{
				R: cf.R,
				G: cf.G,
				B: cf.B,
				A: cf.A * factorA,
			}

			clResult := colorf.ColorOver(clBackground, clForeground)

			coder.Encode(bs[i:], clResult)

			x += dx
		}
		bs = bs[stride:]
		y += dy
	}
	return nil
}

var subpixelShifts3x3 = []float64{
	-1.0 / 3.0,
	0.0,
	+1.0 / 3.0,
}

func calcAlphaSubpixel3x3(x0, y0 float64, dx, dy float64, orbit OrbitFunctor) float64 {

	shift := subpixelShifts3x3
	m := len(shift)

	count := 0
	for iX := 0; iX < m; iX++ {
		for iY := 0; iY < m; iY++ {
			z := Complex{
				Re: x0 + dx*shift[iX],
				Im: y0 + dy*shift[iY],
			}
			if orbit(z) {
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

func calcAlphaSubpixel4x4(x0, y0 float64, dx, dy float64, orbit OrbitFunctor) float64 {

	shift := subpixelShifts4x4
	m := len(shift)

	count := 0
	for iX := 0; iX < m; iX++ {
		for iY := 0; iY < m; iY++ {
			z := Complex{
				Re: x0 + dx*shift[iX],
				Im: y0 + dy*shift[iY],
			}
			if orbit(z) {
				count++
			}
		}
	}

	alpha := float64(count) / float64(m*m)

	return alpha
}

type OrbitFunctor func(z Complex) bool

func makeOrbitFunctor(k int, n int) OrbitFunctor {
	switch k {
	case 0:
		return func(z Complex) bool {
			_, ok := MandelbrotSet(z, n)
			return ok
		}
	case 1:
		return func(z Complex) bool {
			_, ok := MandelbrotSetPow3(z, n)
			return ok
		}
	case 2:
		return func(z Complex) bool {
			_, ok := MandelbrotSetPowM(z, 10, n)
			return ok
		}
	case 3:
		return func(z Complex) bool {
			_, ok := MandelbrotSetPow(z, 2.5, n)
			return ok
		}
	case 4:
		c := Complex{Re: -0.7269, Im: 0.1889}
		return func(z Complex) bool {
			_, ok := JuliaSet(c, z, n)
			return ok
		}
	default:
		return func(z Complex) bool {
			return false
		}
	}
}
