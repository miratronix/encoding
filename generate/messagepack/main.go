package main

import (
	"fmt"
	"github.com/miratronix/zebrapack/cfg"
	"github.com/miratronix/zebrapack/gen"
	"github.com/miratronix/zebrapack/parse"
	"github.com/miratronix/zebrapack/printer"
	"os"
	"strings"
)

func main() {

	// Construct the config
	config := &cfg.ZebraConfig{
		GoFile:   os.Getenv("GOFILE"),
		Encode:   false,
		Tests:    false,
		Marshal:  true,
		UseMsgp2: true,
	}

	// Set the mode
	mode := gen.Marshal | gen.Unmarshal | gen.Size | gen.FieldsEmpty

	// Set the file name
	gen.SetFilename(config.GoFile)

	fmt.Printf(">>> Generating: \"%s\"\n", config.GoFile)
	fs, err := parse.File(config)
	if err != nil {
		panic(err)
	}

	// Compute the output filename
	filename := strings.TrimSuffix(config.GoFile, ".go") + "_gen.go"

	// Write the file
	err = printer.PrintFile(filename, fs, mode, config, config.GoFile)
	if err != nil {
		panic(err)
	}
}
