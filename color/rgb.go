package color

type RGB interface {
	GetRGB() (r, g, b float64)
}

type colorRGB struct {
	r, g, b float64
}

func NewRGB(r, g, b float64) RGB {
	return &colorRGB{
		norm(r),
		norm(g),
		norm(b),
	}
}

func (c *colorRGB) GetRGB() (r, g, b float64) {
	r = c.r
	g = c.g
	b = c.b
	return
}
