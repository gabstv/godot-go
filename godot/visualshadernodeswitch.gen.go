package godot

import (
	"github.com/gabstv/godot-go/gdnative"
)

/*------------------------------------------------------------------------------
//   This code was generated by a tool.
//
//   Changes to this file may cause incorrect behavior and will be lost if
//   the code is regenerated. Any updates should be done in
//   "class.go.tmpl" so they can be included in the generated
//   code.
//----------------------------------------------------------------------------*/

//func NewVisualShaderNodeSwitchFromPointer(ptr gdnative.Pointer) VisualShaderNodeSwitch {
func newVisualShaderNodeSwitchFromPointer(ptr gdnative.Pointer) VisualShaderNodeSwitch {
	owner := gdnative.NewObjectFromPointer(ptr)
	obj := VisualShaderNodeSwitch{}
	obj.SetBaseObject(owner)

	return obj
}

/*
Returns an associated vector if the provided boolean value is [code]true[/code] or [code]false[/code].
*/
type VisualShaderNodeSwitch struct {
	VisualShaderNode
	owner gdnative.Object
}

func (o *VisualShaderNodeSwitch) BaseClass() string {
	return "VisualShaderNodeSwitch"
}

// VisualShaderNodeSwitchImplementer is an interface that implements the methods
// of the VisualShaderNodeSwitch class.
type VisualShaderNodeSwitchImplementer interface {
	VisualShaderNodeImplementer
}
