package cairo

// #include <cairo.h>
import "C"

type Matrix struct {
	matrixNative *C.cairo_matrix_t
}

func NewMatrix() *Matrix {
	return &Matrix{
		matrixNative: &C.cairo_matrix_t{},
	}
}

func (m *Matrix) Init(xx, yx, xy, yy, x0, y0 float64) {
	C.cairo_matrix_init(m.matrixNative,
		C.double(xx), C.double(yx),
		C.double(xy), C.double(yy),
		C.double(x0), C.double(y0),
	)
}

func (m *Matrix) InitIdendity() {
	C.cairo_matrix_init_identity(m.matrixNative)
}

func (m *Matrix) InitTranslate(tx, ty float64) {
	C.cairo_matrix_init_translate(m.matrixNative, C.double(tx), C.double(ty))
}

func (m *Matrix) InitScale(sx, sy float64) {
	C.cairo_matrix_init_scale(m.matrixNative, C.double(sx), C.double(sy))
}

func (m *Matrix) InitRotate(radians float64) {
	C.cairo_matrix_init_rotate(m.matrixNative, C.double(radians))
}

func (m *Matrix) Translate(tx, ty float64) {
	C.cairo_matrix_translate(m.matrixNative, C.double(tx), C.double(ty))
}

func (m *Matrix) Scale(sx, sy float64) {
	C.cairo_matrix_scale(m.matrixNative, C.double(sx), C.double(sy))
}

func (m *Matrix) Rotate(radians float64) {
	C.cairo_matrix_rotate(m.matrixNative, C.double(radians))
}

func (m *Matrix) Invert() Status {
	return Status(C.cairo_matrix_invert(m.matrixNative))
}

func (m *Matrix) Multiply(a, b *Matrix) {
	C.cairo_matrix_multiply(m.matrixNative, a.matrixNative, b.matrixNative)
}

func (m *Matrix) TransformDistance(dx, dy float64) (float64, float64) {
	var (
		x0 = C.double(dx)
		y0 = C.double(dy)
	)
	C.cairo_matrix_transform_distance(m.matrixNative, &x0, &y0)
	return float64(x0), float64(y0)
}

func (m *Matrix) TransformPoint(x, y float64) (float64, float64) {
	var (
		x0 = C.double(x)
		y0 = C.double(y)
	)
	C.cairo_matrix_transform_point(m.matrixNative, &x0, &y0)
	return float64(x0), float64(y0)
}
