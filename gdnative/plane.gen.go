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
#include <gdnative/plane.h>
*/
import "C"

type Plane struct {
	base *C.godot_plane
}

func (t Plane) getBase() *C.godot_plane {
	return t.base
}

// AsString godot_plane_as_string [[const godot_plane * p_self]] godot_string
func (t *Plane) AsString() String {
	arg0 := t.getBase()

	ret := C.go_godot_plane_as_string(GDNative.api, arg0)

	return String{base: ret}

}

// Normalized godot_plane_normalized [[const godot_plane * p_self]] godot_plane
func (t *Plane) Normalized() Plane {
	arg0 := t.getBase()

	ret := C.go_godot_plane_normalized(GDNative.api, arg0)

	return Plane{base: ret}

}

// Center godot_plane_center [[const godot_plane * p_self]] godot_vector3
func (t *Plane) Center() Vector3 {
	arg0 := t.getBase()

	ret := C.go_godot_plane_center(GDNative.api, arg0)

	return Vector3{base: ret}

}

// GetAnyPoint godot_plane_get_any_point [[const godot_plane * p_self]] godot_vector3
func (t *Plane) GetAnyPoint() Vector3 {
	arg0 := t.getBase()

	ret := C.go_godot_plane_get_any_point(GDNative.api, arg0)

	return Vector3{base: ret}

}

// IsPointOver godot_plane_is_point_over [[const godot_plane * p_self] [const godot_vector3 * p_point]] godot_bool
func (t *Plane) IsPointOver(point Vector3) Bool {
	arg0 := t.getBase()
	arg1 := point.getBase()

	ret := C.go_godot_plane_is_point_over(GDNative.api, arg0, arg1)

	return Bool{base: ret}

}

// DistanceTo godot_plane_distance_to [[const godot_plane * p_self] [const godot_vector3 * p_point]] godot_real
func (t *Plane) DistanceTo(point Vector3) Real {
	arg0 := t.getBase()
	arg1 := point.getBase()

	ret := C.go_godot_plane_distance_to(GDNative.api, arg0, arg1)

	return Real{base: ret}

}

// HasPoint godot_plane_has_point [[const godot_plane * p_self] [const godot_vector3 * p_point] [const godot_real p_epsilon]] godot_bool
func (t *Plane) HasPoint(point Vector3, epsilon Real) Bool {
	arg0 := t.getBase()
	arg1 := point.getBase()
	arg2 := epsilon.getBase()

	ret := C.go_godot_plane_has_point(GDNative.api, arg0, arg1, arg2)

	return Bool{base: ret}

}

// Project godot_plane_project [[const godot_plane * p_self] [const godot_vector3 * p_point]] godot_vector3
func (t *Plane) Project(point Vector3) Vector3 {
	arg0 := t.getBase()
	arg1 := point.getBase()

	ret := C.go_godot_plane_project(GDNative.api, arg0, arg1)

	return Vector3{base: ret}

}

// Intersect3 godot_plane_intersect_3 [[const godot_plane * p_self] [godot_vector3 * r_dest] [const godot_plane * p_b] [const godot_plane * p_c]] godot_bool
func (t *Plane) Intersect3(dest Vector3, b Plane, c Plane) Bool {
	arg0 := t.getBase()
	arg1 := dest.getBase()
	arg2 := b.getBase()
	arg3 := c.getBase()

	ret := C.go_godot_plane_intersect_3(GDNative.api, arg0, arg1, arg2, arg3)

	return Bool{base: ret}

}

// IntersectsRay godot_plane_intersects_ray [[const godot_plane * p_self] [godot_vector3 * r_dest] [const godot_vector3 * p_from] [const godot_vector3 * p_dir]] godot_bool
func (t *Plane) IntersectsRay(dest Vector3, from Vector3, dir Vector3) Bool {
	arg0 := t.getBase()
	arg1 := dest.getBase()
	arg2 := from.getBase()
	arg3 := dir.getBase()

	ret := C.go_godot_plane_intersects_ray(GDNative.api, arg0, arg1, arg2, arg3)

	return Bool{base: ret}

}

// IntersectsSegment godot_plane_intersects_segment [[const godot_plane * p_self] [godot_vector3 * r_dest] [const godot_vector3 * p_begin] [const godot_vector3 * p_end]] godot_bool
func (t *Plane) IntersectsSegment(dest Vector3, begin Vector3, end Vector3) Bool {
	arg0 := t.getBase()
	arg1 := dest.getBase()
	arg2 := begin.getBase()
	arg3 := end.getBase()

	ret := C.go_godot_plane_intersects_segment(GDNative.api, arg0, arg1, arg2, arg3)

	return Bool{base: ret}

}

// OperatorNeg godot_plane_operator_neg [[const godot_plane * p_self]] godot_plane
func (t *Plane) OperatorNeg() Plane {
	arg0 := t.getBase()

	ret := C.go_godot_plane_operator_neg(GDNative.api, arg0)

	return Plane{base: ret}

}

// OperatorEqual godot_plane_operator_equal [[const godot_plane * p_self] [const godot_plane * p_b]] godot_bool
func (t *Plane) OperatorEqual(b Plane) Bool {
	arg0 := t.getBase()
	arg1 := b.getBase()

	ret := C.go_godot_plane_operator_equal(GDNative.api, arg0, arg1)

	return Bool{base: ret}

}

// SetNormal godot_plane_set_normal [[godot_plane * p_self] [const godot_vector3 * p_normal]] void
func (t *Plane) SetNormal(normal Vector3) {
	arg0 := t.getBase()
	arg1 := normal.getBase()

	C.go_godot_plane_set_normal(GDNative.api, arg0, arg1)

}

// GetNormal godot_plane_get_normal [[const godot_plane * p_self]] godot_vector3
func (t *Plane) GetNormal() Vector3 {
	arg0 := t.getBase()

	ret := C.go_godot_plane_get_normal(GDNative.api, arg0)

	return Vector3{base: ret}

}

// GetD godot_plane_get_d [[const godot_plane * p_self]] godot_real
func (t *Plane) GetD() Real {
	arg0 := t.getBase()

	ret := C.go_godot_plane_get_d(GDNative.api, arg0)

	return Real{base: ret}

}

// SetD godot_plane_set_d [[godot_plane * p_self] [const godot_real p_d]] void
func (t *Plane) SetD(d Real) {
	arg0 := t.getBase()
	arg1 := d.getBase()

	C.go_godot_plane_set_d(GDNative.api, arg0, arg1)

}
