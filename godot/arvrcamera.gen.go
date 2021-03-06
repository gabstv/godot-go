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

//func NewARVRCameraFromPointer(ptr gdnative.Pointer) ARVRCamera {
func newARVRCameraFromPointer(ptr gdnative.Pointer) ARVRCamera {
	owner := gdnative.NewObjectFromPointer(ptr)
	obj := ARVRCamera{}
	obj.SetBaseObject(owner)

	return obj
}

/*
Undocumented
*/
type ARVRCamera struct {
	Camera
	owner gdnative.Object
}

func (o *ARVRCamera) BaseClass() string {
	return "ARVRCamera"
}

// ARVRCameraImplementer is an interface that implements the methods
// of the ARVRCamera class.
type ARVRCameraImplementer interface {
	CameraImplementer
}
