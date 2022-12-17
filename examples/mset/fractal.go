package main

type Complex struct {
	Re float64
	Im float64
}

// Z[n] = Z[n-1] ^ 2 + C
func MandelbrotSet(z Complex, n int) (int, bool) {

	var zrr, zii, zri float64

	c := z

	for i := 0; i < n; i++ {

		zrr = z.Re * z.Re
		zii = z.Im * z.Im

		if (zrr + zii) > 4.0 {
			return i, false
		}

		zri = z.Re * z.Im

		z.Im = zri + zri + c.Im
		z.Re = zrr - zii + c.Re
	}

	return 0, true
}

func JuliaSet(x, y float64, Cx, Cy float64, n int) (int, bool) {

	var xx, yy, xy float64

	for i := 0; i < n; i++ {

		xx = x * x
		yy = y * y

		if xx+yy > 4.0 {
			return i, false
		}

		xy = x * y
		y = xy + xy + Cy
		x = xx - yy + Cx
	}

	return 0, true
}

// Z[n] = Z[n-1] ^ 3 + C
func MandelbrotSetPow3(x, y float64, n int) (int, bool) {

	var xx, yy float64

	Cx := x
	Cy := y

	for i := 0; i < n; i++ {

		xx = x * x
		yy = y * y

		if xx+yy > 4.0 {
			return i, false
		}

		x = x*(xx-3.0*yy) + Cx
		y = y*(3.0*xx-yy) + Cy
	}

	return 0, true
}
