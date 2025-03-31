package command

import (
	"fmt"
	"strconv"
)

// between dashes is basically the boilerplate to get the command to load and be usable
// interface is not explicitly defined and inherited in Go, you fulfill it by implementing it
// the func (e examplecommand) part of functions makes the functions act as if they were on
// the object in c# terms cmd.Run() for example
// -----
// This code is run even though we don't explicitly instantiate, call, or use it
var _ = RegisterCommand(ExampleCommand{})

// This is basically where we name our command in code (ping or whatever)
type ExampleCommand struct{}

// -----

func (e ExampleCommand) Run(args []string) error {
	if len(args) <= 1 {
		fmt.Println("please provide two numbers to add")
		return nil
	}

	a, era := strconv.Atoi(args[0])
	b, erb := strconv.Atoi(args[1])
	if era != nil {
		fmt.Println("your first number did not parse")
		return era
	}
	if erb != nil {
		fmt.Println("your second number did not parse")
		return erb
	}

	fmt.Println(a + b)
	return nil
}
func (e ExampleCommand) Help() {
	fmt.Println("\tadd two numbers")
}
func (e ExampleCommand) Name() string {
	return "example"
}
