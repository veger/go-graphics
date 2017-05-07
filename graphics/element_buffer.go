package graphics

import "golang.org/x/mobile/gl"

// ElementBuffer is a graphic-related buffer that holds element data (TriangleIndices), used to communicate between engine and graphical backend
type ElementBuffer struct {
	*Buffer

	size int
}

// TriangleIndices contains 3 indices to a databuffer that form a triangle
type TriangleIndices [3]uint32

// NewElementBuffer create a buffer of the BufferType_Element type
// it provides some additional functionalities compares to a regular Buffer
func (e *Engine) NewElementBuffer(elementData []TriangleIndices) (*ElementBuffer, error) {
	bufferData, err := castTriangleIndicesToBytes(elementData)
	if err != nil {
		return nil, err
	}

	b, err := e.NewBuffer(BufferTypeElement, bufferData)
	if err != nil {
		return nil, err
	}
	return &ElementBuffer{
		Buffer: b,
		size:   len(elementData) * 3, // 3 indices per triangle
	}, nil
}

// Render renders the contents of the element buffer
func (eb *ElementBuffer) Render() {
	eb.engine.glctx.DrawElements(gl.TRIANGLES, eb.size, gl.UNSIGNED_INT, 0)
}
