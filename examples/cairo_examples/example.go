package main

import (
	"image"
	"image/color"
	"path/filepath"

	"github.com/gitchander/cairo"
	"github.com/gitchander/cairo/imutil"
)

type Example struct {
	SampleFunc func(c *cairo.Canvas) error
	Dir        string
	FileName   string
	Size       image.Point
}

func (e *Example) Execute() error {

	fileName := filepath.Join(e.Dir, e.FileName)

	surface, err := cairo.NewSurface(cairo.FORMAT_ARGB32, e.Size.X, e.Size.Y)
	if err != nil {
		return err
	}
	defer surface.Destroy()

	canvas, err := cairo.NewCanvas(surface)
	if err != nil {
		return err
	}
	defer canvas.Destroy()

	imutil.CanvasFillColor(canvas, color.White)

	err = e.SampleFunc(canvas)
	if err != nil {
		return err
	}

	if err = surface.WriteToPNG(fileName); err != nil {
		return err
	}

	return nil
}
