package imutil

import (
	"image/color"

	"github.com/gitchander/cairo"
)

func CanvasFillColor(canvas *cairo.Canvas, cl color.Color) {

	surface := canvas.GetTarget()
	if surface == nil {
		return
	}

	r, g, b, a := cl.RGBA()

	const max = 65535

	var (
		R = float64(r) / max
		G = float64(g) / max
		B = float64(b) / max
		A = float64(a) / max
	)

	canvas.Save()

	if a == max {
		canvas.SetSourceRGB(R, G, B)
	} else {
		canvas.SetSourceRGBA(R, G, B, A)
		canvas.SetOperator(cairo.OPERATOR_SOURCE)
	}

	canvas.Paint()
	canvas.Restore()
}

/*
func CanvasFillRGBA(canvas *cairo.Canvas, color RGBA) {

	surface := canvas.GetTarget()
	if surface == nil {
		return
	}

	r, g, b, a := color.GetRGBA()

	canvas.Save()
	canvas.SetSourceRGBA(r, g, b, a)
	canvas.SetOperator(cairo.OPERATOR_SOURCE)
	canvas.Paint()
	canvas.Restore()
}

func SurfaceFillRGB(surface *cairo.Surface, c RGB) {

	canvas, _ := cairo.NewCanvas(surface)
	defer canvas.Destroy()

	r, g, b := c.GetRGB()
	canvas.SetSourceRGB(r, g, b)
	canvas.Paint()
}
*/
