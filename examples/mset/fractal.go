package main

// Z= Z ^ 2 + C
func MandelbrotSet(x, y float64, n int) int {

	var xx, yy, xy float64

	Cx := x
	Cy := y

	for i := 0; i < n; i++ {

		xx = x * x
		yy = y * y

		if xx+yy > 4.0 {
			return i
		}

		xy = x * y
		y = xy + xy + Cy
		x = xx - yy + Cx
	}

	return n
}

func JuliaSet(x, y float64, Cx, Cy float64, n int) int {

	var xx, yy, xy float64

	for i := 0; i < n; i++ {

		xx = x * x
		yy = y * y

		if xx+yy > 4.0 {
			return i
		}

		xy = x * y
		y = xy + xy + Cy
		x = xx - yy + Cx
	}

	return n
}

// Z= Z ^ 3 + C
func MandelbrotSetPow3(x, y float64, n int) int {

	var xx, yy float64

	Cx := x
	Cy := y

	for i := 0; i < n; i++ {

		xx = x * x
		yy = y * y

		if xx+yy > 4.0 {
			return i
		}

		x = x*(xx-3.0*yy) + Cx
		y = y*(3.0*xx-yy) + Cy
	}

	return n
}
