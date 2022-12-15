package main

import (
	"fmt"
	"image/color"

	"github.com/gitchander/cairo/colorf"
)

var (
	White   = color.RGBA{R: 0xff, G: 0xff, B: 0xff, A: 0xff}
	Black   = color.RGBA{R: 0x00, G: 0x00, B: 0x00, A: 0xff}
	Red     = color.RGBA{R: 0xff, G: 0x00, B: 0x00, A: 0xff}
	Green   = color.RGBA{R: 0x00, G: 0x80, B: 0x00, A: 0xff}
	Blue    = color.RGBA{R: 0x00, G: 0x00, B: 0xff, A: 0xff}
	Yellow  = color.RGBA{R: 0xff, G: 0xff, B: 0x00, A: 0xff}
	Cyan    = color.RGBA{R: 0x00, G: 0xff, B: 0xff, A: 0xff}
	Magenta = color.RGBA{R: 0xff, G: 0x00, B: 0xff, A: 0xff}
)

func main() {
	convertExample()
}

func convertExample() {

	cs := []color.Color{
		White,
		Black,
		Red,
		Green,
		Blue,
		Yellow,
		Cyan,
		Magenta,
	}

	for _, c := range cs {
		nc := colorf.NColorfModel.Convert(c).(colorf.NColorf)
		fmt.Printf("color:%3v, rgb:%v\n", c, nc)
	}
}
