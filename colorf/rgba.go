package colorf

import "image/color"

type RGBA struct {
	R, G, B, A float64
}

var _ color.Color = RGBA{}

func (c RGBA) RGBA() (r, g, b, a uint32) {
	r = uint32(c.R * maxUint16)
	g = uint32(c.G * maxUint16)
	b = uint32(c.B * maxUint16)
	a = uint32(c.A * maxUint16)
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
		R: float64(r) / maxUint16,
		G: float64(g) / maxUint16,
		B: float64(b) / maxUint16,
		A: float64(a) / maxUint16,
	}
}
