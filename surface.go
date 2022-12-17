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
	surfaceNative *C.cairo_surface_t
}

func newSurface(surfaceNative *C.cairo_surface_t) (*Surface, error) {

	err := checkCairoStatus(C.cairo_surface_status(surfaceNative))
	if err != nil {
		return nil, err
	}

	s := &Surface{surfaceNative}

	runtime.SetFinalizer(s, (*Surface).destroy)

	return s, nil
}

func (s *Surface) destroy() {
	C.cairo_surface_destroy(s.surfaceNative)
}

func (s *Surface) Destroy() {

	if s.surfaceNative == nil {
		return
	}
	s.destroy()
	s.surfaceNative = nil

	runtime.SetFinalizer(s, nil)
}

func NewSurface(format Format, width, height int) (*Surface, error) {

	surfaceNative := C.cairo_image_surface_create(C.cairo_format_t(format), C.int(width), C.int(height))

	return newSurface(surfaceNative)
}

func NewSurfaceFromPNG(fileName string) (*Surface, error) {

	cstr := newCString(fileName)
	defer freeCString(cstr)

	surfaceNative := C.cairo_image_surface_create_from_png(cstr)

	return newSurface(surfaceNative)
}

func NewSurfaceNative(ptr uintptr) (*Surface, error) {

	surfaceNative := (*C.cairo_surface_t)(unsafe.Pointer(ptr))
	reference := C.cairo_surface_reference(surfaceNative)

	return newSurface(reference)
}

func CreateSurfaceForData(data []byte, format Format, width, height, stride int) (*Surface, error) {
	surfaceNative := C.cairo_image_surface_create_for_data(
		(*C.uchar)(&data[0]),
		C.cairo_format_t(format),
		C.int(width),
		C.int(height),
		C.int(stride),
	)
	return newSurface(surfaceNative)
}

func (s *Surface) Native() uintptr {
	return uintptr(unsafe.Pointer(s.surfaceNative))
}

func (s *Surface) Reference() *Surface {

	reference := C.cairo_surface_reference(s.surfaceNative)

	sr, _ := newSurface(reference)

	return sr
}

func (s *Surface) Finish() {
	C.cairo_surface_finish(s.surfaceNative)
}

func (s *Surface) WriteToPNG(fileName string) error {

	cstr := newCString(fileName)
	defer freeCString(cstr)

	return checkCairoStatus(C.cairo_surface_write_to_png(s.surfaceNative, cstr))
}

func (s *Surface) GetFormat() Format {
	return Format(C.cairo_image_surface_get_format(s.surfaceNative))
}

func (s *Surface) GetWidth() int {
	return int(C.cairo_image_surface_get_width(s.surfaceNative))
}

func (s *Surface) GetHeight() int {
	return int(C.cairo_image_surface_get_height(s.surfaceNative))
}

func (s *Surface) GetStride() int {
	return int(C.cairo_image_surface_get_stride(s.surfaceNative))
}

func (s *Surface) Flush() {
	C.cairo_surface_flush(s.surfaceNative)
}

func (s *Surface) MarkDirty() {
	C.cairo_surface_mark_dirty(s.surfaceNative)
}

func (s *Surface) GetDataLength() int {

	stride := s.GetStride()
	height := s.GetHeight()

	return stride * height
}

func (s *Surface) GetData(data []byte) error {

	dataLen := s.GetDataLength()
	if len(data) != dataLen {
		return newCairoError("Surface.GetData(): invalid data size")
	}

	dataPtr := unsafe.Pointer(C.cairo_image_surface_get_data(s.surfaceNative))
	if dataPtr == nil {
		return newCairoError("Surface.GetData(): can't access surface pixel data")
	}

	C.memcpy(unsafe.Pointer(&data[0]), dataPtr, C.size_t(dataLen))

	return nil
}

func (s *Surface) SetData(data []byte) error {

	dataLen := s.GetDataLength()
	if len(data) != dataLen {
		return newCairoError("Surface.SetData(): invalid data size")
	}

	s.Flush()

	dataPtr := unsafe.Pointer(C.cairo_image_surface_get_data(s.surfaceNative))
	if dataPtr == nil {
		return newCairoError("Surface.SetData(): can't access surface pixel data")
	}

	C.memcpy(dataPtr, unsafe.Pointer(&data[0]), C.size_t(dataLen))

	s.MarkDirty()

	return nil
}
