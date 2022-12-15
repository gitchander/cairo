package colorf

import (
	"fmt"
	"image/color"
)

func ParseColor(s string) (color.Color, error) {

	bs := []byte(s)
	if (len(bs) == 0) || (bs[0] != '#') {
		return nil, fmt.Errorf("invalid color (%s): no symbol %q", s, '#')
	}
	bs = bs[1:]

	ns, ok := decodeNibbles(bs)
	if !ok {
		return nil, fmt.Errorf("invalid color (%s)", s)
	}

	var c color.Color
	switch len(ns) {
	case 3: // rgb
		c = color.NRGBA{
			R: nibblesToByte(ns[0], ns[0]),
			G: nibblesToByte(ns[1], ns[1]),
			B: nibblesToByte(ns[2], ns[2]),
			A: 0xff,
		}
	case 4: // rgba
		c = color.NRGBA{
			R: nibblesToByte(ns[0], ns[0]),
			G: nibblesToByte(ns[1], ns[1]),
			B: nibblesToByte(ns[2], ns[2]),
			A: nibblesToByte(ns[3], ns[3]),
		}
	case 6: // rrggbb
		c = color.NRGBA{
			R: nibblesToByte(ns[0], ns[1]),
			G: nibblesToByte(ns[2], ns[3]),
			B: nibblesToByte(ns[4], ns[5]),
			A: 0xff,
		}
	case 8: // rrggbbaa
		c = color.NRGBA{
			R: nibblesToByte(ns[0], ns[1]),
			G: nibblesToByte(ns[2], ns[3]),
			B: nibblesToByte(ns[4], ns[5]),
			A: nibblesToByte(ns[6], ns[7]),
		}
	default:
		return nil, fmt.Errorf("invalid color (%s)", s)
	}
	return c, nil
}

func MustParseColor(s string) color.Color {
	c, err := ParseColor(s)
	if err != nil {
		panic(err)
	}
	return c
}
