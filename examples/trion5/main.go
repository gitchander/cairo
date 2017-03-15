package main

import (
	"log"
	"math"

	"github.com/gitchander/cairo"
)

// Splitting equilateral triangle into 5 equal parts

const (
	sqrt3 = 1.7320508075688772

	sqrt3div2 = sqrt3 / 2
	sqrt3div4 = sqrt3 / 4
	sqrt3div6 = sqrt3 / 6
)

type Pointf struct {
	X, Y float64
}

func main() {

	surface, err := cairo.NewSurface(cairo.FORMAT_ARGB32, 800, 800)
	if err != nil {
		log.Fatal(err)
	}
	defer surface.Destroy()

	canvas, err := cairo.NewCanvas(surface)
	if err != nil {
		log.Fatal(err)
	}
	defer canvas.Destroy()

	draw(canvas)

	err = surface.WriteToPNG("trion5.png")
	if err != nil {
		log.Fatal(err)
	}
}

var palettes = [][]uint32{
	[]uint32{
		0x588C7E,
		0xF2E394,
		0xD96459,
		0xF2AE72,
		0x8C4646,
	},
	[]uint32{
		0x9ed670,
		0x4d7358,
		0xd64d4d,
		0xe8d174,
		0xe39e54,
	},
}

func draw(canvas *cairo.Canvas) {

	var width, height = canvasSize(canvas)

	radius := float64(min(width, height)) * 0.5

	var center = Pointf{
		X: float64(width) * 0.5,
		Y: float64(height)*0.5 + radius*(sqrt3/2-sqrt3/6)*0.5,
	}

	drawTriangleParts(canvas, center, radius, 0)
}

func drawTriangleParts(canvas *cairo.Canvas, center Pointf, radius float64, angle float64) {

	colors := palettes[1]
	var strokeColor uint32 = 0

	side := radius * sqrt3
	a := side / 5

	m := cairo.NewMatrix()
	m.InitIdendity()
	m.Translate(center.X, center.Y)
	m.Scale(1, -1) // Flip Vertical
	m.Rotate(angle)

	canvas.SetLineWidth(2)
	canvas.SetLineJoin(cairo.LINE_JOIN_ROUND)

	var parts = [5]bool{
		0: true,
		1: true,
		2: true,
		3: true,
		4: true,
	}

	if parts[0] {
		canvas.SetMatrix(m)
		canvas.Translate(-a*2.5, -a*sqrt3*0.75)

		pathForOnePart(canvas, a)
		fillAndStroke(canvas, colors[0], strokeColor)
	}

	if parts[1] {
		canvas.SetMatrix(m)
		canvas.Translate(a*2.5, -a*sqrt3*0.75)
		canvas.Scale(-1, 1)

		pathForOnePart(canvas, a)
		fillAndStroke(canvas, colors[1], strokeColor)
	}

	if parts[2] {
		canvas.SetMatrix(m)
		canvas.Translate(-a*2.0, -a*sqrt3div4)

		pathForOnePart(canvas, a)
		fillAndStroke(canvas, colors[2], strokeColor)
	}

	if parts[3] {
		canvas.SetMatrix(m)
		canvas.Translate(0, a*sqrt3*1.75)
		canvas.Rotate(math.Pi * 4 / 3)

		pathForOnePart(canvas, a)
		fillAndStroke(canvas, colors[3], strokeColor)
	}

	if parts[4] {
		canvas.SetMatrix(m)
		canvas.Translate(a*2, -a*sqrt3div4)
		canvas.Rotate(math.Pi * 2 / 3)

		pathForOnePart(canvas, a)
		fillAndStroke(canvas, colors[4], strokeColor)
	}
}

func canvasSize(canvas *cairo.Canvas) (width, height int) {
	surface := canvas.GetTarget()
	width = surface.GetWidth()
	height = surface.GetHeight()
	return
}

func fillAndStroke(canvas *cairo.Canvas, fillColor, strokeColor uint32) {

	setColorUint32(canvas, fillColor)
	canvas.FillPreserve()

	setColorUint32(canvas, strokeColor)
	canvas.Stroke()
}

func pathForOnePart(canvas *cairo.Canvas, a float64) {

	canvas.MoveTo(0, 0)
	canvas.LineTo(a*0.5, a*sqrt3div2)
	canvas.LineTo(a, 0)
	canvas.ClosePath()

	canvas.MoveTo(a*0.5, a*sqrt3div2)
	canvas.LineTo(a*1.5, a*sqrt3div2)
	canvas.LineTo(a, 0)
	canvas.ClosePath()

	canvas.MoveTo(a*1.5, a*sqrt3div2)
	canvas.LineTo(a*2, 0)
	canvas.LineTo(a, 0)
	canvas.ClosePath()

	canvas.MoveTo(a*1.5, a*sqrt3div2)
	canvas.LineTo(a*2.5, a*sqrt3div2)
	canvas.LineTo(a*2, 0)
	canvas.ClosePath()

	// Propeller 1
	canvas.MoveTo(a*2, 0)
	canvas.LineTo(a*2.25, a*sqrt3div4)
	canvas.LineTo(a*2.5, a*sqrt3div6)
	canvas.ClosePath()

	canvas.MoveTo(a*2.5, a*sqrt3div2)
	canvas.LineTo(a*2.75, a*sqrt3div4)
	canvas.LineTo(a*2.5, a*sqrt3div6)
	canvas.ClosePath()

	canvas.MoveTo(a*3, 0)
	canvas.LineTo(a*2.5, 0)
	canvas.LineTo(a*2.5, a*sqrt3div6)
	canvas.ClosePath()

	// Propeller 2
	canvas.MoveTo(a*2, a*sqrt3)
	canvas.LineTo(a*2.5, a*sqrt3)
	canvas.LineTo(a*2.5, a*(sqrt3+sqrt3div6))
	canvas.ClosePath()

	canvas.MoveTo(a*3, a*sqrt3)
	canvas.LineTo(a*2.75, a*(sqrt3+sqrt3div4))
	canvas.LineTo(a*2.5, a*(sqrt3+sqrt3div6))
	canvas.ClosePath()

	canvas.MoveTo(a*2.5, a*(sqrt3+sqrt3div2))
	canvas.LineTo(a*2.25, a*(sqrt3+sqrt3div4))
	canvas.LineTo(a*2.5, a*(sqrt3+sqrt3div6))
	canvas.ClosePath()
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
