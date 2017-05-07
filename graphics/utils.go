package graphics

import (
	"bytes"
	"encoding/binary"
)

// FloatSize is the size of a 32-bit float
const FloatSize = 4

// CastUint32sToBytes create byte array from uint32 array
func castUint32sToBytes(ints []uint32) ([]byte, error) {
	buf := new(bytes.Buffer)
	err := binary.Write(buf, binary.LittleEndian, ints)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
