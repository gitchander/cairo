package cairo

// #cgo pkg-config: cairo cairo-gobject
// #include <stdlib.h>
// #include <cairo.h>
// #include <cairo-gobject.h>
import "C"

import (
	"runtime"
	"unsafe"
)

type Pattern struct {
	pattern_n *C.cairo_pattern_t
}

func newPattern(pattern_n *C.cairo_pattern_t) (*Pattern, error) {

	err := checkCairoStatus(C.cairo_pattern_status(pattern_n))
	if err != nil {
		return nil, err
	}

	p := &Pattern{pattern_n}

	runtime.SetFinalizer(p, (*Pattern).destroy)

	return p, nil
}

func (p *Pattern) destroy() {
	C.cairo_pattern_destroy(p.pattern_n)
}

func (p *Pattern) Destroy() {

	if p.pattern_n == nil {
		return
	}
	p.destroy()
	p.pattern_n = nil

	runtime.SetFinalizer(p, nil)
}

func NewPatternLinear(x0, y0, x1, y1 float64) (*Pattern, error) {

	pattern_n := C.cairo_pattern_create_linear(C.double(x0), C.double(y0), C.double(x1), C.double(y1))

	return newPattern(pattern_n)
}

func NewPatternRadial(cx0, cy0, radius0, cx1, cy1, radius1 float64) (*Pattern, error) {

	pattern_n := C.cairo_pattern_create_radial(
		C.double(cx0), C.double(cy0), C.double(radius0),
		C.double(cx1), C.double(cy1), C.double(radius1))

	return newPattern(pattern_n)
}

func NewPatternForSurface(s *Surface) (*Pattern, error) {

	pattern_n := C.cairo_pattern_create_for_surface(s.surfaceNative)

	return newPattern(pattern_n)
}

func NewPatternNative(ptr uintptr) (*Pattern, error) {

	pattern_n := (*C.cairo_pattern_t)(unsafe.Pointer(ptr))
	reference := C.cairo_pattern_reference(pattern_n)

	return newPattern(reference)
}

func (p *Pattern) Native() uintptr {
	return uintptr(unsafe.Pointer(p.pattern_n))
}

func (p *Pattern) Reference() *Pattern {

	reference := C.cairo_pattern_reference(p.pattern_n)

	pr, _ := newPattern(reference)

	return pr
}

func (p *Pattern) AddColorStopRGB(offset, red, green, blue float64) {
	C.cairo_pattern_add_color_stop_rgb(p.pattern_n, C.double(offset), C.double(red), C.double(green), C.double(blue))
}

func (p *Pattern) AddColorStopRGBA(offset, red, green, blue, alpha float64) {
	C.cairo_pattern_add_color_stop_rgba(p.pattern_n, C.double(offset), C.double(red), C.double(green), C.double(blue), C.double(alpha))
}

func (p *Pattern) SetExtend(extend Extend) {
	C.cairo_pattern_set_extend(p.pattern_n, C.cairo_extend_t(extend))
}

func (p *Pattern) SetMatrix(m *Matrix) {
	C.cairo_pattern_set_matrix(p.pattern_n, m.matrixNative)
}
