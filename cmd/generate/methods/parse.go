// Package methods is a package that parses the GDNative headers for type definitions
// of methods
package methods

import (
	"github.com/gabstv/godot-go/cmd/generate/gdnative"
)

// Parse will parse the GDNative headers. Takes a list of headers/structs to ignore.
// Definitions in the given headers and definitions
// with the given name will not be added to the returned list of type definitions.
// We'll need to manually create these structures.
func Parse() gdnative.APIs {

	packagePath := "."

	// Parse the GDNative JSON for method data.
	apis := gdnative.Parse(packagePath)
	return apis
}
