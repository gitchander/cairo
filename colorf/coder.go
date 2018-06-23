package colorf

import (
	"errors"
	"image/color"
)

type Coder interface {
	Size() int
	Encode(bs []byte, c color.Color) error
	Decode(bs []byte) (color.Color, error)
}

var CoderBGRA32 = coderBGRA32{}

type coderBGRA32 struct{}

var _ Coder = coderBGRA32{}

// Encode Size
func (coderBGRA32) Size() int {
	return 4
}

func (coderBGRA32) Encode(bs []byte, c color.Color) error {

	if len(bs) < CoderBGRA32.Size() {
		return errors.New("Insufficient data len")
	}

	v := color.RGBAModel.Convert(c).(color.RGBA)

	bs[0] = v.B
	bs[1] = v.G
	bs[2] = v.R
	bs[3] = v.A

	return nil
}

func (coderBGRA32) Decode(bs []byte) (color.Color, error) {

	if len(bs) < CoderBGRA32.Size() {
		return nil, errors.New("Insufficient data len")
	}

	return color.RGBA{
		B: bs[0],
		G: bs[1],
		R: bs[2],
		A: bs[3],
	}, nil
}
