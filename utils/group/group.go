// +build darwin dragonfly freebsd !android,linux netbsd openbsd solaris
// +build cgo

package group

import "os/user"

var implemented = true

type Group struct {
	Gid  string // group id
	Name string // group name
}

func (g *Group) Members() ([]string, error) {
	return groupMembers(g)
}

// UnknownGroupIdError is returned by LookupGroupId when
// a group cannot be found.
type UnknownGroupIdError string

func (e UnknownGroupIdError) Error() string {
	return "group: unknown groupid " + string(e)
}

// UnknownGroupError is returned by LookupGroup when
// a group cannot be found.
type UnknownGroupError string

func (e UnknownGroupError) Error() string {
	return "group: unknown group " + string(e)
}

// CurrentGroup returns the current group.
func Current() (*Group, error) {
	return currentGroup()
}

// LookupGroup looks up a group by name. If the group cannot be found, the
// returned error is of type UnknownGroupError.
func Lookup(groupname string) (*Group, error) {
	return lookupGroup(groupname)
}

// LookupGroupId looks up a group by groupid. If the group cannot be found, the
// returned error is of type UnknownGroupIdError.
func LookupId(gid string) (*Group, error) {
	return lookupGroupId(gid)
}

// UserInGroup reports whether the user is a member of the given group.
func UserInGroup(u *user.User, g *Group) (bool, error) {
	return userInGroup(u, g)
}
