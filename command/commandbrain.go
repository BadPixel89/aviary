package command

import (
	"errors"
	"fmt"
)

var commands = make(map[string]Command)

func RegisterCommand(cmd Command) error {
	_, found := commands[cmd.Name()]
	if found {
		return errors.New("command already registered")
	}

	commands[cmd.Name()] = cmd
	return nil
}

func Parse(args []string) error {
	first := args[0]
	cmd, ok := commands[first]

	if !ok {
		return errors.New("command not found")
	}

	if len(args) == 1 {
		err := cmd.Run(nil)
		return err
	}
	if args[1] == "help" || args[1] == "-h" {
		cmd.Help()
		return nil
	}
	err := cmd.Run(args[1:])
	return err
}

func List() {
	fmt.Println("available commands:")
	for name, _ := range commands {
		fmt.Println("\t" + name)
	}
	fmt.Println("type a command followed by help or -h to view more info about it")
}
