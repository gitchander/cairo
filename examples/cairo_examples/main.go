package main

import (
	"image"
)

func main() {
	err := Run()
	checkError(err)
}

func Run() error {

	var (
		size = image.Point{X: 256, Y: 256}
		dir  = "result"
	)

	err := makeDir(dir)
	if err != nil {
		return err
	}

	es := []Example{
		Example{ExampleHelloWorld, dir, "hello-world.png", size},
		Example{ExampleArc, dir, "arc.png", size},
		Example{ExampleArcNegative, dir, "arc-negative.png", size},
		Example{ExampleClip, dir, "clip.png", size},
		Example{ExampleClipImage, dir, "clip-image.png", size},
		Example{ExampleCurveRectangle, dir, "curve-rectangle.png", size},
		Example{ExampleCurveTo, dir, "curve-to.png", size},
		Example{ExampleGradient, dir, "gradient.png", size},
		Example{ExampleSetLineJoin, dir, "set-line-join.png", size},
		Example{ExampleDonut, dir, "donut.png", size},
		Example{ExampleDash, dir, "dash.png", size},
		Example{ExampleFillAndStroke2, dir, "fill-and-stroke2.png", size},
		Example{ExampleFillStyle, dir, "fill-style.png", size},
		Example{ExampleImage, dir, "image.png", size},
		Example{ExampleImagePattern, dir, "image-pattern.png", size},
		Example{ExampleMultiSegmentCaps, dir, "multi-segment-caps.png", size},
		Example{ExampleRoundedRectangle, dir, "rounded-rectangle.png", size},
		Example{ExampleSetLineCap, dir, "set-line-cap.png", size},
		Example{ExampleText, dir, "text.png", size},
		Example{ExampleTextAlignCenter, dir, "text-align-center.png", size},
		Example{ExampleTextExtents, dir, "text-extents.png", size},
	}

	for _, e := range es {
		err = e.Execute()
		if err != nil {
			return err
		}
	}

	return nil
}
