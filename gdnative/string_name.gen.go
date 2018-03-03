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
#include <gdnative/string_name.h>
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

// NewPointerFromStringName will return an unsafe pointer to the given
// object. This is primarily used in conjunction with MethodBindPtrCall.
func NewPointerFromStringName(obj StringName) Pointer {
	return Pointer{base: unsafe.Pointer(obj.getBase())}
}

// NewStringNameFromPointer will return a StringName from the
// given unsafe pointer. This is primarily used in conjunction with MethodBindPtrCall.
func NewStringNameFromPointer(ptr Pointer) StringName {
	return StringName{base: (*C.godot_string_name)(ptr.getBase())}
}

type StringName struct {
	base *C.godot_string_name
}

func (gdt StringName) getBase() *C.godot_string_name {
	return gdt.base
}

// GetName godot_string_name_get_name [[const godot_string_name * p_self]] godot_string
func (gdt *StringName) GetName() String {
	arg0 := gdt.getBase()

	ret := C.go_godot_string_name_get_name(GDNative.api, arg0)

	return String{base: &ret}

}

// GetHash godot_string_name_get_hash [[const godot_string_name * p_self]] uint32_t
func (gdt *StringName) GetHash() Uint32T {
	arg0 := gdt.getBase()

	ret := C.go_godot_string_name_get_hash(GDNative.api, arg0)

	return Uint32T(ret)
}

// GetDataUniquePointer godot_string_name_get_data_unique_pointer [[const godot_string_name * p_self]] const void *
func (gdt *StringName) GetDataUniquePointer() {
	arg0 := gdt.getBase()

	C.go_godot_string_name_get_data_unique_pointer(GDNative.api, arg0)
}

// OperatorEqual godot_string_name_operator_equal [[const godot_string_name * p_self] [const godot_string_name * p_other]] godot_bool
func (gdt *StringName) OperatorEqual(other StringName) Bool {
	arg0 := gdt.getBase()
	arg1 := other.getBase()

	ret := C.go_godot_string_name_operator_equal(GDNative.api, arg0, arg1)

	return Bool(ret)
}

// OperatorLess godot_string_name_operator_less [[const godot_string_name * p_self] [const godot_string_name * p_other]] godot_bool
func (gdt *StringName) OperatorLess(other StringName) Bool {
	arg0 := gdt.getBase()
	arg1 := other.getBase()

	ret := C.go_godot_string_name_operator_less(GDNative.api, arg0, arg1)

	return Bool(ret)
}

// Destroy godot_string_name_destroy [[godot_string_name * p_self]] void
func (gdt *StringName) Destroy() {
	arg0 := gdt.getBase()

	C.go_godot_string_name_destroy(GDNative.api, arg0)
}
