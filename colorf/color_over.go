package colorf

import (
	"image/color"
)

// ------------------------------------------------------------------------------
func ColorOver1(dc, sc color.Color) color.Color {
	var (
		dc1 = NColorfModel.Convert(dc).(NColorf)
		sc1 = NColorfModel.Convert(sc).(NColorf)
	)
	return colorOver_NColorf(dc1, sc1)
}

func ColorOver2(dc, sc color.Color) color.Color {
	var (
		dc1 = color.RGBA64Model.Convert(dc).(color.RGBA64)
		sc1 = color.RGBA64Model.Convert(sc).(color.RGBA64)
	)
	return colorOver_RGBA64(dc1, sc1)
}

// ------------------------------------------------------------------------------
// Alpha blending
// sc over dc
func colorOver_NColorf(dc, sc NColorf) NColorf {
	A := lerp(dc.A, 1.0, sc.A)
	return NColorf{
		R: lerp(dc.R*dc.A, sc.R, sc.A) / A,
		G: lerp(dc.G*dc.A, sc.G, sc.A) / A,
		B: lerp(dc.B*dc.A, sc.B, sc.A) / A,
		A: A,
	}
}

// ------------------------------------------------------------------------------
func colorOver_RGBA64(dc, sc color.RGBA64) color.RGBA64 {

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
