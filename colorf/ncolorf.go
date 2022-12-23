package colorf

import (
	"image/color"
)

// NRGBAf represents a non-alpha-premultiplied color.
type NRGBAf struct {
	R, G, B, A float64
}

var _ color.Color = NRGBAf{}

func (c NRGBAf) RGBA() (r, g, b, a uint32) {

	// v := c.toRGBAf()
	// return v.RGBA()

	r = colorComponentEncode(c.R * c.A)
	g = colorComponentEncode(c.G * c.A)
	b = colorComponentEncode(c.B * c.A)
	a = colorComponentEncode(c.A)

	return
}

func (c NRGBAf) toRGBAf() RGBAf {
	return RGBAf{
		R: c.R * c.A,
		G: c.G * c.A,
		B: c.B * c.A,
		A: c.A,
	}
}

func nrgbafModel(c color.Color) color.Color {

	if _, ok := c.(NRGBAf); ok {
		return c
	}

	if pa, ok := c.(RGBAf); ok {
		return pa.toNRGBAf()
	}

	r, g, b, a := c.RGBA()
	if a == 0 {
		return NRGBAf{0, 0, 0, 0}
	}

	// var (
	// 	fR = colorComponentDecode(r)
	// 	fG = colorComponentDecode(g)
	// 	fB = colorComponentDecode(b)
	// 	fA = colorComponentDecode(a)
	// )

	// return NRGBAf{
	// 	R: fR / fA,
	// 	G: fG / fA,
	// 	B: fB / fA,
	// 	A: fA,
	// }

	if a == maxColorComponent {
		return NRGBAf{
			R: float64(r) / maxColorComponent,
			G: float64(g) / maxColorComponent,
			B: float64(b) / maxColorComponent,
			A: 1,
		}
	}
	return NRGBAf{
		R: float64(r) / float64(a),
		G: float64(g) / float64(a),
		B: float64(b) / float64(a),
		A: float64(a) / maxColorComponent,
	}
}

var NRGBAfModel color.Model = color.ModelFunc(nrgbafModel)
