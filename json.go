package lingo

import (
	"encoding/json"
)

// jsonEncoding defines a basic JSON encoder/decoder
type jsonEncoding struct{}

// Encode encodes the given interface into a byte slice
func (j *jsonEncoding) Encode(data interface{}) ([]byte, error) {
	return json.Marshal(data)
}

// Decode decodes the given byte slice into the supplied interface
func (j *jsonEncoding) Decode(data []byte, result interface{}) error {
	return json.Unmarshal(data, result)
}
