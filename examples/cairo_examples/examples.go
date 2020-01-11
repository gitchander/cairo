package main

import (
	"math"

	"github.com/gitchander/cairo"
)

const textureFile = "data/romedalen.png"

func ExampleHelloWorld(c *cairo.Canvas) error {

	c.SelectFontFace("Sans", cairo.FONT_SLANT_NORMAL, cairo.FONT_WEIGHT_BOLD)
	c.SetFontSize(32.0)
	c.SetSourceRGB(0.0, 0.0, 0.0)
	c.MoveTo(10.0, 140.0)
	c.ShowText("Hello, World!")

	return nil
}

func ExampleArc(c *cairo.Canvas) error {

	var (
		xc     float64 = 128.0
		yc     float64 = 128.0
		radius float64 = 100.0
	)

	angle1 := DegToRad(45.0)  // angles are specified
	angle2 := DegToRad(180.0) // in radians

	c.SetLineWidth(10.0)
	c.Arc(xc, yc, radius, angle1, angle2)
	c.Stroke()

	// draw helping lines
	c.SetSourceRGBA(1, 0.2, 0.2, 0.6)
	c.SetLineWidth(6.0)

	c.Arc(xc, yc, 10.0, 0, 2*math.Pi)
	c.Fill()

	c.Arc(xc, yc, radius, angle1, angle1)
	c.LineTo(xc, yc)
	c.Arc(xc, yc, radius, angle2, angle2)
	c.LineTo(xc, yc)
	c.Stroke()

	return nil
}

func ExampleArcNegative(c *cairo.Canvas) error {

	var (
		xc     float64 = 128.0
		yc     float64 = 128.0
		radius float64 = 100.0
	)

	angle1 := DegToRad(45.0)  // angles are specified
	angle2 := DegToRad(180.0) // in radians

	c.SetLineWidth(10.0)
	c.ArcNegative(xc, yc, radius, angle1, angle2)
	c.Stroke()

	// draw helping lines
	c.SetSourceRGBA(1, 0.2, 0.2, 0.6)
	c.SetLineWidth(6.0)

	c.Arc(xc, yc, 10.0, 0, 2*math.Pi)
	c.Fill()

	c.Arc(xc, yc, radius, angle1, angle1)
	c.LineTo(xc, yc)
	c.Arc(xc, yc, radius, angle2, angle2)
	c.LineTo(xc, yc)
	c.Stroke()

	return nil
}

func ExampleClip(c *cairo.Canvas) error {

	c.Arc(128.0, 128.0, 76.8, 0, 2*math.Pi)
	c.Clip()

	c.NewPath() // current path is not consumed by cairo_clip()
	c.Rectangle(0, 0, 256, 256)
	c.Fill()
	c.SetSourceRGB(0, 1, 0)
	c.MoveTo(0, 0)
	c.LineTo(256, 256)
	c.MoveTo(256, 0)
	c.LineTo(0, 256)
	c.SetLineWidth(10.0)
	c.Stroke()

	return nil
}

func ExampleClipImage(c *cairo.Canvas) error {

	c.Arc(128.0, 128.0, 76.8, 0, 2*math.Pi)
	c.Clip()
	c.NewPath() // path not consumed by clip()

	image, err := cairo.NewSurfaceFromPNG(textureFile)
	if err != nil {
		return err
	}
	defer image.Destroy()

	w := float64(image.GetWidth())
	h := float64(image.GetHeight())

	c.Scale(256.0/w, 256.0/h)

	c.SetSourceSurface(image, 0, 0)
	c.Paint()

	return nil
}

