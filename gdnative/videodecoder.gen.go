package gdnative

/*------------------------------------------------------------------------------
//   This code was generated by a tool.
//
//   Changes to this file may cause incorrect behavior and will be lost if
//   the code is regenerated. Any updates should be done in
//   "types.go.tmpl" so they can be included in the generated
//   code.
//----------------------------------------------------------------------------*/

/*
#cgo CFLAGS: -I../cgodeps/godot-cpp/godot_headers
#include "gdnative.gen.h"
// #include <cgodeps/godot-cpp/godot_headers/videodecoder/godot_videodecoder.h>
// Include all headers for now. TODO: Look up all the required
// headers we need to import based on the method arguments and return types.
#include <gdnative/aabb.h>
#include <gdnative/array.h>
#include <gdnative/basis.h>
#include <gdnative/color.h>
#include <gdnative/dictionary.h>
#include <gdnative/gdnative.h>
#include <gdnative/node_path.h>
#include <gdnative/plane.h>
#include <gdnative/pool_arrays.h>
#include <gdnative/quat.h>
#include <gdnative/rect2.h>
#include <gdnative/rid.h>
#include <gdnative/string.h>
#include <gdnative/string_name.h>
#include <gdnative/transform.h>
#include <gdnative/transform2d.h>
#include <gdnative/variant.h>
#include <gdnative/vector2.h>
#include <gdnative/vector3.h>
#include <gdnative_api_struct.gen.h>
*/
import "C"
import "unsafe"

// NewEmptyVideodecoderInterfaceGdnative will return a pointer to an empty
// initialized VideodecoderInterfaceGdnative. This is primarily used in
// conjunction with MethodBindPtrCall.
func NewEmptyVideodecoderInterfaceGdnative() Pointer {
	var obj C.godot_videodecoder_interface_gdnative
	return Pointer{base: unsafe.Pointer(&obj)}
}

// NewPointerFromVideodecoderInterfaceGdnative will return an unsafe pointer to the given
// object. This is primarily used in conjunction with MethodBindPtrCall.
func NewPointerFromVideodecoderInterfaceGdnative(obj VideodecoderInterfaceGdnative) Pointer {
	return Pointer{base: unsafe.Pointer(obj.getBase())}
}

// NewVideodecoderInterfaceGdnativeFromPointer will return a VideodecoderInterfaceGdnative from the
// given unsafe pointer. This is primarily used in conjunction with MethodBindPtrCall.
func NewVideodecoderInterfaceGdnativeFromPointer(ptr Pointer) VideodecoderInterfaceGdnative {

	return VideodecoderInterfaceGdnative{base: (*C.godot_videodecoder_interface_gdnative)(ptr.getBase())}
}

type VideodecoderInterfaceGdnative struct {
	base *C.godot_videodecoder_interface_gdnative

	Version GdnativeApiVersion
}

func (gdt VideodecoderInterfaceGdnative) getBase() *C.godot_videodecoder_interface_gdnative {
	return gdt.base
}
