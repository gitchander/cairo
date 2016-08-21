package color

import (
	"math"
)

func norm(channel float64) float64 {

	if channel < 0.0 {
		channel = 0.0
	}

	if channel > 1.0 {
		channel = 1.0
	}

	return channel
}

func round(x float64) float64 {
	return math.Floor(x + 0.5)
}

// Lerp - Linear interpolation
// t= [0, 1]
// (t == 0) => v0
// (t == 1) => v1
func lerp(v0, v1 float64, t float64) float64 {
	return (1.0-t)*v0 + t*v1
}
