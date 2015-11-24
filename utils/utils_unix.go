// +build darwin dragonfly freebsd !android,linux netbsd openbsd solaris
// +build cgo

package utils

import (
	"fmt"
	"os/user"
	"strconv"
	"strings"

	"github.com/milosgajdos83/servpeek/utils/group"
)

func roleToID(role string, name string) (uint64, error) {
	var id string
	switch strings.ToLower(role) {
	case "user":
		user, err := user.Lookup(name)
		if err != nil {
			return 0, err
		}
		id = user.Uid
	case "group":
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
