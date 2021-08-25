package main

import (
	"os"
	"plugin"
)

func main() {
	libPath := "./default/default.so"
	plug, err := plugin.Open(libPath)
	if err != nil {
		os.Exit(1)
	}

	symbol, err := plug.Lookup("ReservedCommander")
	if err != nil {
		os.Exit(1)
	}

	reservedCmd, ok := symbol.(ReservedCommander)
	if !ok {
		os.Exit(1)
	}

	reservedCmd.Execute()
}
