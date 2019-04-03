package lingo

// Encoding defines an encoder/decoder interface
type Encoding interface {
	Encode(interface{}) ([]byte, error)
	Decode([]byte, interface{}) error
}

// JSON exposes a pre-built JSON encoding
var JSON Encoding = &jsonEncoding{}

// MessagePack exposes a pre-built message pack encoding
var MessagePack Encoding = &messagePack{}

// Map exposes a pre-built map encoding
var Map = &mapEncoding{}
