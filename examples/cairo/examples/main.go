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
		Example{sampleArc, dir, "arc.png", size},
		Example{sampleArcNegative, dir, "arc-negative.png", size},
		Example{sampleClip, dir, "clip.png", size},
		Example{sampleClipImage, dir, "clip-image.png", size},
		Example{sampleCurveRectangle, dir, "curve-rectangle.png", size},
		Example{sampleCurveTo, dir, "curve-to.png", size},
		Example{sampleGradient, dir, "gradient.png", size},
		Example{sampleSetLineJoin, dir, "set-line-join.png", size},
		Example{sampleDash, dir, "dash.png", size},
		Example{sampleFillAndStroke2, dir, "fill-and-stroke2.png", size},
		Example{sampleFillStyle, dir, "fill-style.png", size},
		Example{sampleImage, dir, "image.png", size},
		Example{sampleImagePattern, dir, "image-pattern.png", size},
		Example{sampleMultiSegmentCaps, dir, "multi-segment-caps.png", size},
		Example{sampleRoundedRectangle, dir, "rounded-rectangle.png", size},
		Example{sampleSetLineCap, dir, "set-line-cap.png", size},
		Example{sampleText, dir, "text.png", size},
		Example{sampleTextAlignCenter, dir, "text-align-center.png", size},
		Example{sampleTextExtents, dir, "text-extents.png", size},
	}

	for _, e := range es {
		err = e.Execute()
		if err != nil {
			return err
		}
	}

	return nil
}
