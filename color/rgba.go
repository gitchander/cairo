package color

type RGBA interface {
	GetRGBA() (r, g, b, a float64)
}

type colorRGBA struct {
	r, g, b, a float64
}

func NewRGBA(r, g, b, a float64) RGBA {
	return &colorRGBA{
		norm(r),
		norm(g),
		norm(b),
		norm(a),
	}
}

func (this *colorRGBA) GetRGBA() (r, g, b, a float64) {
	r = this.r
	g = this.g
	b = this.b
	a = this.a
	return
}

// Alpha blending
// c = a over b
func Over(a, b RGBA) (c RGBA) {

	aR, aG, aB, aA := a.GetRGBA()
	bR, bG, bB, bA := b.GetRGBA()

	cA := lerp(bA, 1.0, aA)

	cR := lerp(bR*bA, aR, aA) / cA
	cG := lerp(bG*bA, aG, aA) / cA
	cB := lerp(bB*bA, aB, aA) / cA

	return NewRGBA(cR, cG, cB, cA)
}
