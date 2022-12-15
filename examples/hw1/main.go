package main

import (
	"image"
	"log"

	"github.com/gitchander/cairo"
)

func main() {
	size := image.Point{X: 256, Y: 256}
	surface, err := cairo.NewSurface(cairo.FORMAT_ARGB32, size.X, size.Y)
	checkError(err)
	defer surface.Destroy()

	canvas, err := cairo.NewCanvas(surface)
	checkError(err)
	defer canvas.Destroy()

	drawHelloWorld(canvas)

	err = surface.WriteToPNG("test.png")
	checkError(err)
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func drawHelloWorld(c *cairo.Canvas) error {

	c.SelectFontFace("Sans", cairo.FONT_SLANT_NORMAL, cairo.FONT_WEIGHT_BOLD)
	c.SetFontSize(32.0)
	c.SetSourceRGB(0.0, 0.0, 0.0)
	c.MoveTo(10.0, 140.0)
	c.ShowText("Hello, World!")

	return nil
}
