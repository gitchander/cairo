package main

import (
	"image"
	"image/color"
	"path/filepath"

	"github.com/gitchander/cairo"
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

	c, err := cairo.NewCanvas(surface)
	if err != nil {
		return err
	}
	defer c.Destroy()

	c.FillColor(color.White)

	err = e.SampleFunc(c)
	if err != nil {
		return err
	}

	return surface.WriteToPNG(fileName)
}
