package pointf

import (
	"fmt"
	"math"
)

type Point2f struct {
	X, Y float64
}

func Pt2f(x, y float64) Point2f {
	return Point2f{
		X: x,
		Y: y,
	}
}

func (a Point2f) Add(b Point2f) Point2f {
	return Point2f{
		X: a.X + b.X,
		Y: a.Y + b.Y,
	}
}

func (a Point2f) Sub(b Point2f) Point2f {
	return Point2f{
		X: a.X - b.X,
		Y: a.Y - b.Y,
	}
}

func (a Point2f) MulScalar(scalar float64) Point2f {
	return Point2f{
		X: a.X * scalar,
		Y: a.Y * scalar,
	}
}

func (a Point2f) DivScalar(scalar float64) Point2f {
	return Point2f{
		X: a.X / scalar,
		Y: a.Y / scalar,
	}
}

// func (p Point2f) Negative() Point2f {
// 	return Point2f{
// 		X: -p.X,
// 		Y: -p.Y,
// 	}
// }

func (a Point2f) Invert() Point2f {
	return Point2f{
		X: -a.X,
		Y: -a.Y,
	}
}

func (a Point2f) InvertAxisX() Point2f {
	return Point2f{
		X: -a.X,
		Y: a.Y,
	}
}

func (a Point2f) InvertAxisY() Point2f {
	return Point2f{
		X: a.X,
		Y: -a.Y,
	}
}

func (p Point2f) String() string {
	return fmt.Sprintf("(%.12f, %.12f)", p.X, p.Y)
}

// Point2f lerp
func PtLerp(p0, p1 Point2f, t float64) Point2f {
	return Point2f{
		X: lerp(p0.X, p1.X, t),
		Y: lerp(p0.Y, p1.Y, t),
	}
}

func Distance(a, b Point2f) float64 {
	var (
		dx = a.X - b.X
		dy = a.Y - b.Y
	)
	return math.Sqrt(dx*dx + dy*dy)
}
