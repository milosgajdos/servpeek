package commander

import (
	"github.com/milosgajdos83/servpeek/utils/command"
)

// BuildCmd builds Command from cmd name and arguments
func BuildCmd(cmd string, args ...string) *command.Command {
	return command.NewCommand(cmd, args...)

}
