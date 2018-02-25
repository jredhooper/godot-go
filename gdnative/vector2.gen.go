package gdnative

/*------------------------------------------------------------------------------
//   This code was generated by a tool.
//
//   Changes to this file may cause incorrect behavior and will be lost if
//   the code is regenerated. Any updates should be done in
//   "types.go.tmpl" so they can be included in the generated
//   code.
//----------------------------------------------------------------------------*/

/*
#include <gdnative/vector2.h>
*/
import "C"

type Vector2 struct {
	base *C.godot_vector2
}

func (t Vector2) getBase() *C.godot_vector2 {
	return t.base
}

// AsString godot_vector2_as_string [[const godot_vector2 * p_self]] godot_string
func (t *Vector2) AsString() String {
	arg0 := t.getBase()

	ret := C.go_godot_vector2_as_string(GDNative.api, arg0)

	return String{base: ret}

}

// Normalized godot_vector2_normalized [[const godot_vector2 * p_self]] godot_vector2
func (t *Vector2) Normalized() Vector2 {
	arg0 := t.getBase()

	ret := C.go_godot_vector2_normalized(GDNative.api, arg0)

	return Vector2{base: ret}

}

// Length godot_vector2_length [[const godot_vector2 * p_self]] godot_real
func (t *Vector2) Length() Real {
	arg0 := t.getBase()

	ret := C.go_godot_vector2_length(GDNative.api, arg0)

	return Real{base: ret}

}

// Angle godot_vector2_angle [[const godot_vector2 * p_self]] godot_real
func (t *Vector2) Angle() Real {
	arg0 := t.getBase()

	ret := C.go_godot_vector2_angle(GDNative.api, arg0)

	return Real{base: ret}

}

// LengthSquared godot_vector2_length_squared [[const godot_vector2 * p_self]] godot_real
func (t *Vector2) LengthSquared() Real {
	arg0 := t.getBase()

	ret := C.go_godot_vector2_length_squared(GDNative.api, arg0)

	return Real{base: ret}

}

// IsNormalized godot_vector2_is_normalized [[const godot_vector2 * p_self]] godot_bool
func (t *Vector2) IsNormalized() Bool {
	arg0 := t.getBase()

	ret := C.go_godot_vector2_is_normalized(GDNative.api, arg0)

	return Bool{base: ret}

}

// DistanceTo godot_vector2_distance_to [[const godot_vector2 * p_self] [const godot_vector2 * p_to]] godot_real
func (t *Vector2) DistanceTo(to Vector2) Real {
	arg0 := t.getBase()
	arg1 := to.getBase()

	ret := C.go_godot_vector2_distance_to(GDNative.api, arg0, arg1)

	return Real{base: ret}

}

// DistanceSquaredTo godot_vector2_distance_squared_to [[const godot_vector2 * p_self] [const godot_vector2 * p_to]] godot_real
func (t *Vector2) DistanceSquaredTo(to Vector2) Real {
	arg0 := t.getBase()
	arg1 := to.getBase()

	ret := C.go_godot_vector2_distance_squared_to(GDNative.api, arg0, arg1)

	return Real{base: ret}

}

// AngleTo godot_vector2_angle_to [[const godot_vector2 * p_self] [const godot_vector2 * p_to]] godot_real
func (t *Vector2) AngleTo(to Vector2) Real {
	arg0 := t.getBase()
	arg1 := to.getBase()

	ret := C.go_godot_vector2_angle_to(GDNative.api, arg0, arg1)

	return Real{base: ret}

}

// AngleToPoint godot_vector2_angle_to_point [[const godot_vector2 * p_self] [const godot_vector2 * p_to]] godot_real
func (t *Vector2) AngleToPoint(to Vector2) Real {
	arg0 := t.getBase()
	arg1 := to.getBase()

	ret := C.go_godot_vector2_angle_to_point(GDNative.api, arg0, arg1)

	return Real{base: ret}

}

// LinearInterpolate godot_vector2_linear_interpolate [[const godot_vector2 * p_self] [const godot_vector2 * p_b] [const godot_real p_t]] godot_vector2
func (t *Vector2) LinearInterpolate(b Vector2, t Real) Vector2 {
	arg0 := t.getBase()
	arg1 := b.getBase()
	arg2 := t.getBase()

	ret := C.go_godot_vector2_linear_interpolate(GDNative.api, arg0, arg1, arg2)

	return Vector2{base: ret}

}

