package main

import (
	"errors"
	"fmt"
	"image/color"
	"math"
	"os"
	"path/filepath"

	"github.com/gitchander/cairo"
	"github.com/gitchander/cairo/examples/mathf"
)

type Range struct {
	Min   float64
	Max   float64
	Count int
}

func (r *Range) Step() float64 {
	return (r.Max - r.Min) / float64(r.Count-1)
}

type PolarCurve struct {
	Curve  Curve
	Ranges []Range
}

var curves = []PolarCurve{
	PolarCurve{
		Curve: NewSpiral("spiral", 7.0),
		Ranges: []Range{
			Range{Min: 0, Max: math.Pi * 10, Count: 500},
		},
	},
	PolarCurve{
		Curve: NewCardioid("cardioid", 100.0),
		Ranges: []Range{
			Range{Min: 0, Max: math.Pi * 2, Count: 100},
		},
	},
	PolarCurve{
		Curve: NewLemniscate("lemniscate", 200.0),
		Ranges: []Range{
			Range{Min: -math.Pi / 4, Max: math.Pi / 4, Count: 50},
			Range{Min: math.Pi * 3 / 4, Max: math.Pi * 5 / 4, Count: 50},
		},
	},
	PolarCurve{
		Curve: NewCannabis("cannabis", -50.0),
		Ranges: []Range{
			Range{Min: 0, Max: 2.0 * math.Pi, Count: 1000},
		},
	},
	PolarCurve{
		Curve: NewRose("cardioid-rose", 200.0, 0.5),
		Ranges: []Range{
			Range{Min: 0, Max: 2.0 * math.Pi, Count: 100},
		},
	},
	PolarCurve{
		Curve: NewRose("rose3", 200.0, 3.0),
		Ranges: []Range{
			Range{Min: 0, Max: math.Pi, Count: 100},
		},
	},
	PolarCurve{
		Curve: NewRose("rose2", 200.0, 2.0),
		Ranges: []Range{
			Range{Min: 0, Max: 2.0 * math.Pi, Count: 100},
		},
	},
	PolarCurve{
		Curve: NewCircle("circle", -100.0),
		Ranges: []Range{
			Range{Min: 0, Max: 2.0 * math.Pi, Count: 100},
		},
	},
	PolarCurve{
		Curve: NewStrofoid("strofoid", -100.0),
		Ranges: []Range{
			Range{Min: math.Pi - 2.2, Max: math.Pi + 2.2, Count: 100},
		},
	},
	PolarCurve{
		Curve: NewStrofoidKnot("strofoid-knot", -200.0),
		Ranges: []Range{
			Range{Min: math.Pi - 2.0, Max: math.Pi + 2.0, Count: 100},
		},
	},
	PolarCurve{
		Curve: NewParabolaKnot("parabola-knot", -200.0),
		Ranges: []Range{
			Range{Min: -0.37 * math.Pi, Max: +0.37 * math.Pi, Count: 100},
		},
	},
}

func mkDir(dir string) error {

	fi, err := os.Stat(dir)
	if err != nil {
		err = os.Mkdir(dir, os.ModePerm)
		if err != nil {
			return err
		}
	} else {
		if !fi.IsDir() {
			return errors.New("file is not dir")
		}
	}
	return nil
}

func makeCurve(dir string, params PolarCurve) error {

	width, height := 512, 512

	surface, err := cairo.NewSurface(cairo.FORMAT_ARGB32, width, height)
	if err != nil {
		return err
	}

	c, err := cairo.NewCanvas(surface)
	if err != nil {
		return err
	}

	c.FillColor(color.White)

	c.SetLineJoin(cairo.LINE_JOIN_ROUND)
	c.SetLineWidth(1.0)
	c.SetSourceRGB(0.7, 0.7, 0.7)
	DrawAxes(c)

	c.SetLineWidth(2)
	c.SetSourceRGB(0.5, 0, 0)
	PolarDraw(c, width, height, params)

	fileName := filepath.Join(dir, fmt.Sprintf("%s.png", params.Curve.Name()))

	if err = surface.WriteToPNG(fileName); err != nil {
		return err
	}
	surface.Finish()

	return nil
}

func DrawAxes(canvas *cairo.Canvas) {

	surface := canvas.GetTarget()
	if surface == nil {
		return
	}

	var (
		x0 = float64(surface.GetWidth()) * 0.5
		y0 = float64(surface.GetHeight()) * 0.5
	)

	const rd = float64(40)
	k := 6
	m := 80
	du := 2 * math.Pi / float64(m-1)

	// draw circles
	for i := 0; i < k; i++ {

		u := float64(0)
		for j := 0; j < m; j++ {

			s, c := math.Sincos(u)

			r := rd * float64(i+1)
			x := x0 + r*c
			y := y0 + r*s

			if j == 0 {
				canvas.MoveTo(x, y)
			} else {
				canvas.LineTo(x, y)
			}

			u += du
		}
	}

	n := 16
	du = 2 * math.Pi / float64(n)
	u := float64(0)
	k1 := 0
	k2 := k

	// draw rays
	for i := 0; i < n; i++ {

		s, c := math.Sincos(u)

		r := rd * float64(k1)
		x1 := x0 + r*c
		y1 := y0 + r*s

		r = rd * float64(k2)
		x2 := x0 + r*c
		y2 := y0 + r*s

		canvas.MoveTo(x1, y1)
		canvas.LineTo(x2, y2)

		u += du
	}

	// render center cross
	if false {
		canvas.MoveTo(x0, y0-rd)
		canvas.LineTo(x0, y0+rd)

		canvas.MoveTo(x0-rd, y0)
		canvas.LineTo(x0+rd, y0)
	}

	canvas.Stroke()
}

func PolarDraw(canvas *cairo.Canvas, width, height int, params PolarCurve) {

	var center = mathf.Point2f{
		X: float64(width) * 0.5,
		Y: float64(height) * 0.5,
	}

	for _, r := range params.Ranges {

		angleStep := r.Step()
		n := r.Count
		angle := r.Min

		for i := 0; i < n; i++ {

			radius := params.Curve.RadiusByAngle(angle)
			d := mathf.Pt2f(mathf.PolarToCartesian(radius, angle))
			angle += angleStep

			temp := center.Add(d)

			x, y := temp.X, temp.Y

			if i == 0 {
				canvas.MoveTo(x, y)
			} else {
				canvas.LineTo(x, y)
			}
		}

		canvas.Stroke()
	}
}

func main() {

	dir := filepath.Dir(os.Args[0])
	dir = filepath.Join(dir, "result")

	err := mkDir(dir)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	for _, curve := range curves {
		err := makeCurve(dir, curve)
		if err != nil {
			fmt.Println(err.Error())
			break
		}
	}
}
