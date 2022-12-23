package colorf

import (
	"image/color"
)

// RGBAf represents a alpha-premultiplied color.
type RGBAf struct {
	R, G, B, A float64
}

var _ color.Color = RGBAf{}

func (c RGBAf) RGBA() (r, g, b, a uint32) {
	r = colorComponentEncode(c.R)
	g = colorComponentEncode(c.G)
	b = colorComponentEncode(c.B)
	a = colorComponentEncode(c.A)
	return
}

func (c RGBAf) toNRGBAf() NRGBAf {
	return NRGBAf{
		R: c.R / c.A,
		G: c.G / c.A,
		B: c.B / c.A,
		A: c.A,
	}
}

func rgbafModel(c color.Color) color.Color {
	if _, ok := c.(RGBAf); ok {
		return c
	}
	if sa, ok := c.(NRGBAf); ok {
		return sa.toRGBAf()
	}
	r, g, b, a := c.RGBA()
	if a == 0 {
		return RGBAf{0, 0, 0, 0}
	}
	return RGBAf{
		R: colorComponentDecode(r),
		G: colorComponentDecode(g),
		B: colorComponentDecode(b),
		A: colorComponentDecode(a),
	}
}

var RGBAfModel color.Model = color.ModelFunc(rgbafModel)
