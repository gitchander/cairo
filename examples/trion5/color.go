package main

import (
	"github.com/gitchander/cairo"
)

var byteToFloat = func() []float64 {
	p := make([]float64, 256)
	for i := range p {
		p[i] = float64(i) / 255
	}
	return p
}()

func setColorUint32(canvas *cairo.Canvas, c uint32) {
	var (
		r = byte((c >> 16) & 0xFF)
		g = byte((c >> 8) & 0xFF)
		b = byte((c >> 0) & 0xFF)
	)
	setColorRGB(canvas, r, g, b)
}

func setColorRGB(canvas *cairo.Canvas, r, g, b byte) {
	canvas.SetSourceRGB(byteToFloat[r], byteToFloat[g], byteToFloat[b])
}