func ExampleCurveRectangle(c *cairo.Canvas) error {

	// a custom shape that could be wrapped in a function
	var (
		x0         float64 = 25.6 // parameters like cairo_rectangle
		y0         float64 = 25.6
		rectWidth  float64 = 204.8
		rectHeight float64 = 204.8
		radius     float64 = 102.4 // and an approximate curvature radius
	)

	x1 := x0 + rectWidth
	y1 := y0 + rectHeight
	if (rectWidth == 0) || (rectHeight == 0) {
		return nil
	}
	if rectWidth/2 < radius {
		if rectHeight/2 < radius {
			c.MoveTo(x0, (y0+y1)/2)
			c.CurveTo(x0, y0, x0, y0, (x0+x1)/2, y0)
			c.CurveTo(x1, y0, x1, y0, x1, (y0+y1)/2)
			c.CurveTo(x1, y1, x1, y1, (x1+x0)/2, y1)
			c.CurveTo(x0, y1, x0, y1, x0, (y0+y1)/2)
		} else {
			c.MoveTo(x0, y0+radius)
			c.CurveTo(x0, y0, x0, y0, (x0+x1)/2, y0)
			c.CurveTo(x1, y0, x1, y0, x1, y0+radius)
			c.LineTo(x1, y1-radius)
			c.CurveTo(x1, y1, x1, y1, (x1+x0)/2, y1)
			c.CurveTo(x0, y1, x0, y1, x0, y1-radius)
		}
	} else {
		if rectHeight/2 < radius {
			c.MoveTo(x0, (y0+y1)/2)
			c.CurveTo(x0, y0, x0, y0, x0+radius, y0)
			c.LineTo(x1-radius, y0)
			c.CurveTo(x1, y0, x1, y0, x1, (y0+y1)/2)
			c.CurveTo(x1, y1, x1, y1, x1-radius, y1)
			c.LineTo(x0+radius, y1)
			c.CurveTo(x0, y1, x0, y1, x0, (y0+y1)/2)
		} else {
			c.MoveTo(x0, y0+radius)
			c.CurveTo(x0, y0, x0, y0, x0+radius, y0)
			c.LineTo(x1-radius, y0)
			c.CurveTo(x1, y0, x1, y0, x1, y0+radius)
			c.LineTo(x1, y1-radius)
			c.CurveTo(x1, y1, x1, y1, x1-radius, y1)
			c.LineTo(x0+radius, y1)
			c.CurveTo(x0, y1, x0, y1, x0, y1-radius)
		}
	}
	c.ClosePath()

	c.SetSourceRGB(0.5, 0.5, 1)
	c.FillPreserve()
	c.SetSourceRGBA(0.5, 0, 0, 0.5)
	c.SetLineWidth(10.0)
	c.Stroke()

	return nil
}

func ExampleCurveTo(c *cairo.Canvas) error {

	var (
		x, y   float64 = 25.6, 128.0
		x1, y1 float64 = 102.4, 230.4
		x2, y2 float64 = 153.6, 25.6
		x3, y3 float64 = 230.4, 128.0
	)

	c.MoveTo(x, y)
	c.CurveTo(x1, y1, x2, y2, x3, y3)

	c.SetLineWidth(10.0)
	c.Stroke()

	c.SetSourceRGBA(1, 0.2, 0.2, 0.6)
	c.SetLineWidth(6.0)
	c.MoveTo(x, y)
	c.LineTo(x1, y1)
	c.MoveTo(x2, y2)
	c.LineTo(x3, y3)
	c.Stroke()

	return nil
}

func ExampleGradient(c *cairo.Canvas) error {

	// draw background
	{
		p, _ := cairo.NewPatternLinear(0.0, 0.0, 0.0, 256.0)
		p.AddColorStopRGBA(1, 0, 0, 0, 1)
		p.AddColorStopRGBA(0, 1, 1, 1, 1)
		c.Rectangle(0, 0, 256, 256)

		c.SetSource(p)
		c.Fill()
		p.Destroy()
	}

	// draw sphere
	{
		p, _ := cairo.NewPatternRadial(
			115.2, 102.4, 25.6,
			102.4, 102.4, 128.0,
		)
		p.AddColorStopRGBA(0, 1, 1, 1, 1)
		p.AddColorStopRGBA(1, 0, 0, 0, 1)

		c.SetSource(p)
		c.Arc(128.0, 128.0, 76.8, 0, 2*math.Pi)
		c.Fill()
		p.Destroy()
	}

	return nil
}

func ExampleSetLineJoin(c *cairo.Canvas) error {

	c.SetLineWidth(40.96)
	c.MoveTo(76.8, 84.48)
	c.RelLineTo(51.2, -51.2)
	c.RelLineTo(51.2, 51.2)
	c.SetLineJoin(cairo.LINE_JOIN_MITER) // default
	c.Stroke()

	c.MoveTo(76.8, 161.28)
	c.RelLineTo(51.2, -51.2)
	c.RelLineTo(51.2, 51.2)
	c.SetLineJoin(cairo.LINE_JOIN_BEVEL)
	c.Stroke()

	c.MoveTo(76.8, 238.08)
	c.RelLineTo(51.2, -51.2)
	c.RelLineTo(51.2, 51.2)
	c.SetLineJoin(cairo.LINE_JOIN_ROUND)
	c.Stroke()

	return nil
}

func ExampleDonut(c *cairo.Canvas) error {

	var w, h float64 = 256, 256

	c.SetLineWidth(0.5)
	c.Translate(w/2, h/2)
	c.Arc(0, 0, 120, 0, 2*math.Pi)
	c.Stroke()

	for i := 0; i < 36; i++ {
		c.Save()
		c.Rotate(float64(i) * math.Pi / 36)
		c.Scale(0.3, 1)
		c.Arc(0, 0, 120, 0, 2*math.Pi)
		c.Restore()
		c.Stroke()
	}

	return nil
}

