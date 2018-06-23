package colorf

import "image/color"

type RGBA struct {
	R, G, B, A float64
}

var _ color.Color = RGBA{}

func (c RGBA) RGBA() (r, g, b, a uint32) {
	r = channelToMax(c.R)
	g = channelToMax(c.G)
	b = channelToMax(c.B)
	a = channelToMax(c.A)
	return
}

func (c RGBA) Norm() RGBA {
	return RGBA{
		R: norm(c.R),
		G: norm(c.G),
		B: norm(c.B),
		A: norm(c.A),
	}
}

// Alpha blending
// a over b
func (a RGBA) Over(b RGBA) RGBA {
	A := lerp(b.A, 1.0, a.A)
	return RGBA{
		R: lerp(b.R*b.A, a.R, a.A) / A,
		G: lerp(b.G*b.A, a.G, a.A) / A,
		B: lerp(b.B*b.A, a.B, a.A) / A,
		A: A,
	}
}

var RGBAModel color.Model = color.ModelFunc(rgbaModel)

func rgbaModel(c color.Color) color.Color {
	if _, ok := c.(RGBA); ok {
		return c
	}
	r, g, b, a := c.RGBA()
	return RGBA{
		R: float64(r) / channelMax,
		G: float64(g) / channelMax,
		B: float64(b) / channelMax,
		A: float64(a) / channelMax,
	}
}
