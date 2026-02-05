package command

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
