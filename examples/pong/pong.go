package main

import (
	"github.com/shadowapex/go-godot/godot"
	"log"
)

// NewPongClass is a constructor that we can pass to godot.
func NewPongClass() godot.Class {
	pongClass := &PongClass{}

	return pongClass
}

// PongClass is a simple Godot class that can be registered.
type PongClass struct {
	godot.Node2D
}

// Xready is called as soon as the node enters the scene.
func (p *PongClass) X_ready() {
	godot.Log.Println("Pong is ready!")
}

func (p *PongClass) X_process(delta float64) {
	godot.Log.Println("Processing in pong.go!")
	godot.Log.Println("  Delta:", delta)
}

func (p *PongClass) CustomFunc(myPhrase string) {
	godot.Log.Println(myPhrase)
}

// The "init()" function is a special Go function that will be called when this library
// is initialized. Here we can register our Godot classes.
func init() {
	// SetGodotGDNativeInit will set the given function to run on library initialization.
	godot.SetGodotGDNativeInit(func(options *godot.GodotGDNativeInitOptions) {
		log.Println("This is being called from pong.go!")
	})

	// Register will register the given class constructor with Godot.
	godot.Register(NewPongClass)
}

// This never gets called, but it necessary to export as a shared library.
func main() {
}