// CubicInterpolate godot_vector2_cubic_interpolate [[const godot_vector2 * p_self] [const godot_vector2 * p_b] [const godot_vector2 * p_pre_a] [const godot_vector2 * p_post_b] [const godot_real p_t]] godot_vector2
func (t *Vector2) CubicInterpolate(b Vector2, preA Vector2, postB Vector2, t Real) Vector2 {
	arg0 := t.getBase()
	arg1 := b.getBase()
	arg2 := preA.getBase()
	arg3 := postB.getBase()
	arg4 := t.getBase()

	ret := C.go_godot_vector2_cubic_interpolate(GDNative.api, arg0, arg1, arg2, arg3, arg4)

	return Vector2{base: ret}

}

// Rotated godot_vector2_rotated [[const godot_vector2 * p_self] [const godot_real p_phi]] godot_vector2
func (t *Vector2) Rotated(phi Real) Vector2 {
	arg0 := t.getBase()
	arg1 := phi.getBase()

	ret := C.go_godot_vector2_rotated(GDNative.api, arg0, arg1)

	return Vector2{base: ret}

}

// Tangent godot_vector2_tangent [[const godot_vector2 * p_self]] godot_vector2
func (t *Vector2) Tangent() Vector2 {
	arg0 := t.getBase()

	ret := C.go_godot_vector2_tangent(GDNative.api, arg0)

	return Vector2{base: ret}

}

// Floor godot_vector2_floor [[const godot_vector2 * p_self]] godot_vector2
func (t *Vector2) Floor() Vector2 {
	arg0 := t.getBase()

	ret := C.go_godot_vector2_floor(GDNative.api, arg0)

	return Vector2{base: ret}

}

// Snapped godot_vector2_snapped [[const godot_vector2 * p_self] [const godot_vector2 * p_by]] godot_vector2
func (t *Vector2) Snapped(by Vector2) Vector2 {
	arg0 := t.getBase()
	arg1 := by.getBase()

	ret := C.go_godot_vector2_snapped(GDNative.api, arg0, arg1)

	return Vector2{base: ret}

}

// Aspect godot_vector2_aspect [[const godot_vector2 * p_self]] godot_real
func (t *Vector2) Aspect() Real {
	arg0 := t.getBase()

	ret := C.go_godot_vector2_aspect(GDNative.api, arg0)

	return Real{base: ret}

}

// Dot godot_vector2_dot [[const godot_vector2 * p_self] [const godot_vector2 * p_with]] godot_real
func (t *Vector2) Dot(with Vector2) Real {
	arg0 := t.getBase()
	arg1 := with.getBase()

	ret := C.go_godot_vector2_dot(GDNative.api, arg0, arg1)

	return Real{base: ret}

}

// Slide godot_vector2_slide [[const godot_vector2 * p_self] [const godot_vector2 * p_n]] godot_vector2
func (t *Vector2) Slide(n Vector2) Vector2 {
	arg0 := t.getBase()
	arg1 := n.getBase()

	ret := C.go_godot_vector2_slide(GDNative.api, arg0, arg1)

	return Vector2{base: ret}

}

// Bounce godot_vector2_bounce [[const godot_vector2 * p_self] [const godot_vector2 * p_n]] godot_vector2
func (t *Vector2) Bounce(n Vector2) Vector2 {
	arg0 := t.getBase()
	arg1 := n.getBase()

	ret := C.go_godot_vector2_bounce(GDNative.api, arg0, arg1)

	return Vector2{base: ret}

}

// Reflect godot_vector2_reflect [[const godot_vector2 * p_self] [const godot_vector2 * p_n]] godot_vector2
func (t *Vector2) Reflect(n Vector2) Vector2 {
	arg0 := t.getBase()
	arg1 := n.getBase()

	ret := C.go_godot_vector2_reflect(GDNative.api, arg0, arg1)

	return Vector2{base: ret}

}

// Abs godot_vector2_abs [[const godot_vector2 * p_self]] godot_vector2
func (t *Vector2) Abs() Vector2 {
	arg0 := t.getBase()

	ret := C.go_godot_vector2_abs(GDNative.api, arg0)

	return Vector2{base: ret}

}

// Clamped godot_vector2_clamped [[const godot_vector2 * p_self] [const godot_real p_length]] godot_vector2
func (t *Vector2) Clamped(length Real) Vector2 {
	arg0 := t.getBase()
	arg1 := length.getBase()

	ret := C.go_godot_vector2_clamped(GDNative.api, arg0, arg1)

	return Vector2{base: ret}

}

// OperatorAdd godot_vector2_operator_add [[const godot_vector2 * p_self] [const godot_vector2 * p_b]] godot_vector2
func (t *Vector2) OperatorAdd(b Vector2) Vector2 {
	arg0 := t.getBase()
	arg1 := b.getBase()

	ret := C.go_godot_vector2_operator_add(GDNative.api, arg0, arg1)

	return Vector2{base: ret}

}

