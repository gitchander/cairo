package main

import (
	"math"
)

// https://en.wikipedia.org/wiki/Lemniscate_of_Bernoulli

// Lemniscate of Bernoulli
func lemniscate(radius, theta float64) (float64, bool) {
	x := square(radius) * math.Cos(2*theta)
	if x < 0 {
		return 0, false
	}
	return math.Sqrt(x), true
}

func square(a float64) float64 {
	return a * a
}
