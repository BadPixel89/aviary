package command

// use this API to get json info about a MAC
// https://api.maclookup.app/v2/macs/<MAC ADDRESS HERE>

var _ = RegisterCommand(MacCommand{})

type MacCommand struct{}

func (m MacCommand) Run(args []string) error {
	return nil
}

func (m MacCommand) Help() {

}

func (m MacCommand) Name() string {
	return "mac"
}