// OperatorSubtract godot_vector2_operator_subtract [[const godot_vector2 * p_self] [const godot_vector2 * p_b]] godot_vector2
func (t *Vector2) OperatorSubtract(b Vector2) Vector2 {
	arg0 := t.getBase()
	arg1 := b.getBase()

	ret := C.go_godot_vector2_operator_subtract(GDNative.api, arg0, arg1)

	return Vector2{base: ret}

}

// OperatorMultiplyVector godot_vector2_operator_multiply_vector [[const godot_vector2 * p_self] [const godot_vector2 * p_b]] godot_vector2
func (t *Vector2) OperatorMultiplyVector(b Vector2) Vector2 {
	arg0 := t.getBase()
	arg1 := b.getBase()

	ret := C.go_godot_vector2_operator_multiply_vector(GDNative.api, arg0, arg1)

	return Vector2{base: ret}

}

// OperatorMultiplyScalar godot_vector2_operator_multiply_scalar [[const godot_vector2 * p_self] [const godot_real p_b]] godot_vector2
func (t *Vector2) OperatorMultiplyScalar(b Real) Vector2 {
	arg0 := t.getBase()
	arg1 := b.getBase()

	ret := C.go_godot_vector2_operator_multiply_scalar(GDNative.api, arg0, arg1)

	return Vector2{base: ret}

}

// OperatorDivideVector godot_vector2_operator_divide_vector [[const godot_vector2 * p_self] [const godot_vector2 * p_b]] godot_vector2
func (t *Vector2) OperatorDivideVector(b Vector2) Vector2 {
	arg0 := t.getBase()
	arg1 := b.getBase()

	ret := C.go_godot_vector2_operator_divide_vector(GDNative.api, arg0, arg1)

	return Vector2{base: ret}

}

// OperatorDivideScalar godot_vector2_operator_divide_scalar [[const godot_vector2 * p_self] [const godot_real p_b]] godot_vector2
func (t *Vector2) OperatorDivideScalar(b Real) Vector2 {
	arg0 := t.getBase()
	arg1 := b.getBase()

	ret := C.go_godot_vector2_operator_divide_scalar(GDNative.api, arg0, arg1)

	return Vector2{base: ret}

}

// OperatorEqual godot_vector2_operator_equal [[const godot_vector2 * p_self] [const godot_vector2 * p_b]] godot_bool
func (t *Vector2) OperatorEqual(b Vector2) Bool {
	arg0 := t.getBase()
	arg1 := b.getBase()

	ret := C.go_godot_vector2_operator_equal(GDNative.api, arg0, arg1)

	return Bool{base: ret}

}

// OperatorLess godot_vector2_operator_less [[const godot_vector2 * p_self] [const godot_vector2 * p_b]] godot_bool
func (t *Vector2) OperatorLess(b Vector2) Bool {
	arg0 := t.getBase()
	arg1 := b.getBase()

	ret := C.go_godot_vector2_operator_less(GDNative.api, arg0, arg1)

	return Bool{base: ret}

}

// OperatorNeg godot_vector2_operator_neg [[const godot_vector2 * p_self]] godot_vector2
func (t *Vector2) OperatorNeg() Vector2 {
	arg0 := t.getBase()

	ret := C.go_godot_vector2_operator_neg(GDNative.api, arg0)

	return Vector2{base: ret}

}

// SetX godot_vector2_set_x [[godot_vector2 * p_self] [const godot_real p_x]] void
func (t *Vector2) SetX(x Real) {
	arg0 := t.getBase()
	arg1 := x.getBase()

	C.go_godot_vector2_set_x(GDNative.api, arg0, arg1)

}

// SetY godot_vector2_set_y [[godot_vector2 * p_self] [const godot_real p_y]] void
func (t *Vector2) SetY(y Real) {
	arg0 := t.getBase()
	arg1 := y.getBase()

	C.go_godot_vector2_set_y(GDNative.api, arg0, arg1)

}

// GetX godot_vector2_get_x [[const godot_vector2 * p_self]] godot_real
func (t *Vector2) GetX() Real {
	arg0 := t.getBase()

	ret := C.go_godot_vector2_get_x(GDNative.api, arg0)

	return Real{base: ret}

}

// GetY godot_vector2_get_y [[const godot_vector2 * p_self]] godot_real
func (t *Vector2) GetY() Real {
	arg0 := t.getBase()

	ret := C.go_godot_vector2_get_y(GDNative.api, arg0)

	return Real{base: ret}

}
