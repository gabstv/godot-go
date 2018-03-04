package class

import (
	"log"

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

//func NewVisualScriptYieldFromPointer(ptr gdnative.Pointer) VisualScriptYield {
func NewVisualScriptYieldFromPointer(ptr gdnative.Pointer) VisualScriptYield {
	owner := gdnative.NewObjectFromPointer(ptr)
	obj := VisualScriptYield{}
	obj.SetBaseObject(owner)

	return obj
}

/*
Undocumented
*/
type VisualScriptYield struct {
	VisualScriptNode
	owner gdnative.Object
}

func (o *VisualScriptYield) BaseClass() string {
	return "VisualScriptYield"
}

// SetBaseObject will internally set the Godot object inside the struct.
// This is used to call parent methods.
func (o *VisualScriptYield) SetBaseObject(object gdnative.Object) {
	o.owner = object
}

func (o *VisualScriptYield) GetBaseObject() gdnative.Object {
	return o.owner
}

/*
        Undocumented
	Args: [], Returns: float
*/
func (o *VisualScriptYield) GetWaitTime() gdnative.Float {
	log.Println("Calling VisualScriptYield.GetWaitTime()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("VisualScriptYield", "get_wait_time")

	// Call the parent method.
	// float
	retPtr := gdnative.NewEmptyFloat()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewFloatFromPointer(retPtr)
	log.Println("  Got return value: ", ret)
	return ret
}

/*
        Undocumented
	Args: [], Returns: enum.VisualScriptYield::YieldMode
*/

/*
        Undocumented
	Args: [{ false sec float}], Returns: void
*/
func (o *VisualScriptYield) SetWaitTime(sec gdnative.Float) {
	log.Println("Calling VisualScriptYield.SetWaitTime()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromFloat(sec)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("VisualScriptYield", "set_wait_time")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false mode int}], Returns: void
*/
func (o *VisualScriptYield) SetYieldMode(mode gdnative.Int) {
	log.Println("Calling VisualScriptYield.SetYieldMode()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(mode)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("VisualScriptYield", "set_yield_mode")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}