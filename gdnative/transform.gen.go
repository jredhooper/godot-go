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
#include "gdnative.gen.h"
#include <gdnative/transform.h>
// Include all headers for now. TODO: Look up all the required
// headers we need to import based on the method arguments and return types.
#include <gdnative/aabb.h>
#include <gdnative/array.h>
#include <gdnative/basis.h>
#include <gdnative/color.h>
#include <gdnative/dictionary.h>
#include <gdnative/gdnative.h>
#include <gdnative/node_path.h>
#include <gdnative/plane.h>
#include <gdnative/pool_arrays.h>
#include <gdnative/quat.h>
#include <gdnative/rect2.h>
#include <gdnative/rid.h>
#include <gdnative/string.h>
#include <gdnative/string_name.h>
#include <gdnative/transform.h>
#include <gdnative/transform2d.h>
#include <gdnative/variant.h>
#include <gdnative/vector2.h>
#include <gdnative/vector3.h>
#include <gdnative_api_struct.gen.h>
*/
import "C"
import "unsafe"

// NewPointerFromTransform will return an unsafe pointer to the given
// object. This is primarily used in conjunction with MethodBindPtrCall.
func NewPointerFromTransform(obj Transform) Pointer {
	return Pointer{base: unsafe.Pointer(obj.getBase())}
}

// NewTransformFromPointer will return a Transform from the
// given unsafe pointer. This is primarily used in conjunction with MethodBindPtrCall.
func NewTransformFromPointer(ptr Pointer) Transform {
	return Transform{base: (*C.godot_transform)(ptr.getBase())}
}

type Transform struct {
	base *C.godot_transform
}

func (gdt Transform) getBase() *C.godot_transform {
	return gdt.base
}

// NewTransformWithAxisOrigin godot_transform_new_with_axis_origin [[godot_transform * r_dest] [const godot_vector3 * p_x_axis] [const godot_vector3 * p_y_axis] [const godot_vector3 * p_z_axis] [const godot_vector3 * p_origin]] void
func NewTransformWithAxisOrigin(xAxis Vector3, yAxis Vector3, zAxis Vector3, origin Vector3) *Transform {
	var dest C.godot_transform
	arg1 := xAxis.getBase()
	arg2 := yAxis.getBase()
	arg3 := zAxis.getBase()
	arg4 := origin.getBase()
	C.go_godot_transform_new_with_axis_origin(GDNative.api, &dest, arg1, arg2, arg3, arg4)
	return &Transform{base: &dest}
}

// NewTransform godot_transform_new [[godot_transform * r_dest] [const godot_basis * p_basis] [const godot_vector3 * p_origin]] void
func NewTransform(basis Basis, origin Vector3) *Transform {
	var dest C.godot_transform
	arg1 := basis.getBase()
	arg2 := origin.getBase()
	C.go_godot_transform_new(GDNative.api, &dest, arg1, arg2)
	return &Transform{base: &dest}
}

// GetBasis godot_transform_get_basis [[const godot_transform * p_self]] godot_basis
func (gdt *Transform) GetBasis() Basis {
	arg0 := gdt.getBase()

	ret := C.go_godot_transform_get_basis(GDNative.api, arg0)

	return Basis{base: &ret}

}

// SetBasis godot_transform_set_basis [[godot_transform * p_self] [const godot_basis * p_v]] void
func (gdt *Transform) SetBasis(v Basis) {
	arg0 := gdt.getBase()
	arg1 := v.getBase()

	C.go_godot_transform_set_basis(GDNative.api, arg0, arg1)
}

// GetOrigin godot_transform_get_origin [[const godot_transform * p_self]] godot_vector3
func (gdt *Transform) GetOrigin() Vector3 {
	arg0 := gdt.getBase()

	ret := C.go_godot_transform_get_origin(GDNative.api, arg0)

	return Vector3{base: &ret}

}

// SetOrigin godot_transform_set_origin [[godot_transform * p_self] [const godot_vector3 * p_v]] void
func (gdt *Transform) SetOrigin(v Vector3) {
	arg0 := gdt.getBase()
	arg1 := v.getBase()

	C.go_godot_transform_set_origin(GDNative.api, arg0, arg1)
}

// AsString godot_transform_as_string [[const godot_transform * p_self]] godot_string
func (gdt *Transform) AsString() String {
	arg0 := gdt.getBase()

	ret := C.go_godot_transform_as_string(GDNative.api, arg0)

	return String{base: &ret}

}

// Inverse godot_transform_inverse [[const godot_transform * p_self]] godot_transform
func (gdt *Transform) Inverse() Transform {
	arg0 := gdt.getBase()

	ret := C.go_godot_transform_inverse(GDNative.api, arg0)

	return Transform{base: &ret}

}

// AffineInverse godot_transform_affine_inverse [[const godot_transform * p_self]] godot_transform
func (gdt *Transform) AffineInverse() Transform {
	arg0 := gdt.getBase()

	ret := C.go_godot_transform_affine_inverse(GDNative.api, arg0)

	return Transform{base: &ret}

}

// Orthonormalized godot_transform_orthonormalized [[const godot_transform * p_self]] godot_transform
func (gdt *Transform) Orthonormalized() Transform {
	arg0 := gdt.getBase()

	ret := C.go_godot_transform_orthonormalized(GDNative.api, arg0)

	return Transform{base: &ret}

}

