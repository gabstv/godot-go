package godot

import (
	"github.com/gabstv/godot-go/gdnative"
)

// Log is a logging interface that will let you push log messages to Godot. This
// will enable them to show up within the Godot Editor.
var Log = &gdnative.Logger{StackNum: 3}
