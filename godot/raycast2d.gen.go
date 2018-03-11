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

//func NewRayCast2DFromPointer(ptr gdnative.Pointer) RayCast2D {
func newRayCast2DFromPointer(ptr gdnative.Pointer) RayCast2D {
	owner := gdnative.NewObjectFromPointer(ptr)
	obj := RayCast2D{}
	obj.SetBaseObject(owner)

	return obj
}

/*
A RayCast represents a line from its origin to its destination position, [code]cast_to[/code]. It is used to query the 2D space in order to find the closest object along the path of the ray. RayCast2D can ignore some objects by adding them to the exception list via [code]add_exception[/code], by setting proper filtering with collision layers, or by filtering object types with type masks. Only enabled raycasts will be able to query the space and report collisions. RayCast2D calculates intersection every physics frame (see [Node]), and the result is cached so it can be used later until the next frame. If multiple queries are required between physics frames (or during the same frame) use [method force_raycast_update] after adjusting the raycast.
*/
type RayCast2D struct {
	Node2D
	owner gdnative.Object
}

func (o *RayCast2D) BaseClass() string {
	return "RayCast2D"
}

/*
        Adds a collision exception so the ray does not report collisions with the specified node.
	Args: [{ false node Object}], Returns: void
*/
func (o *RayCast2D) AddException(node ObjectImplementer) {
	//log.Println("Calling RayCast2D.AddException()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromObject(node.GetBaseObject())

	// Get the method bind
	methodBind := gdnative.NewMethodBind("RayCast2D", "add_exception")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Adds a collision exception so the ray does not report collisions with the specified [RID].
	Args: [{ false rid RID}], Returns: void
*/
func (o *RayCast2D) AddExceptionRid(rid gdnative.Rid) {
	//log.Println("Calling RayCast2D.AddExceptionRid()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromRid(rid)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("RayCast2D", "add_exception_rid")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Removes all collision exceptions for this ray.
	Args: [], Returns: void
*/
func (o *RayCast2D) ClearExceptions() {
	//log.Println("Calling RayCast2D.ClearExceptions()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("RayCast2D", "clear_exceptions")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Updates the collision information for the ray. Use this method to update the collision information immediately instead of waiting for the next [code]_physics_process[/code] call, for example if the ray or its parent has changed state. Note: [code]enabled == true[/code] is not required for this to work.
	Args: [], Returns: void
*/
func (o *RayCast2D) ForceRaycastUpdate() {
	//log.Println("Calling RayCast2D.ForceRaycastUpdate()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("RayCast2D", "force_raycast_update")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [], Returns: Vector2
*/
func (o *RayCast2D) GetCastTo() gdnative.Vector2 {
	//log.Println("Calling RayCast2D.GetCastTo()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("RayCast2D", "get_cast_to")

	// Call the parent method.
	// Vector2
	retPtr := gdnative.NewEmptyVector2()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewVector2FromPointer(retPtr)
	return ret
}

/*
        Returns the closest object the ray is pointing to. Note that this does not consider the length of the ray, so you must also use [method is_colliding] to check if the object returned is actually colliding with the ray. Example: [codeblock] if RayCast2D.is_colliding(): var collider = RayCast2D.get_collider() [/codeblock]
	Args: [], Returns: Object
*/
func (o *RayCast2D) GetCollider() ObjectImplementer {
	//log.Println("Calling RayCast2D.GetCollider()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("RayCast2D", "get_collider")

	// Call the parent method.
	// Object
	retPtr := gdnative.NewEmptyObject()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := newObjectFromPointer(retPtr)

	// Check to see if we already have an instance of this object in our Go instance registry.
	if instance, ok := InstanceRegistry.Get(ret.GetBaseObject().ID()); ok {
		return instance.(ObjectImplementer)
	}

	// Check to see what kind of class this is and create it. This is generally used with
	// GetNode().
	className := ret.GetClass()
	if className != "Object" {
		actualRet := getActualClass(className, ret.GetBaseObject())
		return actualRet.(ObjectImplementer)
	}

	return &ret
}

/*
        Returns the collision shape of the closest object the ray is pointing to. Note that this does not consider the length of the ray, so you must also use [method is_colliding] to check if the object returned is actually colliding with the ray. Example: [codeblock] if RayCast2D.is_colliding(): var shape = RayCast2D.get_collider_shape() [/codeblock]
	Args: [], Returns: int
*/
func (o *RayCast2D) GetColliderShape() gdnative.Int {
	//log.Println("Calling RayCast2D.GetColliderShape()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("RayCast2D", "get_collider_shape")

	// Call the parent method.
	// int
	retPtr := gdnative.NewEmptyInt()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewIntFromPointer(retPtr)
	return ret
}

/*
        Undocumented
	Args: [], Returns: int
*/
func (o *RayCast2D) GetCollisionMask() gdnative.Int {
	//log.Println("Calling RayCast2D.GetCollisionMask()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("RayCast2D", "get_collision_mask")

	// Call the parent method.
	// int
	retPtr := gdnative.NewEmptyInt()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewIntFromPointer(retPtr)
	return ret
}

/*
        Return an individual bit on the collision mask.
	Args: [{ false bit int}], Returns: bool
*/
func (o *RayCast2D) GetCollisionMaskBit(bit gdnative.Int) gdnative.Bool {
	//log.Println("Calling RayCast2D.GetCollisionMaskBit()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(bit)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("RayCast2D", "get_collision_mask_bit")

	// Call the parent method.
	// bool
	retPtr := gdnative.NewEmptyBool()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewBoolFromPointer(retPtr)
	return ret
}

/*
        Returns the normal of the intersecting object's shape at the collision point.
	Args: [], Returns: Vector2
*/
func (o *RayCast2D) GetCollisionNormal() gdnative.Vector2 {
	//log.Println("Calling RayCast2D.GetCollisionNormal()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("RayCast2D", "get_collision_normal")

	// Call the parent method.
	// Vector2
	retPtr := gdnative.NewEmptyVector2()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewVector2FromPointer(retPtr)
	return ret
}

/*
        Returns the collision point at which the ray intersects the closest object. Note: this point is in the [b]global[/b] coordinate system.
	Args: [], Returns: Vector2
*/
func (o *RayCast2D) GetCollisionPoint() gdnative.Vector2 {
	//log.Println("Calling RayCast2D.GetCollisionPoint()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("RayCast2D", "get_collision_point")

	// Call the parent method.
	// Vector2
	retPtr := gdnative.NewEmptyVector2()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewVector2FromPointer(retPtr)
	return ret
}

/*
        Undocumented
	Args: [], Returns: bool
*/
func (o *RayCast2D) GetExcludeParentBody() gdnative.Bool {
	//log.Println("Calling RayCast2D.GetExcludeParentBody()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("RayCast2D", "get_exclude_parent_body")

	// Call the parent method.
	// bool
	retPtr := gdnative.NewEmptyBool()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewBoolFromPointer(retPtr)
	return ret
}

/*
        Return whether the closest object the ray is pointing to is colliding with the vector (considering the vector length).
	Args: [], Returns: bool
*/
func (o *RayCast2D) IsColliding() gdnative.Bool {
	//log.Println("Calling RayCast2D.IsColliding()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("RayCast2D", "is_colliding")

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
	Args: [], Returns: bool
*/
func (o *RayCast2D) IsEnabled() gdnative.Bool {
	//log.Println("Calling RayCast2D.IsEnabled()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("RayCast2D", "is_enabled")

	// Call the parent method.
	// bool
	retPtr := gdnative.NewEmptyBool()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewBoolFromPointer(retPtr)
	return ret
}

/*
        Removes a collision exception so the ray does report collisions with the specified node.
	Args: [{ false node Object}], Returns: void
*/
func (o *RayCast2D) RemoveException(node ObjectImplementer) {
	//log.Println("Calling RayCast2D.RemoveException()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromObject(node.GetBaseObject())

	// Get the method bind
	methodBind := gdnative.NewMethodBind("RayCast2D", "remove_exception")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Removes a collision exception so the ray does report collisions with the specified [RID].
	Args: [{ false rid RID}], Returns: void
*/
func (o *RayCast2D) RemoveExceptionRid(rid gdnative.Rid) {
	//log.Println("Calling RayCast2D.RemoveExceptionRid()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromRid(rid)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("RayCast2D", "remove_exception_rid")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false local_point Vector2}], Returns: void
*/
func (o *RayCast2D) SetCastTo(localPoint gdnative.Vector2) {
	//log.Println("Calling RayCast2D.SetCastTo()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromVector2(localPoint)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("RayCast2D", "set_cast_to")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false mask int}], Returns: void
*/
func (o *RayCast2D) SetCollisionMask(mask gdnative.Int) {
	//log.Println("Calling RayCast2D.SetCollisionMask()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(mask)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("RayCast2D", "set_collision_mask")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Set/clear individual bits on the collision mask. This makes selecting the areas scanned easier.
	Args: [{ false bit int} { false value bool}], Returns: void
*/
func (o *RayCast2D) SetCollisionMaskBit(bit gdnative.Int, value gdnative.Bool) {
	//log.Println("Calling RayCast2D.SetCollisionMaskBit()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 2, 2)
	ptrArguments[0] = gdnative.NewPointerFromInt(bit)
	ptrArguments[1] = gdnative.NewPointerFromBool(value)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("RayCast2D", "set_collision_mask_bit")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false enabled bool}], Returns: void
*/
func (o *RayCast2D) SetEnabled(enabled gdnative.Bool) {
	//log.Println("Calling RayCast2D.SetEnabled()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromBool(enabled)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("RayCast2D", "set_enabled")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false mask bool}], Returns: void
*/
func (o *RayCast2D) SetExcludeParentBody(mask gdnative.Bool) {
	//log.Println("Calling RayCast2D.SetExcludeParentBody()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromBool(mask)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("RayCast2D", "set_exclude_parent_body")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

// RayCast2DImplementer is an interface that implements the methods
// of the RayCast2D class.
type RayCast2DImplementer interface {
	Node2DImplementer
	AddException(node ObjectImplementer)
	AddExceptionRid(rid gdnative.Rid)
	ClearExceptions()
	ForceRaycastUpdate()
	GetCastTo() gdnative.Vector2
	GetCollider() ObjectImplementer
	GetColliderShape() gdnative.Int
	GetCollisionMask() gdnative.Int
	GetCollisionMaskBit(bit gdnative.Int) gdnative.Bool
	GetCollisionNormal() gdnative.Vector2
	GetCollisionPoint() gdnative.Vector2
	GetExcludeParentBody() gdnative.Bool
	IsColliding() gdnative.Bool
	IsEnabled() gdnative.Bool
	RemoveException(node ObjectImplementer)
	RemoveExceptionRid(rid gdnative.Rid)
	SetCastTo(localPoint gdnative.Vector2)
	SetCollisionMask(mask gdnative.Int)
	SetCollisionMaskBit(bit gdnative.Int, value gdnative.Bool)
	SetEnabled(enabled gdnative.Bool)
	SetExcludeParentBody(mask gdnative.Bool)
}