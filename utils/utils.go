package utils

import (
	"os/user"
	"regexp"
	"strconv"

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

// Uid looks up a username and returns uid
func Uid(username string) (uint32, error) {
	u, err := user.Lookup(username)
	if err != nil {
		return 0, err
	}
	uid, err := strconv.ParseUint(u.Uid, 10, 32)
	if err != nil {
		return 0, err
	}
	return uint32(uid), nil
}
