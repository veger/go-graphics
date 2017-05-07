package graphics

import (
	"fmt"

	"golang.org/x/mobile/gl"
)

// Buffer is a graphic-related buffer, used to communicate between engine and graphical backend
type Buffer struct {
	engine     *Engine
	bufferType BufferType
	glBuffer   gl.Buffer
}

// BufferType defines the graphics buffer types
type BufferType int

const (
	// BufferTypeData holds (vertex) data
	BufferTypeData BufferType = iota

	// BufferTypeElement holds element data (indices to/of objects in a data buffer)
	BufferTypeElement
)

var bufferTypes = map[BufferType]gl.Enum{
	BufferTypeData:    gl.ARRAY_BUFFER,
	BufferTypeElement: gl.ELEMENT_ARRAY_BUFFER,
}

// NewBuffer creates a buffer using/for the underlying graphics rendering structure
// Note that the buffer gets (and stays) bound (for the given BufferType)
func (e *Engine) NewBuffer(bufferType BufferType, verticesData []byte) (*Buffer, error) {
	b := &Buffer{
		engine:     e,
		bufferType: bufferType,
		glBuffer:   e.glctx.CreateBuffer(),
	}

	glBufferType, ok := bufferTypes[bufferType]
	if !ok {
		return nil, fmt.Errorf("Unknown buffer type: %d", bufferType)
	}

	b.bindBuffer()
	e.glctx.BufferData(glBufferType, verticesData, gl.STATIC_DRAW)

	return b, nil
}

// Release releases the Buffer
// Using the Buffer after calling this function results in a panic
func (b *Buffer) Release() {
	b.engine.glctx.DeleteBuffer(b.glBuffer)
	b.engine = nil
}

// BindVariable makes the (32-bit floating point) variable available in shader programs ofr the given position.
// variableSize is the size of the variable, allowed range 1 to 4
// dataSize is the number floating point entries for each data entry
// pointer is the location of the data that needs to be bound in the data entry
func (b *Buffer) BindVariable(position uint, variableSize, dataSize, pointer int) {
	b.bindBuffer()
	p := gl.Attrib{Value: position}
	b.engine.glctx.VertexAttribPointer(p, variableSize, gl.FLOAT, false, dataSize*FloatSize, pointer*FloatSize)
	b.engine.glctx.EnableVertexAttribArray(p)
}

func (b *Buffer) bindBuffer() {
	// TODO Check if already bound (prevent overhead?)
	b.engine.glctx.BindBuffer(bufferTypes[b.bufferType], b.glBuffer)
}

func (b *Buffer) unbindBuffer() {
	b.engine.glctx.BindBuffer(bufferTypes[b.bufferType], gl.Buffer{Value: 0})
}
