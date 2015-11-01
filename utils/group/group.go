// Package group provides utility functions that deal with OS groups

// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build darwin dragonfly freebsd !android,linux netbsd openbsd solaris
// +build cgo

package group

import "os/user"

var implemented = true

// Group is an OS group
type Group struct {
	Gid  string // group id
	Name string // group name
}

// Members returns a slice of OS groups
// It returns error if the group lookup fails or is unsupported
func (g *Group) Members() ([]string, error) {
	return groupMembers(g)
}

// UnknownGroupIDError is returned by LookupGroupId when
// a group cannot be found.
type UnknownGroupIDError string

func (e UnknownGroupIDError) Error() string {
	return "group: unknown groupid " + string(e)
}

// UnknownGroupError is returned by LookupGroup when
// a group cannot be found.
type UnknownGroupError string

func (e UnknownGroupError) Error() string {
	return "group: unknown group " + string(e)
}

// Current returns the current group.
func Current() (*Group, error) {
	return currentGroup()
}

// Lookup looks up a group by name. If the group cannot be found, the
// returned error is of type UnknownGroupError.
func Lookup(groupname string) (*Group, error) {
	return lookupGroup(groupname)
}

// LookupID looks up a group by groupid. If the group cannot be found, the
// returned error is of type UnknownGroupIDError.
func LookupID(gid string) (*Group, error) {
	return lookupGroupID(gid)
}

// UserInGroup reports whether the user is a member of the given group.
func UserInGroup(u *user.User, g *Group) (bool, error) {
	return userInGroup(u, g)
}
