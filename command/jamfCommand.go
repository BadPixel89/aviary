package command

import (
	"fmt"
)

var _ = RegisterCommand(JamfCommand{})

type JamfCommand struct{}

func (j JamfCommand) Run(args []string) error {
	fmt.Println("you wish you could interact with jamf")
	return nil
}

func (j JamfCommand) Help() {
	fmt.Println("interact with jamf (when this is finished)")
}

func (j JamfCommand) Name() string {
	return "jamf"
}
