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

type Canvas struct {
	cr *C.cairo_t // cairo context
}

func newCanvas(canvasNative *C.cairo_t) (*Canvas, error) {

	err := checkCairoStatus(C.cairo_status(canvasNative))
	if err != nil {
		return nil, err
	}

	c := &Canvas{canvasNative}

	runtime.SetFinalizer(c, (*Canvas).destroy)

	return c, nil
}

func (c *Canvas) destroy() {
	C.cairo_destroy(c.cr)
}

func (c *Canvas) Destroy() {

	if c.cr == nil {
		return
	}
	c.destroy()
	c.cr = nil

	runtime.SetFinalizer(c, nil)
}

func NewCanvas(s *Surface) (*Canvas, error) {

	canvas_n := C.cairo_create(s.surfaceNative)

	return newCanvas(canvas_n)
}

func NewCanvasNative(ptr uintptr) (*Canvas, error) {

	canvas_n := (*C.cairo_t)(unsafe.Pointer(ptr))
	reference := C.cairo_reference(canvas_n)

	//reference := C.cairo_reference((*C.cairo_t)(unsafe.Pointer(ptr)))

	return newCanvas(reference)
}

// Native returns native cairo context
func (c *Canvas) Native() uintptr {
	return uintptr(unsafe.Pointer(c.cr))
}

func (c *Canvas) Reference() *Canvas {

	reference := C.cairo_reference(c.cr)

	cr, _ := newCanvas(reference)

	return cr
}

func (c *Canvas) Status() Status {
	return Status(C.cairo_status(c.cr))
}

func (c *Canvas) Save() {
	C.cairo_save(c.cr)
}

func (c *Canvas) Restore() {
	C.cairo_restore(c.cr)
}

func (c *Canvas) GetTarget() *Surface {

	var surfaceNative *C.cairo_surface_t
	surfaceNative = C.cairo_get_target(c.cr)
	if surfaceNative == nil {
		return nil
	}
	return &Surface{surfaceNative}
}

func (c *Canvas) PushGroup() {
	C.cairo_push_group(c.cr)
}

func (c *Canvas) PushGroupWithContent(content Content) {
	C.cairo_push_group_with_content(c.cr, C.cairo_content_t(content))
}

// cairo_pop_group ()

// cairo_pop_group_to_source ()

func (c *Canvas) GetGroupTarget() *Surface {
	var surfaceNative *C.cairo_surface_t
	surfaceNative = C.cairo_get_group_target(c.cr)
	if surfaceNative == nil {
		return nil
	}
	return &Surface{surfaceNative}
}

func (c *Canvas) SetSource(p *Pattern) {
	C.cairo_set_source(c.cr, p.pattern_n)
}

func (c *Canvas) SetSourceSurface(s *Surface, x, y float64) {
	C.cairo_set_source_surface(c.cr, s.surfaceNative, C.double(x), C.double(y))
}

func (c *Canvas) GetSource() *Pattern {

	var (
		patternNative    *C.cairo_pattern_t
		patternReference *C.cairo_pattern_t
	)

	patternNative = C.cairo_get_source(c.cr)
	patternReference = C.cairo_pattern_reference(patternNative)

	return &Pattern{patternReference}
}

func (c *Canvas) SetAntialias(antialias Antialias) {
	C.cairo_set_antialias(c.cr, C.cairo_antialias_t(antialias))
}

func (c *Canvas) GetAntialias() Antialias {
	return Antialias(C.cairo_get_antialias(c.cr))
}

func (c *Canvas) SetDash(dashes []float64, offset float64) {

	if len(dashes) == 0 {
		C.cairo_set_dash(c.cr, nil, 0, 0.0)
		return
	}

	numDashes := C.int(len(dashes))
	ptrDashes := (*C.double)(&dashes[0])

	C.cairo_set_dash(c.cr,
		ptrDashes,
		numDashes,
		C.double(offset))
}

func (c *Canvas) GetDashCount() int {
	return int(C.cairo_get_dash_count(c.cr))
}

// cairo_get_dash ()

