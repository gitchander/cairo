package mathf

import (
	"math"
)

type Complex struct {
	Re float64
	Im float64
}

func (a Complex) Add(b Complex) Complex {
	return Complex{
		Re: a.Re + b.Re,
		Im: a.Im + b.Im,
	}
}

func (a Complex) Sub(b Complex) Complex {
	return Complex{
		Re: a.Re - b.Re,
		Im: a.Im - b.Im,
	}
}

func (a Complex) Mul(b Complex) (c Complex) {
	c.Re = a.Re*b.Re - a.Im*b.Im
	c.Im = a.Im*b.Re + a.Re*b.Im
	return
}

func (a Complex) Div(b Complex) (c Complex) {
	norm := b.Norm()
	c.Re = (a.Re*b.Re + a.Im*b.Im) / norm
	c.Im = (a.Im*b.Re - a.Re*b.Im) / norm
	return
}

func (a Complex) Norm() float64 {
	return (a.Re * a.Re) + (a.Im * a.Im)
}

func (a Complex) AddScalar(scalar float64) (c Complex) {
	c.Re = a.Re + scalar
	c.Im = a.Im
	return
}

func (a Complex) SubScalar(scalar float64) (c Complex) {
	c.Re = a.Re - scalar
	c.Im = a.Im
	return
}

func (a Complex) MulScalar(scalar float64) (c Complex) {
	c.Re = a.Re * scalar
	c.Im = a.Im * scalar
	return
}

func (a Complex) DivScalar(scalar float64) (c Complex) {
	c.Re = a.Re / scalar
	c.Im = a.Im / scalar
	return
}

func (z Complex) Magnitude() float64 {

	var (
		a = math.Abs(z.Re)
		b = math.Abs(z.Im)
	)

	if a == 0 {
		return b
	}
	if b == 0 {
		return a
	}

	var m float64
	if a >= b {
		m = b / a
		m = a * math.Sqrt(1+m*m)
	} else {
		m = a / b
		m = b * math.Sqrt(1+m*m)
	}

	return m
}

func (a Complex) Argument() float64 {
	return math.Atan2(a.Im, a.Re)
}

// b = 1 / a
func (a Complex) Invert() (b Complex) {
	norm := a.Norm()
	return Complex{
		Re: a.Re / norm,
		Im: -a.Im / norm,
	}
}

func (a Complex) Polar() Polar {
	return Polar{
		Rho: a.Magnitude(),
		Phi: a.Argument(),
	}
}

func (a Complex) Power(p float64) Complex {
	c := Polar{
		Rho: math.Exp(p * 0.5 * math.Log(a.Norm())),
		Phi: a.Argument() * p,
	}.ToCartesian()
	return Complex{
		Re: c.X,
		Im: c.Y,
	}
}

func (a Complex) PowerN(n int) Complex {
	b := Complex{Re: 1}
	for n > 0 {
		if isOdd(n) { // n is odd
			b = b.Mul(a) // b = b * a
		}
		n /= 2
		a = a.Mul(a)
	}
	return b
}

func isOdd(n int) bool {
	return ((n % 2) != 0)
}
