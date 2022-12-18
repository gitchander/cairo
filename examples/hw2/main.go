package main

import (
	"image"
	"log"
	"math"

	"github.com/gitchander/cairo"
	"github.com/gitchander/cairo/examples/mathf"
)

func main() {
	size := image.Point{X: 256, Y: 256}
	surface, err := cairo.NewSurface(cairo.FORMAT_ARGB32, size.X, size.Y)
	checkError(err)
	defer surface.Destroy()

	canvas, err := cairo.NewCanvas(surface)
	checkError(err)
	defer canvas.Destroy()

	draw(canvas, size)

	err = surface.WriteToPNG("test.png")
	checkError(err)
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func draw(canvas *cairo.Canvas, size image.Point) {
	canvas.Rectangle(0, 0, float64(size.X), float64(size.Y))
	canvas.SetSourceRGB(1, 1, 1)
	canvas.Fill()

	n := image.Point{X: 9, Y: 9}
	const wcoef = 0.7

	var (
		dx = float64(size.X) / float64(n.X)
		dy = float64(size.Y) / float64(n.Y)

		x0 = dx * (1 - wcoef) / 2
		y0 = dy * (1 - wcoef) / 2
	)

	for x := 0; x < n.X; x++ {
		for y := 0; y < n.Y; y++ {
			canvas.Rectangle(x0+dx*float64(x), y0+dy*float64(y), dx*wcoef, dy*wcoef)
		}
	}

	//	canvas.SetSourceRGB(1, 0, 0)
	//	canvas.Fill()

	var (
		x = float64(size.X) / 2
		y = float64(size.Y) / 2
	)

	pattern, err := cairo.NewPatternRadial(x, y, 0, x, y, float64(minInt(size.X, size.Y))/2)
	checkError(err)
	defer pattern.Destroy()
	pattern.AddColorStopRGB(0, 0.75, 0.15, 0.99)
	pattern.AddColorStopRGB(0.9, 1, 1, 1)
	canvas.SetSource(pattern)
	canvas.Fill()

	canvas.SetFontSize(30)
	canvas.SelectFontFace("Georgia", cairo.FONT_SLANT_NORMAL, cairo.FONT_WEIGHT_BOLD)
	canvas.SetSourceRGB(0.1, 0.0, 0.5)

	lines := []string{
		"Hello,",
		"World!",
		"cairo is",
		"the best!",
	}

	center := mathf.Point2f{
		X: float64(size.X),
		Y: float64(size.Y),
	}.DivScalar(2)
	drawLines(canvas, lines, center)
}

func minInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func maxFloat64(a, b float64) float64 {
	if a > b {
		return a
	}
	return b
}

func drawLines(c *cairo.Canvas, lines []string, center mathf.Point2f) {

	ts := make([]*cairo.TextExtents, len(lines))
	for i, line := range lines {
		t := new(cairo.TextExtents)
		c.TextExtents(line, t)
		ts[i] = t
	}

	var size mathf.Point2f
	var dY float64
	for _, t := range ts {
		size.X = maxFloat64(size.X, t.Width)
		size.Y += t.Height
		dY = maxFloat64(dY, t.Height)
	}
	dY *= 0.5
	size.Y += dY * float64(len(ts)-1)

	p := center.Sub(size.DivScalar(2))
	if len(ts) > 0 {
		p.Y += ts[0].Height
	}
	for i, t := range ts {
		c.MoveTo(p.X, p.Y)
		c.SetSourceRGB(0.1, 0.0, 0.5)
		c.ShowText(lines[i])
		p.Y += t.Height + dY
	}
}

func drawPoint(c *cairo.Canvas, p mathf.Point2f) {
	radius := 3.0
	c.Arc(p.X, p.Y, radius, 0, 2*math.Pi)
	c.SetSourceRGB(1, 1, 1)
	c.FillPreserve()
	c.SetLineWidth(1)
	c.SetSourceRGB(0, 0, 0)
	c.Stroke()
}
