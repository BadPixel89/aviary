package command

var _ = RegisterCommand(DiscLoadCommand{})

type DiscLoadCommand struct{}

func (c DiscLoadCommand) Run(args []string) error {
	// calculate disk load here

	return nil
}

func (c DiscLoadCommand) Help() {

}

func (c DiscLoadCommand) Name() string {
	return "discload"
}
