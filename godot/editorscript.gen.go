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

//func NewEditorScriptFromPointer(ptr gdnative.Pointer) EditorScript {
func newEditorScriptFromPointer(ptr gdnative.Pointer) EditorScript {
	owner := gdnative.NewObjectFromPointer(ptr)
	obj := EditorScript{}
	obj.SetBaseObject(owner)

	return obj
}

/*
Scripts extending this class and implementing its [method _run] method can be executed from the Script Editor's [b]File > Run[/b] menu option (or by pressing [kbd]Ctrl + Shift + X[/kbd]) while the editor is running. This is useful for adding custom in-editor functionality to Godot. For more complex additions, consider using [EditorPlugin]s instead. [b]Note:[/b] Extending scripts need to have [code]tool[/code] mode enabled. [b]Example script:[/b] [codeblock] tool extends EditorScript func _run(): print("Hello from the Godot Editor!") [/codeblock] [b]Note:[/b] The script is run in the Editor context, which means the output is visible in the console window started with the Editor (stdout) instead of the usual Godot [b]Output[/b] dock.
*/
type EditorScript struct {
	Reference
	owner gdnative.Object
}

func (o *EditorScript) BaseClass() string {
	return "EditorScript"
}

/*
        This method is executed by the Editor when [b]File > Run[/b] is used.
	Args: [], Returns: void
*/
func (o *EditorScript) X_Run() {
	//log.Println("Calling EditorScript.X_Run()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("EditorScript", "_run")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Adds [code]node[/code] as a child of the root node in the editor context. [b]Warning:[/b] The implementation of this method is currently disabled.
	Args: [{ false node Node}], Returns: void
*/
func (o *EditorScript) AddRootNode(node NodeImplementer) {
	//log.Println("Calling EditorScript.AddRootNode()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromObject(node.GetBaseObject())

	// Get the method bind
	methodBind := gdnative.NewMethodBind("EditorScript", "add_root_node")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Returns the [EditorInterface] singleton instance.
	Args: [], Returns: EditorInterface
*/
func (o *EditorScript) GetEditorInterface() EditorInterfaceImplementer {
	//log.Println("Calling EditorScript.GetEditorInterface()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("EditorScript", "get_editor_interface")

	// Call the parent method.
	// EditorInterface
	retPtr := gdnative.NewEmptyObject()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := newEditorInterfaceFromPointer(retPtr)

	// Check to see if we already have an instance of this object in our Go instance registry.
	if instance, ok := InstanceRegistry.Get(ret.GetBaseObject().ID()); ok {
		return instance.(EditorInterfaceImplementer)
	}

	// Check to see what kind of class this is and create it. This is generally used with
	// GetNode().
	className := ret.GetClass()
	if className != "EditorInterface" {
		actualRet := getActualClass(className, ret.GetBaseObject())
		return actualRet.(EditorInterfaceImplementer)
	}

	return &ret
}

/*
        Returns the Editor's currently active scene.
	Args: [], Returns: Node
*/
func (o *EditorScript) GetScene() NodeImplementer {
	//log.Println("Calling EditorScript.GetScene()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("EditorScript", "get_scene")

	// Call the parent method.
	// Node
	retPtr := gdnative.NewEmptyObject()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := newNodeFromPointer(retPtr)

	// Check to see if we already have an instance of this object in our Go instance registry.
	if instance, ok := InstanceRegistry.Get(ret.GetBaseObject().ID()); ok {
		return instance.(NodeImplementer)
	}

	// Check to see what kind of class this is and create it. This is generally used with
	// GetNode().
	className := ret.GetClass()
	if className != "Node" {
		actualRet := getActualClass(className, ret.GetBaseObject())
		return actualRet.(NodeImplementer)
	}

	return &ret
}

// EditorScriptImplementer is an interface that implements the methods
// of the EditorScript class.
type EditorScriptImplementer interface {
	ReferenceImplementer
	X_Run()
	AddRootNode(node NodeImplementer)
	GetEditorInterface() EditorInterfaceImplementer
	GetScene() NodeImplementer
}
