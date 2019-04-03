package lingo

import (
	"github.com/miratronix/structs"
	"github.com/mitchellh/mapstructure"
)

// mapEncoding defines a simple map encoder/decoder
type mapEncoding struct{}

// Encode encodes a structure to a map
func (m *mapEncoding) Encode(data interface{}) map[string]interface{} {
	wrap := structs.New(data)
	wrap.TagName = "json"
	return wrap.Map()
}

// Decode decodes a map to a structure
func (m *mapEncoding) Decode(data interface{}, result interface{}) error {
	config := &mapstructure.DecoderConfig{
		Metadata: nil,
		TagName:  "json",
		Result:   result,
	}

	decoder, err := mapstructure.NewDecoder(config)
	if err != nil {
		return err
	}

	return decoder.Decode(data)
}
