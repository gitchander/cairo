package main

import (
	"image/color"
)

func ColorOver(dc, sc color.Color) color.Color {
	var (
		dc1 = color.RGBA64Model.Convert(dc).(color.RGBA64)
		sc1 = color.RGBA64Model.Convert(sc).(color.RGBA64)
	)
	return colorOverRGBA64(dc1, sc1)
}

func colorOverRGBA64(dc, sc color.RGBA64) color.RGBA64 {

	// m is the maximum color value returned by image.Color.RGBA.
	const m = 1<<16 - 1

	a := m - uint32(sc.A)

	return color.RGBA64{
		R: uint16((uint32(dc.R)*a)/m) + sc.R,
		G: uint16((uint32(dc.G)*a)/m) + sc.G,
		B: uint16((uint32(dc.B)*a)/m) + sc.B,
		A: uint16((uint32(dc.A)*a)/m) + sc.A,
	}
}
