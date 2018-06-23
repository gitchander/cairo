package colorf

import "image/color"

type RGB struct {
	R, G, B float64
}

var _ color.Color = RGB{}

func (c RGB) RGBA() (r, g, b, a uint32) {
	r = channelToMax(c.R)
	g = channelToMax(c.G)
	b = channelToMax(c.B)
	a = channelMax
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
		R: float64(r) / channelMax,
		G: float64(g) / channelMax,
		B: float64(b) / channelMax,
	}
}