// Rotated godot_transform_rotated [[const godot_transform * p_self] [const godot_vector3 * p_axis] [const godot_real p_phi]] godot_transform
func (gdt *Transform) Rotated(axis Vector3, phi Real) Transform {
	arg0 := gdt.getBase()
	arg1 := axis.getBase()
	arg2 := phi.getBase()

	ret := C.go_godot_transform_rotated(GDNative.api, arg0, arg1, arg2)

	return Transform{base: &ret}

}

// Scaled godot_transform_scaled [[const godot_transform * p_self] [const godot_vector3 * p_scale]] godot_transform
func (gdt *Transform) Scaled(scale Vector3) Transform {
	arg0 := gdt.getBase()
	arg1 := scale.getBase()

	ret := C.go_godot_transform_scaled(GDNative.api, arg0, arg1)

	return Transform{base: &ret}

}

// Translated godot_transform_translated [[const godot_transform * p_self] [const godot_vector3 * p_ofs]] godot_transform
func (gdt *Transform) Translated(ofs Vector3) Transform {
	arg0 := gdt.getBase()
	arg1 := ofs.getBase()

	ret := C.go_godot_transform_translated(GDNative.api, arg0, arg1)

	return Transform{base: &ret}

}

// LookingAt godot_transform_looking_at [[const godot_transform * p_self] [const godot_vector3 * p_target] [const godot_vector3 * p_up]] godot_transform
func (gdt *Transform) LookingAt(target Vector3, up Vector3) Transform {
	arg0 := gdt.getBase()
	arg1 := target.getBase()
	arg2 := up.getBase()

	ret := C.go_godot_transform_looking_at(GDNative.api, arg0, arg1, arg2)

	return Transform{base: &ret}

}

// XformPlane godot_transform_xform_plane [[const godot_transform * p_self] [const godot_plane * p_v]] godot_plane
func (gdt *Transform) XformPlane(v Plane) Plane {
	arg0 := gdt.getBase()
	arg1 := v.getBase()

	ret := C.go_godot_transform_xform_plane(GDNative.api, arg0, arg1)

	return Plane{base: &ret}

}

// XformInvPlane godot_transform_xform_inv_plane [[const godot_transform * p_self] [const godot_plane * p_v]] godot_plane
func (gdt *Transform) XformInvPlane(v Plane) Plane {
	arg0 := gdt.getBase()
	arg1 := v.getBase()

	ret := C.go_godot_transform_xform_inv_plane(GDNative.api, arg0, arg1)

	return Plane{base: &ret}

}

// NewTransformIdentity godot_transform_new_identity [[godot_transform * r_dest]] void
func NewTransformIdentity() *Transform {
	var dest C.godot_transform
	C.go_godot_transform_new_identity(GDNative.api, &dest)
	return &Transform{base: &dest}
}

// OperatorEqual godot_transform_operator_equal [[const godot_transform * p_self] [const godot_transform * p_b]] godot_bool
func (gdt *Transform) OperatorEqual(b Transform) Bool {
	arg0 := gdt.getBase()
	arg1 := b.getBase()

	ret := C.go_godot_transform_operator_equal(GDNative.api, arg0, arg1)

	return Bool(ret)
}

// OperatorMultiply godot_transform_operator_multiply [[const godot_transform * p_self] [const godot_transform * p_b]] godot_transform
func (gdt *Transform) OperatorMultiply(b Transform) Transform {
	arg0 := gdt.getBase()
	arg1 := b.getBase()

	ret := C.go_godot_transform_operator_multiply(GDNative.api, arg0, arg1)

	return Transform{base: &ret}

}

// XformVector3 godot_transform_xform_vector3 [[const godot_transform * p_self] [const godot_vector3 * p_v]] godot_vector3
func (gdt *Transform) XformVector3(v Vector3) Vector3 {
	arg0 := gdt.getBase()
	arg1 := v.getBase()

	ret := C.go_godot_transform_xform_vector3(GDNative.api, arg0, arg1)

	return Vector3{base: &ret}

}

// XformInvVector3 godot_transform_xform_inv_vector3 [[const godot_transform * p_self] [const godot_vector3 * p_v]] godot_vector3
func (gdt *Transform) XformInvVector3(v Vector3) Vector3 {
	arg0 := gdt.getBase()
	arg1 := v.getBase()

	ret := C.go_godot_transform_xform_inv_vector3(GDNative.api, arg0, arg1)

	return Vector3{base: &ret}

}

// XformAabb godot_transform_xform_aabb [[const godot_transform * p_self] [const godot_aabb * p_v]] godot_aabb
func (gdt *Transform) XformAabb(v Aabb) Aabb {
	arg0 := gdt.getBase()
	arg1 := v.getBase()

	ret := C.go_godot_transform_xform_aabb(GDNative.api, arg0, arg1)

	return Aabb{base: &ret}

}

// XformInvAabb godot_transform_xform_inv_aabb [[const godot_transform * p_self] [const godot_aabb * p_v]] godot_aabb
func (gdt *Transform) XformInvAabb(v Aabb) Aabb {
	arg0 := gdt.getBase()
	arg1 := v.getBase()

	ret := C.go_godot_transform_xform_inv_aabb(GDNative.api, arg0, arg1)

	return Aabb{base: &ret}

}
