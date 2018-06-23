package colorf

import (
	"errors"
)

const (
	encodeFactor = 255.0
	decodeFactor = 1.0 / encodeFactor
)

type Coder interface {
	Size() int
	Encode(bs []byte, c RGBA) (err error)
	Decode(bs []byte) (c RGBA, err error)
}

var CoderBGRA32 = coderBGRA32{}

type coderBGRA32 struct{}

// Encode Size
func (coderBGRA32) Size() int {
	return 4
}

func (coderBGRA32) Encode(bs []byte, c RGBA) error {

	if len(bs) < CoderBGRA32.Size() {
		return errors.New("ColorRGBA.Encode(): wrong data size")
	}

	bs[0] = byte(round(c.B * encodeFactor))
	bs[1] = byte(round(c.G * encodeFactor))
	bs[2] = byte(round(c.R * encodeFactor))
	bs[3] = byte(round(c.A * encodeFactor))

	return nil
}

func (coderBGRA32) Decode(bs []byte) (c RGBA, err error) {

	if len(bs) < CoderBGRA32.Size() {
		err = errors.New("ColorRGBA.Decode(): wrong data size")
		return
	}

	return RGBA{
		B: float64(bs[0]) * decodeFactor,
		G: float64(bs[1]) * decodeFactor,
		R: float64(bs[2]) * decodeFactor,
		A: float64(bs[3]) * decodeFactor,
	}, nil
}
