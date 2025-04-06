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
	if len(args) <= 2 {
		fmt.Println("please provide an operation and two numbers")
		return nil
	}
	a, era := strconv.Atoi(args[1])
	b, erb := strconv.Atoi(args[2])
	if era != nil {
		fmt.Println("your first number did not parse")
		return era
	}
	if erb != nil {
		fmt.Println("your second number did not parse")
		return erb
	}

	switch args[0] {
	case "-add":
		fmt.Println(add(a, b))
	case "-sub":
		fmt.Println(subtract(a, b))
	case "-mul":
		fmt.Println(multiply(a, b))
	case "-div":
		fmt.Println(divide(a, b))
	default:
		fmt.Println("your operation did not parse")
	}
	return nil
}
func add(n1 int, n2 int) int {
	return (n1 + n2)
}
func subtract(n1 int, n2 int) int {
	return (n1 - n2)
}
func multiply(n1 int, n2 int) int {
	return n1 * n2
}
func divide(n1 int, n2 int) float32 {
	n1f := float32(n1)
	n2f := float32(n2)
	if n1 == 0 || n2 == 0 {
		return 0
	}
	return (n1f / n2f)
}
func (e ExampleCommand) Help() {
	fmt.Println("description:")
	fmt.Println("\tperform math operations on the numbers passed in to this command")
	fmt.Println("usage:")
	fmt.Println("\texample [operation] <n1> <n2>")
	fmt.Println("operation flag can be one of:")
	fmt.Println("\t-add -sub -mul -div")
}
func (e ExampleCommand) Name() string {
	return "example"
}
