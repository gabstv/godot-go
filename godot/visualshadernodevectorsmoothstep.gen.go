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

//func NewVisualShaderNodeVectorSmoothStepFromPointer(ptr gdnative.Pointer) VisualShaderNodeVectorSmoothStep {
func newVisualShaderNodeVectorSmoothStepFromPointer(ptr gdnative.Pointer) VisualShaderNodeVectorSmoothStep {
	owner := gdnative.NewObjectFromPointer(ptr)
	obj := VisualShaderNodeVectorSmoothStep{}
	obj.SetBaseObject(owner)

	return obj
}

/*
Translates to [code]smoothstep(edge0, edge1, x)[/code] in the shader language, where [code]x[/code] is a vector. Returns [code]0.0[/code] if [code]x[/code] is smaller than [code]edge0[/code] and [code]1.0[/code] if [code]x[/code] is larger than [code]edge1[/code]. Otherwise the return value is interpolated between [code]0.0[/code] and [code]1.0[/code] using Hermite polynomials.
*/
type VisualShaderNodeVectorSmoothStep struct {
	VisualShaderNode
	owner gdnative.Object
}

func (o *VisualShaderNodeVectorSmoothStep) BaseClass() string {
	return "VisualShaderNodeVectorSmoothStep"
}

// VisualShaderNodeVectorSmoothStepImplementer is an interface that implements the methods
// of the VisualShaderNodeVectorSmoothStep class.
type VisualShaderNodeVectorSmoothStepImplementer interface {
	VisualShaderNodeImplementer
}
