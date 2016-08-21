package main

import "math"

type Curve interface {
	Name() string
	RadiusByAngle(Angle float64) float64
}

type Named struct {
	name string
}

func (this *Named) Name() string {
	return this.name
}

type Spiral struct {
	Named
	a float64
}

func NewSpiral(name string, a float64) Curve {
	return &Spiral{
		Named: Named{name},
		a:     a,
	}
}

func (this *Spiral) RadiusByAngle(Angle float64) float64 {
	return this.a * Angle
}

// r= a * (1 - cos(angle))
type Cardioid struct {
	Named
	a float64
}

func NewCardioid(name string, a float64) Curve {
	return &Cardioid{
		Named: Named{name},
		a:     a,
	}
}

func (this *Cardioid) RadiusByAngle(Angle float64) float64 {
	return this.a * (1.0 - math.Cos(Angle))
}

type Lemniscate struct {
	Named
	a float64
}

func NewLemniscate(name string, a float64) Curve {
	return &Lemniscate{
		Named: Named{name},
		a:     a,
	}
}

func (this *Lemniscate) RadiusByAngle(Angle float64) float64 {

	c := math.Cos(2 * Angle)
	if c < 0.0 {
		return 0.0
	}
	return math.Sqrt(this.a * this.a * c)
}

type Cannabis struct {
	Named
	a float64
}

func NewCannabis(name string, a float64) Curve {
	return &Cannabis{
		Named: Named{name},
		a:     a,
	}
}

func (this *Cannabis) RadiusByAngle(Angle float64) float64 {

	return this.a * (1.0 + 9.0/10.0*math.Cos(8.0*Angle)) *
		(1.0 + 1.0/10.0*math.Cos(24.0*Angle)) *
		(9.0/10.0 + 1.0/10.0*math.Cos(200.0*Angle)) *
		(1.0 + math.Sin(Angle))
}

// r= a * sin(k * angle)
type Rose struct {
	Named
	a float64
	k float64
}

func NewRose(name string, a, k float64) Curve {
	return &Rose{
		Named: Named{name},
		a:     a,
		k:     k,
	}
}

func (this *Rose) RadiusByAngle(Angle float64) float64 {
	return this.a * math.Sin(Angle*this.k)
}

type Circle struct {
	Named
	a float64
}

func NewCircle(name string, a float64) Curve {
	return &Circle{
		Named: Named{name},
		a:     a,
	}
}

func (this *Circle) RadiusByAngle(Angle float64) float64 {
	return 2.0 * this.a * math.Sin(Angle)
}

type Strofoid struct {
	Named
	b float64
}

func NewStrofoid(name string, a float64) Curve {
	return &Strofoid{
		Named: Named{name},
		b:     a,
	}
}

func (this *Strofoid) RadiusByAngle(Angle float64) float64 {

	sin, cos := math.Sincos(Angle)

	strofoid := this.b * (1.0 + cos) / sin

	return strofoid
}

type StrofoidKnot struct {
	Named
	a, b float64
}

func NewStrofoidKnot(name string, c float64) Curve {
	return &StrofoidKnot{
		Named: Named{name},
		a:     c,
		b:     2.0 * c,
	}
}

func (this *StrofoidKnot) RadiusByAngle(Angle float64) float64 {

	sin, cos := math.Sincos(Angle)

	circle := 2.0 * this.a * sin
	strofoid := this.b * (1.0 + cos) / sin
	knot := circle - strofoid

	return knot
}

type ParabolaKnot struct {
	Named
	a, p float64
}

func NewParabolaKnot(name string, c float64) Curve {
	return &ParabolaKnot{
		Named: Named{name},
		a:     c,
		p:     c / 4.0,
	}
}

func (this *ParabolaKnot) RadiusByAngle(Angle float64) float64 {

	sin, cos := math.Sincos(Angle)

	circle := 2.0 * this.a * sin
	parabola := 2.0 * this.p * sin / (cos * cos)
	knot := circle - parabola

	return knot
}
