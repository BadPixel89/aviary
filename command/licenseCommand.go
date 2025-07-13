package command

import (
	"fmt"
)

var _ = RegisterCommand(LicenseCommand{})

type LicenseCommand struct{}

func (l LicenseCommand) Run(args []string) error {
	//figure out how to embed and read a text file
	//fill it with GPL and spuff it out
	fmt.Println("print the GPL to console when I figure that stuff out")
	return nil
}
func (l LicenseCommand) Help() {
	fmt.Println("print the GPL to console")
}
func (l LicenseCommand) Name() string {
	return "license"
}
