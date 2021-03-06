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

//func NewVisualShaderNodeBooleanConstantFromPointer(ptr gdnative.Pointer) VisualShaderNodeBooleanConstant {
func newVisualShaderNodeBooleanConstantFromPointer(ptr gdnative.Pointer) VisualShaderNodeBooleanConstant {
	owner := gdnative.NewObjectFromPointer(ptr)
	obj := VisualShaderNodeBooleanConstant{}
	obj.SetBaseObject(owner)

	return obj
}

/*
Has only one output port and no inputs. Translated to [code]bool[/code] in the shader language.
*/
type VisualShaderNodeBooleanConstant struct {
	VisualShaderNode
	owner gdnative.Object
}

func (o *VisualShaderNodeBooleanConstant) BaseClass() string {
	return "VisualShaderNodeBooleanConstant"
}

/*
        Undocumented
	Args: [], Returns: bool
*/
func (o *VisualShaderNodeBooleanConstant) GetConstant() gdnative.Bool {
	//log.Println("Calling VisualShaderNodeBooleanConstant.GetConstant()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("VisualShaderNodeBooleanConstant", "get_constant")

	// Call the parent method.
	// bool
	retPtr := gdnative.NewEmptyBool()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewBoolFromPointer(retPtr)
	return ret
}

/*
        Undocumented
	Args: [{ false value bool}], Returns: void
*/
func (o *VisualShaderNodeBooleanConstant) SetConstant(value gdnative.Bool) {
	//log.Println("Calling VisualShaderNodeBooleanConstant.SetConstant()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromBool(value)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("VisualShaderNodeBooleanConstant", "set_constant")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

// VisualShaderNodeBooleanConstantImplementer is an interface that implements the methods
// of the VisualShaderNodeBooleanConstant class.
type VisualShaderNodeBooleanConstantImplementer interface {
	VisualShaderNodeImplementer
	GetConstant() gdnative.Bool
	SetConstant(value gdnative.Bool)
}
