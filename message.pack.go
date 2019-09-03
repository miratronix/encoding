package lingo

import (
	"bytes"
	"github.com/miratronix/msgpack"
)

// messagePack defines a basic message pack encoder/decoder
type messagePack struct{}

// Encode encodes the given interface into a byte slice
func (m *messagePack) Encode(data interface{}) ([]byte, error) {
	var buf bytes.Buffer
	err := msgpack.NewEncoder(&buf).UseJSONTag(true).Encode(data)
	return buf.Bytes(), err
}

// Decode decodes the given byte slice into the supplied interface
func (m *messagePack) Decode(data []byte, result interface{}) error {
	return msgpack.NewDecoder(bytes.NewReader(data)).UseJSONTag(true).Decode(result)
}