func ExampleDash(c *cairo.Canvas) error {

	var (
		dashes = []float64{
			50.0, // ink
			10.0, // skip
			10.0, // ink
			10.0, // skip
		}

		offset float64 = -50.0
	)

	c.SetDash(dashes, offset)
	c.SetLineWidth(10.0)

	c.MoveTo(128.0, 25.6)
	c.LineTo(230.4, 230.4)
	c.RelLineTo(-102.4, 0.0)
	c.CurveTo(51.2, 230.4, 51.2, 128.0, 128.0, 128.0)

	c.Stroke()

	return nil
}

func ExampleFillAndStroke2(c *cairo.Canvas) error {

	c.MoveTo(128.0, 25.6)
	c.LineTo(230.4, 230.4)
	c.RelLineTo(-102.4, 0.0)
	c.CurveTo(51.2, 230.4, 51.2, 128.0, 128.0, 128.0)
	c.ClosePath()

	c.MoveTo(64.0, 25.6)
	c.RelLineTo(51.2, 51.2)
	c.RelLineTo(-51.2, 51.2)
	c.RelLineTo(-51.2, -51.2)
	c.ClosePath()

	c.SetLineWidth(10.0)
	c.SetSourceRGB(0, 0, 1)
	c.FillPreserve()
	c.SetSourceRGB(0, 0, 0)
	c.Stroke()

	return nil
}

func ExampleFillStyle(c *cairo.Canvas) error {

	c.SetLineWidth(6)

	c.Rectangle(12, 12, 232, 70)
	c.NewSubPath()
	c.Arc(64, 64, 40, 0, 2*math.Pi)
	c.NewSubPath()
	c.ArcNegative(192, 64, 40, 0, -2*math.Pi)

	c.SetFillRule(cairo.FILL_RULE_EVEN_ODD)
	c.SetSourceRGB(0, 0.7, 0)
	c.FillPreserve()
	c.SetSourceRGB(0, 0, 0)
	c.Stroke()

	c.Translate(0, 128)
	c.Rectangle(12, 12, 232, 70)
	c.NewSubPath()
	c.Arc(64, 64, 40, 0, 2*math.Pi)
	c.NewSubPath()
	c.ArcNegative(192, 64, 40, 0, -2*math.Pi)

	c.SetFillRule(cairo.FILL_RULE_WINDING)
	c.SetSourceRGB(0, 0, 0.9)
	c.FillPreserve()
	c.SetSourceRGB(0, 0, 0)
	c.Stroke()

	return nil
}

func ExampleImage(c *cairo.Canvas) error {

	image, err := cairo.NewSurfaceFromPNG(textureFile)
	if err != nil {
		return err
	}
	defer image.Destroy()

	w := float64(image.GetWidth())
	h := float64(image.GetHeight())

	c.Translate(128.0, 128.0)
	c.Rotate(45 * math.Pi / 180)
	c.Scale(256.0/w, 256.0/h)
	c.Translate(-0.5*w, -0.5*h)

	c.SetSourceSurface(image, 0, 0)
	c.Paint()

	return nil
}

func ExampleImagePattern(c *cairo.Canvas) error {

	image, err := cairo.NewSurfaceFromPNG(textureFile)
	if err != nil {
		return err
	}
	defer image.Destroy()

	w := float64(image.GetWidth())
	h := float64(image.GetHeight())

	pattern, _ := cairo.NewPatternForSurface(image)
	defer pattern.Destroy()

	pattern.SetExtend(cairo.EXTEND_REPEAT)

	c.Translate(128.0, 128.0)
	c.Rotate(math.Pi / 4)
	c.Scale(1/math.Sqrt(2), 1/math.Sqrt(2))
	c.Translate(-128.0, -128.0)

	matrix := cairo.NewMatrix()

	matrix.InitScale(w/256.0*5.0, h/256.0*5.0)
	pattern.SetMatrix(matrix)

	c.SetSource(pattern)

	c.Rectangle(0, 0, 256.0, 256.0)
	c.Fill()

	return nil
}

func ExampleMultiSegmentCaps(c *cairo.Canvas) error {

	c.MoveTo(50.0, 75.0)
	c.LineTo(200.0, 75.0)

	c.MoveTo(50.0, 125.0)
	c.LineTo(200.0, 125.0)

	c.MoveTo(50.0, 175.0)
	c.LineTo(200.0, 175.0)

	c.SetLineWidth(30.0)
	c.SetLineCap(cairo.LINE_CAP_ROUND)
	c.Stroke()

	return nil
}

