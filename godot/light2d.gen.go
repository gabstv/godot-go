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

// Light2DMode is an enum for Mode values.
type Light2DMode int

const (
	Light2DModeAdd  Light2DMode = 0
	Light2DModeMask Light2DMode = 3
	Light2DModeMix  Light2DMode = 2
	Light2DModeSub  Light2DMode = 1
)

// Light2DShadowFilter is an enum for ShadowFilter values.
type Light2DShadowFilter int

const (
	Light2DShadowFilterNone  Light2DShadowFilter = 0
	Light2DShadowFilterPcf13 Light2DShadowFilter = 5
	Light2DShadowFilterPcf3  Light2DShadowFilter = 1
	Light2DShadowFilterPcf5  Light2DShadowFilter = 2
	Light2DShadowFilterPcf7  Light2DShadowFilter = 3
	Light2DShadowFilterPcf9  Light2DShadowFilter = 4
)

//func NewLight2DFromPointer(ptr gdnative.Pointer) Light2D {
func newLight2DFromPointer(ptr gdnative.Pointer) Light2D {
	owner := gdnative.NewObjectFromPointer(ptr)
	obj := Light2D{}
	obj.SetBaseObject(owner)

	return obj
}

/*
Casts light in a 2D environment. Light is defined by a (usually grayscale) texture, a color, an energy value, a mode (see constants), and various other parameters (range and shadows-related). [b]Note:[/b] Light2D can also be used as a mask.
*/
type Light2D struct {
	Node2D
	owner gdnative.Object
}

func (o *Light2D) BaseClass() string {
	return "Light2D"
}

/*
        Undocumented
	Args: [], Returns: Color
*/
func (o *Light2D) GetColor() gdnative.Color {
	//log.Println("Calling Light2D.GetColor()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Light2D", "get_color")

	// Call the parent method.
	// Color
	retPtr := gdnative.NewEmptyColor()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewColorFromPointer(retPtr)
	return ret
}

/*
        Undocumented
	Args: [], Returns: float
*/
func (o *Light2D) GetEnergy() gdnative.Real {
	//log.Println("Calling Light2D.GetEnergy()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Light2D", "get_energy")

	// Call the parent method.
	// float
	retPtr := gdnative.NewEmptyReal()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewRealFromPointer(retPtr)
	return ret
}

/*
        Undocumented
	Args: [], Returns: float
*/
func (o *Light2D) GetHeight() gdnative.Real {
	//log.Println("Calling Light2D.GetHeight()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Light2D", "get_height")

	// Call the parent method.
	// float
	retPtr := gdnative.NewEmptyReal()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewRealFromPointer(retPtr)
	return ret
}

