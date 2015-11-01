// Package utils provides useful utility functions
package utils

import (
	"fmt"
	"os/user"
	"strconv"

	"github.com/milosgajdos83/servpeek/utils/command"
	"github.com/milosgajdos83/servpeek/utils/group"
)

// BuildCmd builds Command from cmd name and arguments
func BuildCmd(cmd string, args ...string) *command.Command {
	return command.NewCommand(cmd, args...)

}

// RoleToID converts username/groupname to their numeric id representations: uid/gid
// Returns error if the role is not supported, usernamd/groupname have not been found
func RoleToID(role string, name string) (uint64, error) {
	var id string
	switch role {
	case "user", "User":
		user, err := user.Lookup(name)
		if err != nil {
			return 0, err
		}
		id = user.Uid
	case "group", "Group":
		group, err := group.Lookup(name)
		if err != nil {
			return 0, err
		}
		id = group.Gid
	default:
		return 0, fmt.Errorf("Unsupported role: %s", role)
	}
	// Parse uid/gid
	numID, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		return 0, err
	}

	return numID, nil
}