func ExampleRoundedRectangle(c *cairo.Canvas) error {

	// a custom shape that could be wrapped in a function
	var (
		x             float64 = 25.6 // parameters like cairo_rectangle
		y             float64 = 25.6
		width         float64 = 204.8
		height        float64 = 204.8
		aspect        float64 = 1.0           // aspect ratio
		corner_radius float64 = height / 10.0 // and corner curvature radius
		radius        float64 = corner_radius / aspect
		degrees       float64 = math.Pi / 180.0
	)

	c.NewSubPath()
	c.Arc(x+width-radius, y+radius, radius, -90*degrees, 0*degrees)
	c.Arc(x+width-radius, y+height-radius, radius, 0*degrees, 90*degrees)
	c.Arc(x+radius, y+height-radius, radius, 90*degrees, 180*degrees)
	c.Arc(x+radius, y+radius, radius, 180*degrees, 270*degrees)
	c.ClosePath()

	c.SetSourceRGB(0.5, 0.5, 1)
	c.FillPreserve()
	c.SetSourceRGBA(0.5, 0, 0, 0.5)
	c.SetLineWidth(10.0)
	c.Stroke()

	return nil
}

func ExampleSetLineCap(c *cairo.Canvas) error {

	c.SetLineWidth(30.0)
	c.SetLineCap(cairo.LINE_CAP_BUTT) // default
	c.MoveTo(64.0, 50.0)
	c.LineTo(64.0, 200.0)
	c.Stroke()
	c.SetLineCap(cairo.LINE_CAP_ROUND)
	c.MoveTo(128.0, 50.0)
	c.LineTo(128.0, 200.0)
	c.Stroke()
	c.SetLineCap(cairo.LINE_CAP_SQUARE)
	c.MoveTo(192.0, 50.0)
	c.LineTo(192.0, 200.0)
	c.Stroke()

	// draw helping lines
	c.SetSourceRGB(1, 0.2, 0.2)
	c.SetLineWidth(2.56)
	c.MoveTo(64.0, 50.0)
	c.LineTo(64.0, 200.0)
	c.MoveTo(128.0, 50.0)
	c.LineTo(128.0, 200.0)
	c.MoveTo(192.0, 50.0)
	c.LineTo(192.0, 200.0)
	c.Stroke()

	return nil
}

func ExampleText(c *cairo.Canvas) error {

	c.SelectFontFace("Sans", cairo.FONT_SLANT_NORMAL, cairo.FONT_WEIGHT_BOLD)
	c.SetFontSize(90.0)

	c.MoveTo(10.0, 135.0)
	c.ShowText("Hello")

	c.MoveTo(70.0, 165.0)
	c.TextPath("void")
	c.SetSourceRGB(0.5, 0.5, 1)
	c.FillPreserve()
	c.SetSourceRGB(0, 0, 0)
	c.SetLineWidth(2.56)
	c.Stroke()

	// draw helping lines
	c.SetSourceRGBA(1, 0.2, 0.2, 0.6)
	c.Arc(10.0, 135.0, 5.12, 0, 2*math.Pi)
	c.ClosePath()
	c.Arc(70.0, 165.0, 5.12, 0, 2*math.Pi)
	c.Fill()

	return nil
}

func ExampleTextAlignCenter(c *cairo.Canvas) error {

	var extents cairo.TextExtents

	text := "cairo"
	var x, y float64

	c.SelectFontFace("Sans", cairo.FONT_SLANT_NORMAL, cairo.FONT_WEIGHT_NORMAL)

	c.SetFontSize(52.0)
	c.TextExtents(text, &extents)
	x = 128.0 - (extents.Width/2 + extents.BearingX)
	y = 128.0 - (extents.Height/2 + extents.BearingY)

	c.MoveTo(x, y)
	c.ShowText(text)

	// draw helping lines
	c.SetSourceRGBA(1, 0.2, 0.2, 0.6)
	c.SetLineWidth(6.0)
	c.Arc(x, y, 10.0, 0, 2*math.Pi)
	c.Fill()
	c.MoveTo(128.0, 0)
	c.RelLineTo(0, 256)
	c.MoveTo(0, 128.0)
	c.RelLineTo(256, 0)
	c.Stroke()

	return nil
}

func ExampleTextExtents(c *cairo.Canvas) error {

	var extents cairo.TextExtents

	text := "cairo"
	var x, y float64

	c.SelectFontFace("Sans", cairo.FONT_SLANT_NORMAL, cairo.FONT_WEIGHT_NORMAL)

	c.SetFontSize(100.0)
	c.TextExtents(text, &extents)

	x = 25.0
	y = 150.0

	c.MoveTo(x, y)
	c.ShowText(text)

	// draw helping lines
	c.SetSourceRGBA(1, 0.2, 0.2, 0.6)
	c.SetLineWidth(6.0)
	c.Arc(x, y, 10.0, 0, 2*math.Pi)
	c.Fill()
	c.MoveTo(x, y)
	c.RelLineTo(0, -extents.Height)
	c.RelLineTo(extents.Width, 0)
	c.RelLineTo(extents.BearingX, -extents.BearingY)
	c.Stroke()

	return nil
}
