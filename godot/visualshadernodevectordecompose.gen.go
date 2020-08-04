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

//func NewVisualShaderNodeVectorDecomposeFromPointer(ptr gdnative.Pointer) VisualShaderNodeVectorDecompose {
func newVisualShaderNodeVectorDecomposeFromPointer(ptr gdnative.Pointer) VisualShaderNodeVectorDecompose {
	owner := gdnative.NewObjectFromPointer(ptr)
	obj := VisualShaderNodeVectorDecompose{}
	obj.SetBaseObject(owner)

	return obj
}

/*
Takes a [code]vec3[/code] and decomposes it into three scalar values that can be used as separate inputs.
*/
type VisualShaderNodeVectorDecompose struct {
	VisualShaderNode
	owner gdnative.Object
}

func (o *VisualShaderNodeVectorDecompose) BaseClass() string {
	return "VisualShaderNodeVectorDecompose"
}

// VisualShaderNodeVectorDecomposeImplementer is an interface that implements the methods
// of the VisualShaderNodeVectorDecompose class.
type VisualShaderNodeVectorDecomposeImplementer interface {
	VisualShaderNodeImplementer
}
