package command

type Command interface {
	Run(args []string) error
	Help()
	Name() string
}
