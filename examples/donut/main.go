package main

import (
	"image"
	"image/color"
	"log"
	"math"

	"github.com/gitchander/cairo"
	. "github.com/gitchander/cairo/examples/mathf"
)

const tau = 2.0 * math.Pi

func main() {
	checkError(makeDonutImage())
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func makeDonutImage() error {

	fileName := "donut.png"
	size := image.Pt(256, 256)

	surface, err := cairo.NewSurface(cairo.FORMAT_ARGB32, size.X, size.Y)
	if err != nil {
		return err
	}
	defer surface.Destroy()

	c, err := cairo.NewCanvas(surface)
	if err != nil {
		return err
	}
	defer c.Destroy()

	c.FillColor(color.White)

	err = drawDonut(c, size)
	if err != nil {
		return err
	}

	return surface.WriteToPNG(fileName)
}

func drawDonut(c *cairo.Canvas, size image.Point) error {

	var (
		vp     = Pt2f(float64(size.X), float64(size.Y))
		center = vp.DivScalar(2)

		vm = vmin(vp)
	)

	var (
		radius    = 45 * vm
		lineWidth = 0.2 * vm
	)

	c.SetLineWidth(lineWidth)
	c.Translate(center.X, center.Y)
	c.Arc(0, 0, radius, 0, tau)
	c.Stroke()

	n := 36
	da := math.Pi / float64(n)

	for i := 0; i < n; i++ {
		c.Save()
		c.Rotate(float64(i) * da)
		c.Scale(0.3, 1)
		c.Arc(0, 0, radius, 0, tau)
		c.Restore()
		c.Stroke()
	}

	return nil
}

func min(a, b float64) float64 {
	if a < b {
		return a
	}
	return b
}

// viewport min
func vmin(viewportSize Point2f) float64 {
	return min(viewportSize.X, viewportSize.Y) / 100.0
}
