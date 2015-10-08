package commander

// Builds Command from cmd name and arguments
func BuildCmd(cmd string, args ...string) *Command {
	return NewCommand(cmd, args...)

}
