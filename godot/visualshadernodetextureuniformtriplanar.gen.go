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

//func NewVisualShaderNodeTextureUniformTriplanarFromPointer(ptr gdnative.Pointer) VisualShaderNodeTextureUniformTriplanar {
func newVisualShaderNodeTextureUniformTriplanarFromPointer(ptr gdnative.Pointer) VisualShaderNodeTextureUniformTriplanar {
	owner := gdnative.NewObjectFromPointer(ptr)
	obj := VisualShaderNodeTextureUniformTriplanar{}
	obj.SetBaseObject(owner)

	return obj
}

/*
Performs a lookup operation on the texture provided as a uniform for the shader, with support for triplanar mapping.
*/
type VisualShaderNodeTextureUniformTriplanar struct {
	VisualShaderNodeTextureUniform
	owner gdnative.Object
}

func (o *VisualShaderNodeTextureUniformTriplanar) BaseClass() string {
	return "VisualShaderNodeTextureUniformTriplanar"
}

// VisualShaderNodeTextureUniformTriplanarImplementer is an interface that implements the methods
// of the VisualShaderNodeTextureUniformTriplanar class.
type VisualShaderNodeTextureUniformTriplanarImplementer interface {
	VisualShaderNodeTextureUniformImplementer
}