func (c *Canvas) SetFillRule(fillRule FillRule) {
	C.cairo_set_fill_rule(c.cr, C.cairo_fill_rule_t(fillRule))
}

func (c *Canvas) GetFillRule() FillRule {
	return FillRule(C.cairo_get_fill_rule(c.cr))
}

func (c *Canvas) SetLineCap(lineCap LineCap) {
	C.cairo_set_line_cap(c.cr, C.cairo_line_cap_t(lineCap))
}

func (c *Canvas) GetLineCap() LineCap {
	return LineCap(C.cairo_get_line_cap(c.cr))
}

func (c *Canvas) SetLineJoin(lineJoin LineJoin) {
	C.cairo_set_line_join(c.cr, C.cairo_line_join_t(lineJoin))
}

func (c *Canvas) GetLineJoin() LineJoin {
	return LineJoin(C.cairo_get_line_join(c.cr))
}

func (c *Canvas) SetLineWidth(width float64) {
	C.cairo_set_line_width(c.cr, C.double(width))
}

func (c *Canvas) GetLineWidth() float64 {
	return float64(C.cairo_get_line_width(c.cr))
}

func (c *Canvas) SetMiterLimit(limit float64) {
	C.cairo_set_miter_limit(c.cr, C.double(limit))
}

func (c *Canvas) GetMiterLimit() float64 {
	return float64(C.cairo_get_miter_limit(c.cr))
}

func (c *Canvas) SetOperator(operator Operator) {
	C.cairo_set_operator(c.cr, C.cairo_operator_t(operator))
}

func (c *Canvas) GetOperator() Operator {
	return Operator(C.cairo_get_operator(c.cr))
}

func (c *Canvas) SetTolerance(tolerance float64) {
	C.cairo_set_tolerance(c.cr, C.double(tolerance))
}

func (c *Canvas) GetTolerance() float64 {
	return float64(C.cairo_get_tolerance(c.cr))
}

func (c *Canvas) Clip() {
	C.cairo_clip(c.cr)
}

func (c *Canvas) ClipPreserve() {
	C.cairo_clip_preserve(c.cr)
}

// cairo_clip_extents ()

func (c *Canvas) InClip(x, y float64) bool {
	var b C.cairo_bool_t
	b = C.cairo_in_clip(c.cr, C.double(x), C.double(y))
	return boolGolang(b)
}

func (c *Canvas) ResetClip() {
	C.cairo_reset_clip(c.cr)
}

// cairo_rectangle_list_destroy ()

// cairo_copy_clip_rectangle_list ()

func (c *Canvas) Fill() {
	C.cairo_fill(c.cr)
}

func (c *Canvas) FillPreserve() {
	C.cairo_fill_preserve(c.cr)
}

// cairo_fill_extents ()

func (c *Canvas) InFill(x, y float64) bool {
	var b C.cairo_bool_t
	b = C.cairo_in_fill(c.cr, C.double(x), C.double(y))
	return boolGolang(b)
}

// cairo_mask ()

// cairo_mask_surface ()

func (c *Canvas) Paint() {
	C.cairo_paint(c.cr)
}

func (c *Canvas) PaintWithAlpha(alpha float64) {
	C.cairo_paint_with_alpha(c.cr, C.double(alpha))
}

func (c *Canvas) Stroke() {
	C.cairo_stroke(c.cr)
}

func (c *Canvas) StrokePreserve() {
	C.cairo_stroke_preserve(c.cr)
}

// cairo_stroke_extents ()

func (c *Canvas) InStroke(x, y float64) bool {
	var b C.cairo_bool_t
	b = C.cairo_in_stroke(c.cr, C.double(x), C.double(y))
	return boolGolang(b)
}

func (c *Canvas) CopyPage() {
	C.cairo_copy_page(c.cr)
}

func (c *Canvas) ShowPage() {
	C.cairo_show_page(c.cr)
}

func (c *Canvas) GetReferenceCount() uint {
	return uint(C.cairo_get_reference_count(c.cr))
}

// cairo_set_user_data ()

// cairo_get_user_data ()

// ------------------------------------------
func (c *Canvas) MoveTo(x, y float64) {
	C.cairo_move_to(c.cr, C.double(x), C.double(y))
}

