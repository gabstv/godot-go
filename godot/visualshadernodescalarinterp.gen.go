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

//func NewVisualShaderNodeScalarInterpFromPointer(ptr gdnative.Pointer) VisualShaderNodeScalarInterp {
func newVisualShaderNodeScalarInterpFromPointer(ptr gdnative.Pointer) VisualShaderNodeScalarInterp {
	owner := gdnative.NewObjectFromPointer(ptr)
	obj := VisualShaderNodeScalarInterp{}
	obj.SetBaseObject(owner)

	return obj
}

/*
Translates to [code]mix(a, b, weight)[/code] in the shader language.
*/
type VisualShaderNodeScalarInterp struct {
	VisualShaderNode
	owner gdnative.Object
}

func (o *VisualShaderNodeScalarInterp) BaseClass() string {
	return "VisualShaderNodeScalarInterp"
}

// VisualShaderNodeScalarInterpImplementer is an interface that implements the methods
// of the VisualShaderNodeScalarInterp class.
type VisualShaderNodeScalarInterpImplementer interface {
	VisualShaderNodeImplementer
}
