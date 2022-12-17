package colorf

import (
	"fmt"
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

func checkSizeBGRA32(bs []byte) error {
	var (
		haveSize = len(bs)
		wantSize = coderBGRA32{}.Size()
	)
	if haveSize < wantSize {
		return fmt.Errorf("Insufficient data length: have %d, want %d", haveSize, wantSize)
	}
	return nil
}

func (coderBGRA32) Encode(bs []byte, c color.Color) error {

	if err := checkSizeBGRA32(bs); err != nil {
		return err
	}

	v := color.NRGBAModel.Convert(c).(color.NRGBA)

	bs[0] = v.B
	bs[1] = v.G
	bs[2] = v.R
	bs[3] = v.A

	return nil
}

func (coderBGRA32) Decode(bs []byte) (color.Color, error) {

	if err := checkSizeBGRA32(bs); err != nil {
		return nil, err
	}

	c := color.NRGBA{
		B: bs[0],
		G: bs[1],
		R: bs[2],
		A: bs[3],
	}

	return c, nil
}
