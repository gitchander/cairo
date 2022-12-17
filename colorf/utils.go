package colorf

// ------------------------------------------------------------------------------
func clamp(a float64, min, max float64) float64 {
	if a < min {
		a = min
	}
	if a > max {
		a = max
	}
	return a
}

// clamp color component
func clamp01(a float64) float64 {
	return clamp(a, 0, 1)
}

// ------------------------------------------------------------------------------
// Lerp - Linear interpolation
// t= [0, 1]
// (t == 0) => v0
// (t == 1) => v1
func lerp(v0, v1 float64, t float64) float64 {
	return (1.0-t)*v0 + t*v1
}

// ------------------------------------------------------------------------------
// https://en.wikipedia.org/wiki/Nibble
func byteToNibbles(b byte) (hi, lo byte) {
	hi = b >> 4
	lo = b & 0xf
	return
}

func nibblesToByte(hi, lo byte) (b byte) {
	b |= hi << 4
	b |= lo & 0xf
	return
}

func decodeNibbles(bs []byte) ([]byte, bool) {
	ns := make([]byte, len(bs))
	for i, b := range bs {
		n, ok := decodeNibble(b)
		if !ok {
			return nil, false
		}
		ns[i] = n
	}
	return ns, true
}

func decodeNibble(b byte) (byte, bool) {
	if ('0' <= b) && (b <= '9') {
		return b - '0', true
	}
	if ('a' <= b) && (b <= 'f') {
		return b - 'a' + 10, true
	}
	if ('A' <= b) && (b <= 'F') {
		return b - 'A' + 10, true
	}
	return 0, false
}

// ------------------------------------------------------------------------------
