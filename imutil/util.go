package imutil

import (
	"image/color"

	"github.com/gitchander/cairo"
)

func CanvasSetColor(canvas *cairo.Canvas, cr color.Color) {
	r, g, b, a := cr.RGBA()
	const max = 0xffff
	var (
		R = float64(r) / max
		G = float64(g) / max
		B = float64(b) / max
	)
	if a == max {
		canvas.SetSourceRGB(R, G, B)
	} else {
		var A = float64(a) / max
		canvas.SetSourceRGBA(R, G, B, A)
	}
}

func CanvasFillColor(canvas *cairo.Canvas, cr color.Color) {
	canvas.Save()
	CanvasSetColor(canvas, cr)
	canvas.Paint()
	canvas.Restore()
}
