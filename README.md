# lingo
A basic collection of encoding/decoding packages and tools.

## Usage
Lingo ships with three encodings, all of which use the standard `json` struct tag:

* JSON:
```go
import "github.com/miratronix/lingo"
e := lingo.JSON

// Encode a structure as bytes, using the 'json' struct tag
bytes, err := e.Encode(someStructure)

// Decode bytes into a structure, using the 'json' struct tag
err := e.Decode(bytes, &someStructure)
```

* Message Pack:
```go
import "github.com/miratronix/lingo"
e := lingo.MessagePack

// Encode a structure as bytes, using the 'json' struct tag
bytes, err := e.Encode(someStructure)

// Decode bytes into a structure, using the 'json' struct tag
err := e.Decode(bytes, &someStructure)
```

* Map:
```go
import "github.com/miratronix/lingo"
e := lingo.Map

// Encode a structure into a map, using the 'json' struct tag
mapped := e.Encode(&someStructure)

// Decode a map into a structure, using the 'json' struct tag
err := e.Decode(someMap, &someStructure)
```
