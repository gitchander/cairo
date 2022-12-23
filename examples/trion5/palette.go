package main

import (
	"image/color"

	"github.com/gitchander/cairo"
	"github.com/gitchander/cairo/colorf"
)

var palette1 = []color.Color{
	color.RGBA{0x58, 0x8C, 0x7E, 0xff},
	color.RGBA{0xF2, 0xE3, 0x94, 0xff},
	color.RGBA{0xD9, 0x64, 0x59, 0xff},
	color.RGBA{0xF2, 0xAE, 0x72, 0xff},
	color.RGBA{0x8C, 0x46, 0x46, 0xff},
}

var palette2 = []color.Color{
	color.RGBA{0x9e, 0xd6, 0x70, 0xff},
	color.RGBA{0x4d, 0x73, 0x58, 0xff},
	color.RGBA{0xd6, 0x4d, 0x4d, 0xff},
	color.RGBA{0xe8, 0xd1, 0x74, 0xff},
	color.RGBA{0xe3, 0x9e, 0x54, 0xff},
}

func ColorIndex(i int) color.Color {
	pal := palette2
	return pal[mod(i, len(pal))]
}

func mod(a, b int) int {
	d := a % b
	if d < 0 {
		d += b
	}
	return d
}

func canvasSetColor(canvas *cairo.Canvas, cl color.Color) {
	v := colorf.NRGBAfModel.Convert(cl).(colorf.NRGBAf)
	canvas.SetSourceRGBA(v.R, v.G, v.B, v.A)
}
