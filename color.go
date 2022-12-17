package cairo

// #cgo pkg-config: cairo cairo-gobject
// #include <stdlib.h>
// #include <cairo.h>
// #include <cairo-gobject.h>
import "C"

import (
	"image/color"

	"github.com/gitchander/cairo/colorf"
)

const maxColorComponent = 0xffff

func (c *Canvas) SetSourceRGB(red, green, blue float64) {
	C.cairo_set_source_rgb(c.cr,
		C.double(red), C.double(green), C.double(blue))
}

func (c *Canvas) SetSourceRGBA(red, green, blue, alpha float64) {
	C.cairo_set_source_rgba(c.cr,
		C.double(red), C.double(green), C.double(blue), C.double(alpha))
}

func (c *Canvas) setSourceColor1(r color.Color) {

	if cf, ok := r.(colorf.NColorf); ok {
		c.SetSourceRGBA(cf.R, cf.G, cf.B, cf.A)
		return
	}

	cu := color.NRGBA64Model.Convert(r).(color.NRGBA64)

	cf := colorf.NColorf{
		R: float64(cu.R) / maxColorComponent,
		G: float64(cu.G) / maxColorComponent,
		B: float64(cu.B) / maxColorComponent,
		A: float64(cu.A) / maxColorComponent,
	}

	useRGB := (cu.A == maxColorComponent)

	if useRGB {
		c.SetSourceRGB(cf.R, cf.G, cf.B)
	} else {
		c.SetSourceRGBA(cf.R, cf.G, cf.B, cf.A)
	}
}

func (c *Canvas) setSourceColor2(cr color.Color) {

	cf := colorf.NColorfModel.Convert(cr).(colorf.NColorf)

	_, _, _, a := cr.RGBA()
	useRGB := (a == maxColorComponent)

	if useRGB {
		c.SetSourceRGB(cf.R, cf.G, cf.B)
	} else {
		c.SetSourceRGBA(cf.R, cf.G, cf.B, cf.A)
	}
}

func (c *Canvas) setSourceColor3(cr color.Color) {

	cf := colorf.NColorfModel.Convert(cr).(colorf.NColorf)

	useRGB := (cf.A == 1.0)

	if useRGB {
		c.SetSourceRGB(cf.R, cf.G, cf.B)
	} else {
		c.SetSourceRGBA(cf.R, cf.G, cf.B, cf.A)
	}
}

func (c *Canvas) SetSourceColor(cr color.Color) {
	// c.setSourceColor1(cr)
	c.setSourceColor2(cr)
	//c.setSourceColor3(cr)
}
