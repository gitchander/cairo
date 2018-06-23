package colorf

import "image/color"

type RGB struct {
	R, G, B float64
}

var _ color.Color = RGB{}

func (c RGB) RGBA() (r, g, b, a uint32) {
	r = uint32(c.R * maxUint16)
	g = uint32(c.G * maxUint16)
	b = uint32(c.B * maxUint16)
	a = maxUint16
	return
}

func (c RGB) Norm() RGB {
	return RGB{
		R: norm(c.R),
		G: norm(c.G),
		B: norm(c.B),
	}
}

var RGBModel color.Model = color.ModelFunc(rgbModel)

func rgbModel(c color.Color) color.Color {
	if _, ok := c.(RGB); ok {
		return c
	}
	r, g, b, _ := c.RGBA()
	return RGB{
		R: float64(r) / maxUint16,
		G: float64(g) / maxUint16,
		B: float64(b) / maxUint16,
	}
}
