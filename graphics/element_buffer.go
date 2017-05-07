package graphics

import "golang.org/x/mobile/gl"

// ElementBuffer is a graphic-related buffer that holds element data, used to communicate between engine and graphical backend
type ElementBuffer struct {
	*Buffer

	size int
}

// NewElementBuffer create a buffer of the BufferType_Element type
// it provides some additional functionalities compares to a regular Buffer
// TODO change elementData type to a 'array of triangle indices'
func (e *Engine) NewElementBuffer(elementData []uint32) (*ElementBuffer, error) {
	bufferData, err := castUint32sToBytes(elementData)
	if err != nil {
		return nil, err
	}

	b, err := e.NewBuffer(BufferTypeElement, bufferData)
	if err != nil {
		return nil, err
	}
	return &ElementBuffer{
		Buffer: b,
		size:   len(elementData),
	}, nil
}

// Render renders the contents of the element buffer
func (eb *ElementBuffer) Render() {
	eb.engine.glctx.DrawElements(gl.TRIANGLES, eb.size, gl.UNSIGNED_INT, 0)
}