/*
        Undocumented
	Args: [], Returns: int
*/
func (o *Light2D) GetItemCullMask() gdnative.Int {
	//log.Println("Calling Light2D.GetItemCullMask()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Light2D", "get_item_cull_mask")

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
func (o *Light2D) GetItemShadowCullMask() gdnative.Int {
	//log.Println("Calling Light2D.GetItemShadowCullMask()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Light2D", "get_item_shadow_cull_mask")

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
func (o *Light2D) GetLayerRangeMax() gdnative.Int {
	//log.Println("Calling Light2D.GetLayerRangeMax()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Light2D", "get_layer_range_max")

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
func (o *Light2D) GetLayerRangeMin() gdnative.Int {
	//log.Println("Calling Light2D.GetLayerRangeMin()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Light2D", "get_layer_range_min")

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
	Args: [], Returns: enum.Light2D::Mode
*/
func (o *Light2D) GetMode() Light2DMode {
	//log.Println("Calling Light2D.GetMode()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Light2D", "get_mode")

	// Call the parent method.
	// enum.Light2D::Mode
	retPtr := gdnative.NewEmptyInt()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewIntFromPointer(retPtr)
	return Light2DMode(ret)
}

/*
        Undocumented
	Args: [], Returns: int
*/
func (o *Light2D) GetShadowBufferSize() gdnative.Int {
	//log.Println("Calling Light2D.GetShadowBufferSize()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Light2D", "get_shadow_buffer_size")

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
	Args: [], Returns: Color
*/
func (o *Light2D) GetShadowColor() gdnative.Color {
	//log.Println("Calling Light2D.GetShadowColor()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Light2D", "get_shadow_color")

	// Call the parent method.
	// Color
	retPtr := gdnative.NewEmptyColor()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewColorFromPointer(retPtr)
	return ret
}

/*
        Undocumented
	Args: [], Returns: enum.Light2D::ShadowFilter
*/
func (o *Light2D) GetShadowFilter() Light2DShadowFilter {
	//log.Println("Calling Light2D.GetShadowFilter()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Light2D", "get_shadow_filter")

	// Call the parent method.
	// enum.Light2D::ShadowFilter
	retPtr := gdnative.NewEmptyInt()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewIntFromPointer(retPtr)
	return Light2DShadowFilter(ret)
}

/*
        Undocumented
	Args: [], Returns: float
*/
func (o *Light2D) GetShadowGradientLength() gdnative.Real {
	//log.Println("Calling Light2D.GetShadowGradientLength()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Light2D", "get_shadow_gradient_length")

	// Call the parent method.
	// float
	retPtr := gdnative.NewEmptyReal()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewRealFromPointer(retPtr)
	return ret
}

/*
        Undocumented
	Args: [], Returns: float
*/
func (o *Light2D) GetShadowSmooth() gdnative.Real {
	//log.Println("Calling Light2D.GetShadowSmooth()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Light2D", "get_shadow_smooth")

	// Call the parent method.
	// float
	retPtr := gdnative.NewEmptyReal()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewRealFromPointer(retPtr)
	return ret
}

/*
        Undocumented
	Args: [], Returns: Texture
*/
func (o *Light2D) GetTexture() TextureImplementer {
	//log.Println("Calling Light2D.GetTexture()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Light2D", "get_texture")

	// Call the parent method.
	// Texture
	retPtr := gdnative.NewEmptyObject()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := newTextureFromPointer(retPtr)

	// Check to see if we already have an instance of this object in our Go instance registry.
	if instance, ok := InstanceRegistry.Get(ret.GetBaseObject().ID()); ok {
		return instance.(TextureImplementer)
	}

	// Check to see what kind of class this is and create it. This is generally used with
	// GetNode().
	className := ret.GetClass()
	if className != "Texture" {
		actualRet := getActualClass(className, ret.GetBaseObject())
		return actualRet.(TextureImplementer)
	}

	return &ret
}

/*
        Undocumented
	Args: [], Returns: Vector2
*/
func (o *Light2D) GetTextureOffset() gdnative.Vector2 {
	//log.Println("Calling Light2D.GetTextureOffset()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Light2D", "get_texture_offset")

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
	Args: [], Returns: float
*/
func (o *Light2D) GetTextureScale() gdnative.Real {
	//log.Println("Calling Light2D.GetTextureScale()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Light2D", "get_texture_scale")

	// Call the parent method.
	// float
	retPtr := gdnative.NewEmptyReal()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewRealFromPointer(retPtr)
	return ret
}

/*
        Undocumented
	Args: [], Returns: int
*/
func (o *Light2D) GetZRangeMax() gdnative.Int {
	//log.Println("Calling Light2D.GetZRangeMax()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Light2D", "get_z_range_max")

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
func (o *Light2D) GetZRangeMin() gdnative.Int {
	//log.Println("Calling Light2D.GetZRangeMin()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Light2D", "get_z_range_min")

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
	Args: [], Returns: bool
*/
func (o *Light2D) IsEditorOnly() gdnative.Bool {
	//log.Println("Calling Light2D.IsEditorOnly()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Light2D", "is_editor_only")

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
func (o *Light2D) IsEnabled() gdnative.Bool {
	//log.Println("Calling Light2D.IsEnabled()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Light2D", "is_enabled")

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
func (o *Light2D) IsShadowEnabled() gdnative.Bool {
	//log.Println("Calling Light2D.IsShadowEnabled()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Light2D", "is_shadow_enabled")

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
	Args: [{ false color Color}], Returns: void
*/
func (o *Light2D) SetColor(color gdnative.Color) {
	//log.Println("Calling Light2D.SetColor()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromColor(color)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Light2D", "set_color")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false editor_only bool}], Returns: void
*/
func (o *Light2D) SetEditorOnly(editorOnly gdnative.Bool) {
	//log.Println("Calling Light2D.SetEditorOnly()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromBool(editorOnly)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Light2D", "set_editor_only")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false enabled bool}], Returns: void
*/
func (o *Light2D) SetEnabled(enabled gdnative.Bool) {
	//log.Println("Calling Light2D.SetEnabled()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromBool(enabled)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Light2D", "set_enabled")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false energy float}], Returns: void
*/
func (o *Light2D) SetEnergy(energy gdnative.Real) {
	//log.Println("Calling Light2D.SetEnergy()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromReal(energy)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Light2D", "set_energy")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false height float}], Returns: void
*/
func (o *Light2D) SetHeight(height gdnative.Real) {
	//log.Println("Calling Light2D.SetHeight()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromReal(height)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Light2D", "set_height")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false item_cull_mask int}], Returns: void
*/
func (o *Light2D) SetItemCullMask(itemCullMask gdnative.Int) {
	//log.Println("Calling Light2D.SetItemCullMask()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(itemCullMask)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Light2D", "set_item_cull_mask")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false item_shadow_cull_mask int}], Returns: void
*/
func (o *Light2D) SetItemShadowCullMask(itemShadowCullMask gdnative.Int) {
	//log.Println("Calling Light2D.SetItemShadowCullMask()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(itemShadowCullMask)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Light2D", "set_item_shadow_cull_mask")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false layer int}], Returns: void
*/
func (o *Light2D) SetLayerRangeMax(layer gdnative.Int) {
	//log.Println("Calling Light2D.SetLayerRangeMax()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(layer)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Light2D", "set_layer_range_max")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false layer int}], Returns: void
*/
func (o *Light2D) SetLayerRangeMin(layer gdnative.Int) {
	//log.Println("Calling Light2D.SetLayerRangeMin()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(layer)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Light2D", "set_layer_range_min")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false mode int}], Returns: void
*/
func (o *Light2D) SetMode(mode gdnative.Int) {
	//log.Println("Calling Light2D.SetMode()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(mode)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Light2D", "set_mode")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false size int}], Returns: void
*/
func (o *Light2D) SetShadowBufferSize(size gdnative.Int) {
	//log.Println("Calling Light2D.SetShadowBufferSize()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(size)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Light2D", "set_shadow_buffer_size")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false shadow_color Color}], Returns: void
*/
func (o *Light2D) SetShadowColor(shadowColor gdnative.Color) {
	//log.Println("Calling Light2D.SetShadowColor()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromColor(shadowColor)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Light2D", "set_shadow_color")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false enabled bool}], Returns: void
*/
func (o *Light2D) SetShadowEnabled(enabled gdnative.Bool) {
	//log.Println("Calling Light2D.SetShadowEnabled()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromBool(enabled)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Light2D", "set_shadow_enabled")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false filter int}], Returns: void
*/
func (o *Light2D) SetShadowFilter(filter gdnative.Int) {
	//log.Println("Calling Light2D.SetShadowFilter()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(filter)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Light2D", "set_shadow_filter")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false multiplier float}], Returns: void
*/
func (o *Light2D) SetShadowGradientLength(multiplier gdnative.Real) {
	//log.Println("Calling Light2D.SetShadowGradientLength()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromReal(multiplier)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Light2D", "set_shadow_gradient_length")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false smooth float}], Returns: void
*/
func (o *Light2D) SetShadowSmooth(smooth gdnative.Real) {
	//log.Println("Calling Light2D.SetShadowSmooth()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromReal(smooth)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Light2D", "set_shadow_smooth")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false texture Texture}], Returns: void
*/
func (o *Light2D) SetTexture(texture TextureImplementer) {
	//log.Println("Calling Light2D.SetTexture()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromObject(texture.GetBaseObject())

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Light2D", "set_texture")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false texture_offset Vector2}], Returns: void
*/
func (o *Light2D) SetTextureOffset(textureOffset gdnative.Vector2) {
	//log.Println("Calling Light2D.SetTextureOffset()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromVector2(textureOffset)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Light2D", "set_texture_offset")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false texture_scale float}], Returns: void
*/
func (o *Light2D) SetTextureScale(textureScale gdnative.Real) {
	//log.Println("Calling Light2D.SetTextureScale()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromReal(textureScale)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Light2D", "set_texture_scale")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false z int}], Returns: void
*/
func (o *Light2D) SetZRangeMax(z gdnative.Int) {
	//log.Println("Calling Light2D.SetZRangeMax()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(z)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Light2D", "set_z_range_max")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false z int}], Returns: void
*/
func (o *Light2D) SetZRangeMin(z gdnative.Int) {
	//log.Println("Calling Light2D.SetZRangeMin()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(z)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Light2D", "set_z_range_min")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

// Light2DImplementer is an interface that implements the methods
// of the Light2D class.
type Light2DImplementer interface {
	Node2DImplementer
	GetColor() gdnative.Color
	GetEnergy() gdnative.Real
	GetHeight() gdnative.Real
	GetItemCullMask() gdnative.Int
	GetItemShadowCullMask() gdnative.Int
	GetLayerRangeMax() gdnative.Int
	GetLayerRangeMin() gdnative.Int
	GetShadowBufferSize() gdnative.Int
	GetShadowColor() gdnative.Color
	GetShadowGradientLength() gdnative.Real
	GetShadowSmooth() gdnative.Real
	GetTexture() TextureImplementer
	GetTextureOffset() gdnative.Vector2
	GetTextureScale() gdnative.Real
	GetZRangeMax() gdnative.Int
	GetZRangeMin() gdnative.Int
	IsEditorOnly() gdnative.Bool
	IsEnabled() gdnative.Bool
	IsShadowEnabled() gdnative.Bool
	SetColor(color gdnative.Color)
	SetEditorOnly(editorOnly gdnative.Bool)
	SetEnabled(enabled gdnative.Bool)
	SetEnergy(energy gdnative.Real)
	SetHeight(height gdnative.Real)
	SetItemCullMask(itemCullMask gdnative.Int)
	SetItemShadowCullMask(itemShadowCullMask gdnative.Int)
	SetLayerRangeMax(layer gdnative.Int)
	SetLayerRangeMin(layer gdnative.Int)
	SetMode(mode gdnative.Int)
	SetShadowBufferSize(size gdnative.Int)
	SetShadowColor(shadowColor gdnative.Color)
	SetShadowEnabled(enabled gdnative.Bool)
	SetShadowFilter(filter gdnative.Int)
	SetShadowGradientLength(multiplier gdnative.Real)
	SetShadowSmooth(smooth gdnative.Real)
	SetTexture(texture TextureImplementer)
	SetTextureOffset(textureOffset gdnative.Vector2)
	SetTextureScale(textureScale gdnative.Real)
	SetZRangeMax(z gdnative.Int)
	SetZRangeMin(z gdnative.Int)
}
