package graphics

import (
	"bytes"
	"encoding/binary"
)

// FloatSize is the size of a 32-bit float
const FloatSize = 4

// CastUint32sToBytes create byte array from a TriangleIndices array
func castTriangleIndicesToBytes(indices []TriangleIndices) ([]byte, error) {
	buf := new(bytes.Buffer)
	err := binary.Write(buf, binary.LittleEndian, indices)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
