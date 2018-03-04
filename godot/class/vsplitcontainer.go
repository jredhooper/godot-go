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

//func NewVSplitContainerFromPointer(ptr gdnative.Pointer) VSplitContainer {
func NewVSplitContainerFromPointer(ptr gdnative.Pointer) VSplitContainer {
	owner := gdnative.NewObjectFromPointer(ptr)
	obj := VSplitContainer{}
	obj.SetBaseObject(owner)

	return obj
}

/*
Vertical split container. See [SplitContainer]. This goes from left to right.
*/
type VSplitContainer struct {
	SplitContainer
	owner gdnative.Object
}

func (o *VSplitContainer) BaseClass() string {
	return "VSplitContainer"
}

// SetBaseObject will internally set the Godot object inside the struct.
// This is used to call parent methods.
func (o *VSplitContainer) SetBaseObject(object gdnative.Object) {
	o.owner = object
}

func (o *VSplitContainer) GetBaseObject() gdnative.Object {
	return o.owner
}