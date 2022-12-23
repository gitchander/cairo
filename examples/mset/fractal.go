package main

import (
	. "github.com/gitchander/cairo/examples/mathf"
)

// ------------------------------------------------------------------------------
// Mandelbrot set
// https://en.wikipedia.org/wiki/Mandelbrot_set
// ------------------------------------------------------------------------------
// Z[0] = 0
// Z[n] = Z[n-1] ^ 2 + C
// ------------------------------------------------------------------------------
func MandelbrotSet(c Complex, n int) (int, bool) {
	var z Complex
	for i := 0; i < n; i++ {
		if z.Norm() > 4.0 {
			return i, false
		}
		z = z.Mul(z).Add(c)
	}
	return 0, true
}

// ------------------------------------------------------------------------------
// Z[0] = z0
// Z[n] = Z[n-1] ^ 2 + C
func JuliaSet(c, z0 Complex, n int) (int, bool) {
	z := z0
	for i := 0; i < n; i++ {
		if z.Norm() > 4.0 {
			return i, false
		}
		z = z.Mul(z).Add(c)
	}
	return 0, true
}

// ------------------------------------------------------------------------------
// Z[0] = 0
// Z[n] = Z[n-1] ^ 3 + C
// ------------------------------------------------------------------------------
func MandelbrotSetPow3(c Complex, n int) (int, bool) {
	var z Complex
	for i := 0; i < n; i++ {
		if z.Norm() > 4.0 {
			return i, false
		}
		z = z.PowerN(3).Add(c)
	}
	return 0, true
}

// ------------------------------------------------------------------------------
// Z[0] = 0
// Z[n] = Z[n-1] ^ 3 + C
// ------------------------------------------------------------------------------
func MandelbrotSetPowM(c Complex, m int, n int) (int, bool) {
	var z Complex
	for i := 0; i < n; i++ {
		if z.Norm() > 4.0 {
			return i, false
		}
		z = z.PowerN(m).Add(c)
	}
	return 0, true
}

// ------------------------------------------------------------------------------
func MandelbrotSetPow(c Complex, p float64, n int) (int, bool) {
	var z Complex
	for i := 0; i < n; i++ {
		if z.Norm() > 4.0 {
			return i, false
		}
		z = z.Power(p).Add(c)
	}
	return 0, true
}
