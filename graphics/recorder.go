package graphics

import (
	"errors"

	"golang.org/x/mobile/gl"
)

// Recorder is an object that is able to record the graphics state and easily and efficiently switch (back) to such a state
//
// It is able to record:
// * active (Element)Buffer objects and their properties
type Recorder struct {
	engine *Engine
	vao    gl.VertexArray
}

// NewRecorder creates a new Recorder object
func (e *Engine) NewRecorder() (*Recorder, error) {
	glctx3, ok := e.glctx.(gl.Context3)
	if !ok {
		return nil, errors.New("Graphics backend does not support Recorder objects")
	}

	vao := glctx3.CreateVertexArray()

	return &Recorder{
		engine: e,
		vao:    vao,
	}, nil
}

// Release releases the Recorder
// Using the Recorder after calling this function results in a panic
func (r *Recorder) Release() {
	r.engine.glctx.(gl.Context3).DeleteVertexArray(r.vao)
	r.engine = nil
}

// Activate activates the Recorder (to switch to the active state and start recording it)
func (r *Recorder) Activate() {
	// TODO Check if already bound (prevent overhead?)
	r.engine.glctx.(gl.Context3).BindVertexArray(r.vao)
}

// Deactivate deactivates the Recorder
func (r *Recorder) Deactivate() {
	r.engine.glctx.(gl.Context3).BindVertexArray(gl.VertexArray{Value: 0})
}
