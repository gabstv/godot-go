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

//func NewAudioEffectLowShelfFilterFromPointer(ptr gdnative.Pointer) AudioEffectLowShelfFilter {
func newAudioEffectLowShelfFilterFromPointer(ptr gdnative.Pointer) AudioEffectLowShelfFilter {
	owner := gdnative.NewObjectFromPointer(ptr)
	obj := AudioEffectLowShelfFilter{}
	obj.SetBaseObject(owner)

	return obj
}

/*

 */
type AudioEffectLowShelfFilter struct {
	AudioEffectFilter
	owner gdnative.Object
}

func (o *AudioEffectLowShelfFilter) BaseClass() string {
	return "AudioEffectLowShelfFilter"
}

// AudioEffectLowShelfFilterImplementer is an interface that implements the methods
// of the AudioEffectLowShelfFilter class.
type AudioEffectLowShelfFilterImplementer interface {
	AudioEffectFilterImplementer
}
