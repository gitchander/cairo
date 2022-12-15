package main

import "math"

type Curve interface {
	Name() string
	RadiusByAngle(Angle float64) float64
}

type Named struct {
	name string
}

func (p *Named) Name() string {
	return p.name
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

func (p *Spiral) RadiusByAngle(Angle float64) float64 {
	return p.a * Angle
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

func (p *Cardioid) RadiusByAngle(Angle float64) float64 {
	return p.a * (1.0 - math.Cos(Angle))
}

type Lemniscate struct {
	Named
	radius float64
}

func NewLemniscate(name string, radius float64) Curve {
	return &Lemniscate{
		Named:  Named{name},
		radius: radius,
	}
}

func (p *Lemniscate) RadiusByAngle(angle float64) float64 {
	r, _ := lemniscate(p.radius, angle)
	return r
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

func (p *Cannabis) RadiusByAngle(Angle float64) float64 {
	return p.a * (1.0 + 9.0/10.0*math.Cos(8.0*Angle)) *
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

func (p *Rose) RadiusByAngle(Angle float64) float64 {
	return p.a * math.Sin(Angle*p.k)
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

func (p *Circle) RadiusByAngle(Angle float64) float64 {
	return 2.0 * p.a * math.Sin(Angle)
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

func (p *Strofoid) RadiusByAngle(Angle float64) float64 {

	sin, cos := math.Sincos(Angle)

	strofoid := p.b * (1.0 + cos) / sin

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

func (p *StrofoidKnot) RadiusByAngle(Angle float64) float64 {

	sin, cos := math.Sincos(Angle)

	circle := 2.0 * p.a * sin
	strofoid := p.b * (1.0 + cos) / sin
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

func (p *ParabolaKnot) RadiusByAngle(Angle float64) float64 {

	sin, cos := math.Sincos(Angle)

	circle := 2.0 * p.a * sin
	parabola := 2.0 * p.p * sin / (cos * cos)
	knot := circle - parabola

	return knot
}
