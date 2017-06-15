// Package graphics provides a graphics engine acting as an (OpenGL) abstraction layer between the application and the graphical backend
package graphics

import (
	"errors"

	"golang.org/x/mobile/gl"
)

// ViewportUpdater is used to let objects express their interests in changes in the viewport dimensions
type ViewportUpdater interface {
	UpdateViewport(width, height int)
}

// Engine is the graphical engine that sits between the game engine and the graphical backend
type Engine struct {
	glctx gl.Context
	vu    []ViewportUpdater
	vW    int
	vH    int
}

// NewEngine creates a new graphics engine
//
// Engine.UpdateViewport() can be used to set initial viewport (it makes sure that all ViewportUpdaters are notified as well)
func NewEngine(context interface{}) (*Engine, error) {
	glctx, ok := context.(gl.Context)
	if !ok {
		return nil, errors.New("Only golang.org/x/mobile/gl Context is supported")
	}
	return &Engine{
		glctx: glctx,
		vu:    make([]ViewportUpdater, 0),
	}, nil
}

// StartRender starts the rendering phase (clears window/screen)
func (e *Engine) StartRender() {
	e.glctx.Viewport(0, 0, e.vW, e.vH)
	e.glctx.ClearColor(0, 0, 0, 1)
	e.glctx.Clear(gl.COLOR_BUFFER_BIT)
}

// RegisterViewportUpdater registers a ViewportUpdater to the list of objects that are called by UpdateViewport()
func (e *Engine) RegisterViewportUpdater(vu ViewportUpdater) {
	e.vu = append(e.vu, vu)
}

// DeregisterViewportUpdater deregisters the ViewportUpdater from the list
func (e *Engine) DeregisterViewportUpdater(vu ViewportUpdater) {
	for i, vu1 := range e.vu {
		if vu == vu1 {
			e.vu[i] = e.vu[len(e.vu)-1]
			e.vu = e.vu[:len(e.vu)-1]
			return
		}
	}
	// Not found, ignore
}

// UpdateViewport notifies all registered ViewportUpdater objects about the (new) viewport dimensions
func (e *Engine) UpdateViewport(width, height int) {
	e.vW = width
	e.vH = height
	for _, vu := range e.vu {
		vu.UpdateViewport(width, height)
	}
}
