// Package graphics provides a graphics engine acting as an (OpenGL) abstraction layer between the application and the graphical backend
package graphics

import (
	"errors"

	"golang.org/x/mobile/gl"
)

// Engine is the graphical engine that sits between the game engine and the graphical backend
type Engine struct {
	glctx gl.Context
}

// NewEngine creates a new graphics engine
func NewEngine(context interface{}) (*Engine, error) {
	glctx, ok := context.(gl.Context)
	if !ok {
		return nil, errors.New("Only golang.org/x/mobile/gl Context is supported")
	}
	return &Engine{
		glctx: glctx,
	}, nil
}

// StartRender starts the rendering phase (clears window/screen)
func (e *Engine) StartRender() {
	e.glctx.ClearColor(0, 0, 0, 1)
	e.glctx.Clear(gl.COLOR_BUFFER_BIT)
}
