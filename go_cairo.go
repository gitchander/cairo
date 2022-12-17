package cairo

// #include <cairo.h>
import "C"

import (
	"image/color"
	"math"
)

const tau = 2.0 * math.Pi

func (c *Canvas) Circle(xc, yc float64, radius float64) {
	c.Arc(xc, yc, radius, 0, tau)
}

func (c *Canvas) FillColor(cr color.Color) {
	c.Save()
	c.SetSourceColor(cr)
	c.Paint()
	c.Restore()
}
