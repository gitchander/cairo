package cairo

// #cgo pkg-config: cairo cairo-gobject
// #include <stdlib.h>
// #include <string.h>
// #include <cairo.h>
// #include <cairo-gobject.h>
import "C"

import (
	"runtime"
	"unsafe"
)

type Surface struct {
	surface_n *C.cairo_surface_t
}

func newSurface(surface_n *C.cairo_surface_t) (*Surface, error) {

	status := Status(C.cairo_surface_status(surface_n))
	if status != STATUS_SUCCESS {
		return nil, NewErrorFromStatus(status)
	}

	s := &Surface{surface_n}

	runtime.SetFinalizer(s, (*Surface).destroy)

	return s, nil
}

func (s *Surface) destroy() {
	C.cairo_surface_destroy(s.surface_n)
}

func (s *Surface) Destroy() {

	if s.surface_n == nil {
		return
	}
	s.destroy()
	s.surface_n = nil

	runtime.SetFinalizer(s, nil)
}

func NewSurface(format Format, width, height int) (*Surface, error) {

	surface_n := C.cairo_image_surface_create(C.cairo_format_t(format), C.int(width), C.int(height))

	return newSurface(surface_n)
}

func NewSurfaceFromPNG(fileName string) (*Surface, error) {

	cstr := newCString(fileName)
	defer freeCString(cstr)

	surface_n := C.cairo_image_surface_create_from_png(cstr)

	return newSurface(surface_n)
}

func NewSurfaceNative(ptr uintptr) (*Surface, error) {

	surface_n := (*C.cairo_surface_t)(unsafe.Pointer(ptr))
	reference := C.cairo_surface_reference(surface_n)

	return newSurface(reference)
}

func (s *Surface) Native() uintptr {
	return uintptr(unsafe.Pointer(s.surface_n))
}

func (s *Surface) Reference() *Surface {

	reference := C.cairo_surface_reference(s.surface_n)

	sr, _ := newSurface(reference)

	return sr
}

func (s *Surface) Finish() {
	C.cairo_surface_finish(s.surface_n)
}

func (s *Surface) WriteToPNG(fileName string) error {

	cstr := newCString(fileName)
	defer freeCString(cstr)

	status := Status(C.cairo_surface_write_to_png(s.surface_n, cstr))
	if status != STATUS_SUCCESS {
		return NewErrorFromStatus(status)
	}

	return nil
}

func (s *Surface) GetFormat() Format {
	return Format(C.cairo_image_surface_get_format(s.surface_n))
}

func (s *Surface) GetWidth() int {
	return int(C.cairo_image_surface_get_width(s.surface_n))
}

func (s *Surface) GetHeight() int {
	return int(C.cairo_image_surface_get_height(s.surface_n))
}

func (s *Surface) GetStride() int {
	return int(C.cairo_image_surface_get_stride(s.surface_n))
}

func (s *Surface) Flush() {
	C.cairo_surface_flush(s.surface_n)
}

func (s *Surface) MarkDirty() {
	C.cairo_surface_mark_dirty(s.surface_n)
}

func (s *Surface) GetDataLength() int {

	stride := s.GetStride()
	height := s.GetHeight()

	return stride * height
}

func (s *Surface) GetData(data []byte) error {

	dataLen := s.GetDataLength()
	if len(data) != dataLen {
		return newError("Surface.GetData(): invalid data size")
	}

	dataPtr := unsafe.Pointer(C.cairo_image_surface_get_data(s.surface_n))
	if dataPtr == nil {
		return newError("Surface.GetData(): can't access surface pixel data")
	}

	C.memcpy(unsafe.Pointer(&data[0]), dataPtr, C.size_t(dataLen))

	return nil
}

func (s *Surface) SetData(data []byte) error {

	dataLen := s.GetDataLength()
	if len(data) != dataLen {
		return newError("Surface.SetData(): invalid data size")
	}

	s.Flush()

	dataPtr := unsafe.Pointer(C.cairo_image_surface_get_data(s.surface_n))
	if dataPtr == nil {
		return newError("Surface.SetData(): can't access surface pixel data")
	}

	C.memcpy(dataPtr, unsafe.Pointer(&data[0]), C.size_t(dataLen))

	s.MarkDirty()

	return nil
}
