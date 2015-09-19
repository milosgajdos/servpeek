package utils

import "github.com/milosgajdos83/servpeek/utils/command"

// Ugly hack to parse pkgQueryInfo
type ParseHints struct {
	ListPrefix      string
	ListMinFields   int
	ListVersionIdx  int
	QueryPrefix     string
	QueryMinFields  int
	QueryVersionIdx int
}

// Builds Command from cmd name and arguments
func BuildCmd(cmd string, args ...string) *command.Command {
	return command.NewCommand(cmd, args...)

}
