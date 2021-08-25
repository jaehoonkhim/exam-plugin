package main

import (
	"fmt"
	"plugin"
)

func main() {
	libPath := "./default/default.so"
	plug, err := plugin.Open(libPath)
	if err != nil {
		panic(err)
	}

	symbol, err := plug.Lookup("ReservedCommander")
	if err != nil {
		panic(err)
	}

	reservedCmd, ok := symbol.(ReservedCommander)
	if !ok {
		panic(fmt.Errorf("%s", "not supported symbol"))
	}

	reservedCmd.Execute()
}
