package godot

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

//func NewWorld2DFromPointer(ptr gdnative.Pointer) World2D {
func newWorld2DFromPointer(ptr gdnative.Pointer) World2D {
	owner := gdnative.NewObjectFromPointer(ptr)
	obj := World2D{}
	obj.SetBaseObject(owner)

	return obj
}

/*
Class that has everything pertaining to a 2D world. A physics space, a visual scenario and a sound space. 2D nodes register their resources into the current 2D world.
*/
type World2D struct {
	Resource
	owner gdnative.Object
}

func (o *World2D) BaseClass() string {
	return "World2D"
}

/*
        Undocumented
	Args: [], Returns: RID
*/
func (o *World2D) GetCanvas() gdnative.Rid {
	//log.Println("Calling World2D.GetCanvas()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("World2D", "get_canvas")

	// Call the parent method.
	// RID
	retPtr := gdnative.NewEmptyRid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewRidFromPointer(retPtr)
	return ret
}

/*
        Undocumented
	Args: [], Returns: Physics2DDirectSpaceState
*/
func (o *World2D) GetDirectSpaceState() Physics2DDirectSpaceStateImplementer {
	//log.Println("Calling World2D.GetDirectSpaceState()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("World2D", "get_direct_space_state")

	// Call the parent method.
	// Physics2DDirectSpaceState
	retPtr := gdnative.NewEmptyObject()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := newPhysics2DDirectSpaceStateFromPointer(retPtr)

	// Check to see if we already have an instance of this object in our Go instance registry.
	if instance, ok := InstanceRegistry.Get(ret.GetBaseObject().ID()); ok {
		return instance.(Physics2DDirectSpaceStateImplementer)
	}

	// Check to see what kind of class this is and create it. This is generally used with
	// GetNode().
	className := ret.GetClass()
	if className != "Physics2DDirectSpaceState" {
		actualRet := getActualClass(className, ret.GetBaseObject())
		return actualRet.(Physics2DDirectSpaceStateImplementer)
	}

	return &ret
}

/*
        Undocumented
	Args: [], Returns: RID
*/
func (o *World2D) GetSpace() gdnative.Rid {
	//log.Println("Calling World2D.GetSpace()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("World2D", "get_space")

	// Call the parent method.
	// RID
	retPtr := gdnative.NewEmptyRid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewRidFromPointer(retPtr)
	return ret
}

// World2DImplementer is an interface that implements the methods
// of the World2D class.
type World2DImplementer interface {
	ResourceImplementer
	GetCanvas() gdnative.Rid
	GetDirectSpaceState() Physics2DDirectSpaceStateImplementer
	GetSpace() gdnative.Rid
}