func (c *Canvas) LineTo(x, y float64) {
	C.cairo_line_to(c.cr, C.double(x), C.double(y))
}

func (c *Canvas) RelLineTo(dx, dy float64) {
	C.cairo_rel_line_to(c.cr, C.double(dx), C.double(dy))
}

func (c *Canvas) Rectangle(x, y, width, height float64) {
	C.cairo_rectangle(c.cr,
		C.double(x), C.double(y),
		C.double(width), C.double(height))
}

func (c *Canvas) NewPath() {
	C.cairo_new_path(c.cr)
}

func (c *Canvas) NewSubPath() {
	C.cairo_new_sub_path(c.cr)
}

func (c *Canvas) ClosePath() {
	C.cairo_close_path(c.cr)
}

func (c *Canvas) Arc(xc, yc, radius, angle1, angle2 float64) {
	C.cairo_arc(c.cr, C.double(xc), C.double(yc), C.double(radius), C.double(angle1), C.double(angle2))
}

func (c *Canvas) ArcNegative(xc, yc, radius, angle1, angle2 float64) {
	C.cairo_arc_negative(c.cr, C.double(xc), C.double(yc), C.double(radius), C.double(angle1), C.double(angle2))
}

func (c *Canvas) CurveTo(x1, y1, x2, y2, x3, y3 float64) {
	C.cairo_curve_to(c.cr,
		C.double(x1), C.double(y1),
		C.double(x2), C.double(y2),
		C.double(x3), C.double(y3))
}

// Transformations

func (c *Canvas) Scale(sx, sy float64) {
	C.cairo_scale(c.cr, C.double(sx), C.double(sy))
}

func (c *Canvas) Translate(tx, ty float64) {
	C.cairo_translate(c.cr, C.double(tx), C.double(ty))
}

func (c *Canvas) Rotate(angle float64) {
	C.cairo_rotate(c.cr, C.double(angle))
}

func (c *Canvas) Transform(matrix *Matrix) {
	C.cairo_transform(c.cr, matrix.matrixNative)
}

func (c *Canvas) SetMatrix(matrix *Matrix) {
	C.cairo_set_matrix(c.cr, matrix.matrixNative)
}

func (c *Canvas) GetMatrix(matrix *Matrix) {
	C.cairo_get_matrix(c.cr, matrix.matrixNative)
}

func (c *Canvas) IdentityMatrix() {
	C.cairo_identity_matrix(c.cr)
}

// Font

func (c *Canvas) SelectFontFace(family string, fontSlant FontSlant, fontWeight FontWeight) {

	cstrFamily := newCString(family)
	defer freeCString(cstrFamily)

	C.cairo_select_font_face(c.cr, cstrFamily,
		C.cairo_font_slant_t(fontSlant),
		C.cairo_font_weight_t(fontWeight))
}

func (c *Canvas) SetFontSize(size float64) {
	C.cairo_set_font_size(c.cr, C.double(size))
}

// Text

func (c *Canvas) ShowText(text string) {

	cstr := newCString(text)
	defer freeCString(cstr)

	C.cairo_show_text(c.cr, cstr)
}

func (c *Canvas) TextPath(text string) {

	cstr := newCString(text)
	defer freeCString(cstr)

	C.cairo_text_path(c.cr, cstr)
}

type TextExtents struct {
	BearingX float64
	BearingY float64
	Width    float64
	Height   float64
	AdvanceX float64
	AdvanceY float64
}

func (c *Canvas) TextExtents(text string, textExtents *TextExtents) {

	if textExtents == nil {
		return
	}

	cstr := newCString(text)
	defer freeCString(cstr)

	var extents C.cairo_text_extents_t

	C.cairo_text_extents(c.cr, cstr, &extents)

	textExtents.BearingX = float64(extents.x_bearing)
	textExtents.BearingY = float64(extents.y_bearing)
	textExtents.Width = float64(extents.width)
	textExtents.Height = float64(extents.height)
	textExtents.AdvanceX = float64(extents.x_advance)
	textExtents.AdvanceY = float64(extents.y_advance)
}
