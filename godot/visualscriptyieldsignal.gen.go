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

// VisualScriptYieldSignalCallMode is an enum for CallMode values.
type VisualScriptYieldSignalCallMode int

const (
	VisualScriptYieldSignalCallModeInstance VisualScriptYieldSignalCallMode = 2
	VisualScriptYieldSignalCallModeNodePath VisualScriptYieldSignalCallMode = 1
	VisualScriptYieldSignalCallModeSelf     VisualScriptYieldSignalCallMode = 0
)

//func NewVisualScriptYieldSignalFromPointer(ptr gdnative.Pointer) VisualScriptYieldSignal {
func newVisualScriptYieldSignalFromPointer(ptr gdnative.Pointer) VisualScriptYieldSignal {
	owner := gdnative.NewObjectFromPointer(ptr)
	obj := VisualScriptYieldSignal{}
	obj.SetBaseObject(owner)

	return obj
}

/*
Undocumented
*/
type VisualScriptYieldSignal struct {
	VisualScriptNode
	owner gdnative.Object
}

func (o *VisualScriptYieldSignal) BaseClass() string {
	return "VisualScriptYieldSignal"
}

/*
        Undocumented
	Args: [], Returns: NodePath
*/
func (o *VisualScriptYieldSignal) GetBasePath() gdnative.NodePath {
	//log.Println("Calling VisualScriptYieldSignal.GetBasePath()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("VisualScriptYieldSignal", "get_base_path")

	// Call the parent method.
	// NodePath
	retPtr := gdnative.NewEmptyNodePath()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewNodePathFromPointer(retPtr)
	return ret
}

/*
        Undocumented
	Args: [], Returns: String
*/
func (o *VisualScriptYieldSignal) GetBaseType() gdnative.String {
	//log.Println("Calling VisualScriptYieldSignal.GetBaseType()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("VisualScriptYieldSignal", "get_base_type")

	// Call the parent method.
	// String
	retPtr := gdnative.NewEmptyString()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewStringFromPointer(retPtr)
	return ret
}

/*
        Undocumented
	Args: [], Returns: enum.VisualScriptYieldSignal::CallMode
*/
func (o *VisualScriptYieldSignal) GetCallMode() VisualScriptYieldSignalCallMode {
	//log.Println("Calling VisualScriptYieldSignal.GetCallMode()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("VisualScriptYieldSignal", "get_call_mode")

	// Call the parent method.
	// enum.VisualScriptYieldSignal::CallMode
	retPtr := gdnative.NewEmptyInt()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewIntFromPointer(retPtr)
	return VisualScriptYieldSignalCallMode(ret)
}

/*
        Undocumented
	Args: [], Returns: String
*/
func (o *VisualScriptYieldSignal) GetSignal() gdnative.String {
	//log.Println("Calling VisualScriptYieldSignal.GetSignal()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("VisualScriptYieldSignal", "get_signal")

	// Call the parent method.
	// String
	retPtr := gdnative.NewEmptyString()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewStringFromPointer(retPtr)
	return ret
}

/*
        Undocumented
	Args: [{ false base_path NodePath}], Returns: void
*/
func (o *VisualScriptYieldSignal) SetBasePath(basePath gdnative.NodePath) {
	//log.Println("Calling VisualScriptYieldSignal.SetBasePath()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromNodePath(basePath)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("VisualScriptYieldSignal", "set_base_path")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false base_type String}], Returns: void
*/
func (o *VisualScriptYieldSignal) SetBaseType(baseType gdnative.String) {
	//log.Println("Calling VisualScriptYieldSignal.SetBaseType()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromString(baseType)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("VisualScriptYieldSignal", "set_base_type")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false mode int}], Returns: void
*/
func (o *VisualScriptYieldSignal) SetCallMode(mode gdnative.Int) {
	//log.Println("Calling VisualScriptYieldSignal.SetCallMode()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(mode)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("VisualScriptYieldSignal", "set_call_mode")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false signal String}], Returns: void
*/
func (o *VisualScriptYieldSignal) SetSignal(signal gdnative.String) {
	//log.Println("Calling VisualScriptYieldSignal.SetSignal()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromString(signal)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("VisualScriptYieldSignal", "set_signal")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

// VisualScriptYieldSignalImplementer is an interface that implements the methods
// of the VisualScriptYieldSignal class.
type VisualScriptYieldSignalImplementer interface {
	VisualScriptNodeImplementer
	GetBasePath() gdnative.NodePath
	GetBaseType() gdnative.String
	GetSignal() gdnative.String
	SetBasePath(basePath gdnative.NodePath)
	SetBaseType(baseType gdnative.String)
	SetCallMode(mode gdnative.Int)
	SetSignal(signal gdnative.String)
}
