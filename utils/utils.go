package utils

import (
	"regexp"

	"github.com/milosgajdos83/servpeek/utils/command"
)

// Ugly hack to parse Command outpyt into *PkgInfo
type ParseHints struct {
	ListFilter  *regexp.Regexp
	ListMatch   *regexp.Regexp
	QueryFilter *regexp.Regexp
	QueryMatch  *regexp.Regexp
}

// Builds Command from cmd name and arguments
func BuildCmd(cmd string, args ...string) *command.Command {
	return command.NewCommand(cmd, args...)

}
