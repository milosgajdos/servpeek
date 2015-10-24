package utils

import (
	"fmt"
	"os/user"
	"strconv"

	"github.com/milosgajdos83/servpeek/utils/group"
)

// Converts username/groupname to their numeric representations
// Returns error if the role is not supported, usernamd/groupname have not been found
func RoleToId(role string, name string) (uint64, error) {
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
	numId, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		return 0, err
	}

	return numId, nil
}
