package colorf

import (
	"image"
	"image/color"
	"math"
)

const maxColorComponent = 0xffff

// NColorf represents a non-alpha-premultiplied color
type NColorf struct {
	R, G, B, A float64
}

var _ color.Color = NColorf{}

func _() {
	image.Pt(0, 0)
}

// NClrf is shorthand for NColorf{R, G, B, A}.
func NClrf(r, g, b, a float64) NColorf {
	return NColorf{
		R: r,
		G: g,
		B: b,
		A: a,
	}.clamp()
}

func (c NColorf) clamp() NColorf {
	return NColorf{
		R: clamp01(c.R),
		G: clamp01(c.G),
		B: clamp01(c.B),
		A: clamp01(c.A),
	}
}

func (c NColorf) RGBA() (r, g, b, a uint32) {
	// return c.v1_RGBA()
	// return c.v2_RGBA()
	return c.v3_RGBA()
}

func (c NColorf) v1_RGBA() (r, g, b, a uint32) {

	cc := c.clamp()

	u := color.NRGBA64{
		R: uint16(colorComponentConvert(cc.R)),
		G: uint16(colorComponentConvert(cc.G)),
		B: uint16(colorComponentConvert(cc.B)),
		A: uint16(colorComponentConvert(cc.A)),
	}

	return u.RGBA()
}

func (c NColorf) v2_RGBA() (r, g, b, a uint32) {

	cc := c.clamp()

	// alpha-premultiple
	{
		cc.R *= cc.A
		cc.G *= cc.A
		cc.B *= cc.A
	}

	r = colorComponentConvert(cc.R)
	g = colorComponentConvert(cc.G)
	b = colorComponentConvert(cc.B)
	a = colorComponentConvert(cc.A)

	return
}

func (c NColorf) v3_RGBA() (r, g, b, a uint32) {

	cc := c.clamp()

	r = colorComponentConvert(cc.R)
	g = colorComponentConvert(cc.G)
	b = colorComponentConvert(cc.B)
	a = colorComponentConvert(cc.A)

	// alpha-premultiple
	{
		r = (r * a) / maxColorComponent
		g = (g * a) / maxColorComponent
		b = (b * a) / maxColorComponent
	}

	return
}

func colorComponentConvert(v float64) uint32 {
	return uint32(math.Round(v * maxColorComponent))
}

// ------------------------------------------------------------------------------
func ncolorfModel(c color.Color) color.Color {
	if _, ok := c.(NColorf); ok {
		return c
	}
	r, g, b, a := c.RGBA()
	if a == 0 {
		return NColorf{0, 0, 0, 0}
	}
	if a == maxColorComponent {
		return NColorf{
			R: float64(r) / maxColorComponent,
			G: float64(g) / maxColorComponent,
			B: float64(b) / maxColorComponent,
			A: 1,
		}
	}
	return NColorf{
		R: float64(r) / float64(a),
		G: float64(g) / float64(a),
		B: float64(b) / float64(a),
		A: float64(a) / maxColorComponent,
	}
}

var NColorfModel color.Model = color.ModelFunc(ncolorfModel)
