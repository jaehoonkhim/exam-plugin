package main

import "fmt"

type Command struct{}

func (c *Command) Execute() {
	fmt.Println("default ReservedCommand(Execute)")
}

var ReservedCommander Command
