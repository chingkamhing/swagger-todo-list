package main

import (
	"os"

	flags "github.com/jessevdk/go-flags"
)

var parser = flags.NewParser(nil, flags.Default)

func main() {
	if _, err := parser.Parse(); err != nil {
		code := 1
		if fe, ok := err.(*flags.Error); ok {
			if fe.Type == flags.ErrHelp {
				code = 0
			}
		}
		os.Exit(code)
	}
}
