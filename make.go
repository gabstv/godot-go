package main

// Copy submodule to track dependencies (make sure that they are up to date)

//go:generate rsync -rv --exclude=.git --exclude=.github godot_engine cgodeps
//go:generate rsync -rv --exclude=.git --exclude=.github godot-cpp cgodeps

// Generate the classes.
//go:generate go run cmd/generate/main.go

func main() {}
