package colorf

import (
	"math"
)

const (
	maxColorComponent = (1 << 16) - 1
	// maxColorComponent = 0xffff
)

// It clamps color component.
func clampColorComponent(c float64) float64 {
	return clamp(c, 0, 1)
}

func colorComponentEncode(c float64) uint32 {
	c = clampColorComponent(c)
	return uint32(math.Round(c * maxColorComponent))
}

func colorComponentDecode(u uint32) float64 {
	return float64(u) / maxColorComponent
}
