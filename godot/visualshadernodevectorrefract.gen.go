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

//func NewVisualShaderNodeVectorRefractFromPointer(ptr gdnative.Pointer) VisualShaderNodeVectorRefract {
func newVisualShaderNodeVectorRefractFromPointer(ptr gdnative.Pointer) VisualShaderNodeVectorRefract {
	owner := gdnative.NewObjectFromPointer(ptr)
	obj := VisualShaderNodeVectorRefract{}
	obj.SetBaseObject(owner)

	return obj
}

/*
Translated to [code]refract(I, N, eta)[/code] in the shader language, where [code]I[/code] is the incident vector, [code]N[/code] is the normal vector and [code]eta[/code] is the ratio of the indices of the refraction.
*/
type VisualShaderNodeVectorRefract struct {
	VisualShaderNode
	owner gdnative.Object
}

func (o *VisualShaderNodeVectorRefract) BaseClass() string {
	return "VisualShaderNodeVectorRefract"
}

// VisualShaderNodeVectorRefractImplementer is an interface that implements the methods
// of the VisualShaderNodeVectorRefract class.
type VisualShaderNodeVectorRefractImplementer interface {
	VisualShaderNodeImplementer
}
