package cairo

// #cgo pkg-config: cairo cairo-gobject
// #include <stdlib.h>
// #include <cairo.h>
// #include <cairo-gobject.h>
import "C"

import (
	"unsafe"
)

func boolCairo(b bool) C.cairo_bool_t {
	if b {
		return C.cairo_bool_t(1)
	}
	return C.cairo_bool_t(0)
}

func boolGolang(b C.cairo_bool_t) bool {
	if b != 0 {
		return true
	}
	return false
}

func newCString(s string) *C.char {
	return C.CString(s)
}

func freeCString(p *C.char) {
	C.free(unsafe.Pointer(p))
}
