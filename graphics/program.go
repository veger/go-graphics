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

// UniformLocation is the location of a uniform in a Program
type UniformLocation int32

// NewProgram creates a new program, optionally provide custom vertex and/or fragment shader(s)
// The default program uses position (index = 0) and texture (index = 3) data
func (e *Engine) NewProgram(vertexShader, fragmentShader string) (*Program, error) {
	if vertexShader == "" {
		vertexShader = vertexShaderDefault
	}
	if fragmentShader == "" {
		fragmentShader = fragmentShaderDefault
	}
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

// GetUniformLocation gets the location of a uniform in the program
func (p *Program) GetUniformLocation(name string) UniformLocation {
	u := p.engine.glctx.GetUniformLocation(p.shader, name)
	return UniformLocation(u.Value)
}

// SetUnitformMatrix4 sets the value of the uniform at location to mat
func (p *Program) SetUnitformMatrix4(location UniformLocation, mat []float32) {
	p.engine.glctx.UniformMatrix4fv(gl.Uniform{Value: int32(location)}, mat)
}

const vertexShaderDefault = `#version 300 es

layout (location = 0) in vec3 position;
layout (location = 1) in vec2 texCoord;

out vec2 TexCoord;

void main() {
	gl_Position = vec4(position, 1.0);
	TexCoord = texCoord;
}`

const fragmentShaderDefault = `#version 300 es
precision mediump float;

in vec2 TexCoord;

out vec4 color;

uniform sampler2D ourTexture;

void main() {
	color = texture(ourTexture, TexCoord);
}`
