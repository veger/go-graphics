package utils

import (
	"bytes"
	"encoding/binary"
)

// CastToBytes casts data to their (LittleEndian) byte representation
func CastToBytes(data interface{}) ([]byte, error) {
	buf := new(bytes.Buffer)
	err := binary.Write(buf, binary.LittleEndian, data)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
