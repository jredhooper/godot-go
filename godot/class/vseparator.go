package class

import (
	"github.com/shadowapex/godot-go/gdnative"
)

/*------------------------------------------------------------------------------
//   This code was generated by a tool.
//
//   Changes to this file may cause incorrect behavior and will be lost if
//   the code is regenerated. Any updates should be done in
//   "class.go.tmpl" so they can be included in the generated
//   code.
//----------------------------------------------------------------------------*/

//func NewVSeparatorFromPointer(ptr gdnative.Pointer) VSeparator {
func NewVSeparatorFromPointer(ptr gdnative.Pointer) VSeparator {
	owner := gdnative.NewObjectFromPointer(ptr)
	obj := VSeparator{}
	obj.SetBaseObject(owner)

	return obj
}

/*
Vertical version of [Separator]. It is used to separate objects horizontally, though (but it looks vertical!).
*/
type VSeparator struct {
	Separator
	owner gdnative.Object
}

func (o *VSeparator) BaseClass() string {
	return "VSeparator"
}

// SetBaseObject will internally set the Godot object inside the struct.
// This is used to call parent methods.
func (o *VSeparator) SetBaseObject(object gdnative.Object) {
	o.owner = object
}

func (o *VSeparator) GetBaseObject() gdnative.Object {
	return o.owner
}