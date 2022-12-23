package colorf

import (
	"image/color"
)

func ColorOver(dc, sc color.Color) color.Color {
	switch 2 {
	case 0:
		return colorOver1(dc, sc)
	case 1:
		return colorOver2(dc, sc)
	case 2:
		return colorOver3(dc, sc)
	default:
		return color.Black
	}
}

func colorOver1(dc, sc color.Color) color.Color {
	var (
		dc1 = NRGBAfModel.Convert(dc).(NRGBAf)
		sc1 = NRGBAfModel.Convert(sc).(NRGBAf)
	)
	return colorOver_NRGBAf(dc1, sc1)
}

func colorOver2(dc, sc color.Color) color.Color {
	var (
		dc1 = color.RGBA64Model.Convert(dc).(color.RGBA64)
		sc1 = color.RGBA64Model.Convert(sc).(color.RGBA64)
	)
	return colorOver_RGBA64(dc1, sc1)
}

func colorOver3(dc, sc color.Color) color.Color {
	var (
		dc1 = RGBAfModel.Convert(dc).(RGBAf)
		sc1 = RGBAfModel.Convert(sc).(RGBAf)
	)
	return colorOver_RGBAf(dc1, sc1)
}

// ------------------------------------------------------------------------------
// Alpha blending
// sc over dc
func colorOver_NRGBAf(dc, sc NRGBAf) NRGBAf {

	// Straight alpha
	// result = (dest.RGB * (1 - source.A)) + (source.RGB * source.A)
	// result = lerp(dest.RGB, source.RGB, source.A)

	A := lerp(dc.A, 1.0, sc.A)

	return NRGBAf{
		R: lerp(dc.R*dc.A, sc.R, sc.A) / A,
		G: lerp(dc.G*dc.A, sc.G, sc.A) / A,
		B: lerp(dc.B*dc.A, sc.B, sc.A) / A,
		A: A,
	}
}

// ------------------------------------------------------------------------------
func colorOver_RGBAf(dc, sc RGBAf) RGBAf {

	// Premultiplied alpha
	// result = (dest.RGB * (1 - source.A)) + source.RGB

	a := (1.0 - sc.A)

	return RGBAf{
		R: dc.R*a + sc.R,
		G: dc.G*a + sc.G,
		B: dc.B*a + sc.G,
		A: dc.A*a + sc.A,
	}
}

// ------------------------------------------------------------------------------
func colorOver_RGBA64(dc, sc color.RGBA64) color.RGBA64 {

	// Premultiplied alpha
	// result = (dest.RGB * (1 - source.A)) + source.RGB

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
