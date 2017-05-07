package graphics

import (
	"golang.org/x/mobile/exp/gl/glutil"
	"golang.org/x/mobile/gl"
)

// Program is a graphics program that is send to the backend to render the Buffer data
type Program struct {
	engine *Engine
	shader gl.Program
}

// NewProgram creates a new (default) program, that uses position (index = 0) and texture (index = 3) data
func (e *Engine) NewProgram() (*Program, error) {
	shader, err := glutil.CreateProgram(e.glctx, vertexShader, fragmentShader)
	if err != nil {
		return nil, err
	}

	return &Program{
		engine: e,
		shader: shader,
	}, nil
}

// Release releases the Program
// Using the Program after calling this function results in a panic
func (p *Program) Release() {
	p.engine.glctx.DeleteProgram(p.shader)
	p.engine = nil
}

// Activate actives the program (and deactives the current one)
func (p *Program) Activate() {
	p.engine.glctx.UseProgram(p.shader)
}

const vertexShader = `#version 300 es

layout (location = 0) in vec3 position;
layout (location = 1) in vec2 texCoord;

out vec2 TexCoord;

void main() {
	gl_Position = vec4(position.x, position.y, position.z, 1.0);
	TexCoord = texCoord;
}`

const fragmentShader = `#version 300 es
precision mediump float;

in vec2 TexCoord;

out vec4 color;

uniform sampler2D ourTexture;

void main() {
	color = texture(ourTexture, TexCoord);
}`
