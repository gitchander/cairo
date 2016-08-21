package color

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

func NewCoderBGRA32() Coder {
	return &coderBGRA32{}
}

type coderBGRA32 struct{}

// Encode Size
func (this *coderBGRA32) Size() int {
	return 4
}

func (this *coderBGRA32) Encode(bs []byte, c RGBA) error {

	if len(bs) < this.Size() {
		return errors.New("ColorRGBA.Encode(): wrong data size")
	}

	var r, g, b, a = c.GetRGBA()

	bs[0] = byte(round(b * encodeFactor))
	bs[1] = byte(round(g * encodeFactor))
	bs[2] = byte(round(r * encodeFactor))
	bs[3] = byte(round(a * encodeFactor))

	return nil
}

func (this *coderBGRA32) Decode(bs []byte) (c RGBA, err error) {

	if len(bs) < this.Size() {
		err = errors.New("ColorRGBA.Decode(): wrong data size")
		return
	}

	var (
		b = float64(bs[0]) * decodeFactor
		g = float64(bs[1]) * decodeFactor
		r = float64(bs[2]) * decodeFactor
		a = float64(bs[3]) * decodeFactor
	)

	c = NewRGBA(r, g, b, a)

	return
}
