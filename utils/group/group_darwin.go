// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build darwin
// +build cgo

package group

/*
#include <unistd.h>
#include <sys/types.h>
#include <pwd.h>
#include <grp.h>
#include <stdlib.h>

static int mygetgrouplist(const char *user, int group, int *groups,
	int *ngroups) {
  return getgrouplist(user, group, groups, ngroups);
}

static inline gid_t group_at(int i, gid_t *groups) {
  return groups[i];
}
*/
import "C"
import (
	"fmt"
	"os/user"
	"strconv"
	"syscall"
	"unsafe"
)

func userInGroup(u *user.User, g *Group) (bool, error) {
	if u.Gid == g.Gid {
		return true, nil
	}
	gid, err := strconv.Atoi(u.Gid)
	if err != nil {
		return false, err
	}

	nameC := C.CString(u.Username)
	defer C.free(unsafe.Pointer(nameC))
	groupC := C.int(gid)
	ngroupsC := C.int(0)

	C.mygetgrouplist(nameC, groupC, nil, &ngroupsC)
	ngroups := int(ngroupsC)

	groups := C.malloc(C.size_t(int(unsafe.Sizeof(groupC)) * ngroups))
	defer C.free(groups)

	rv := C.mygetgrouplist(nameC, groupC, (*C.int)(groups), &ngroupsC)
	if rv == -1 {
		return false, fmt.Errorf("user: membership of %s in %s: %s", u.Username, g.Name, syscall.Errno(rv))
	}

	ngroups = int(ngroupsC)
	for i := 0; i < ngroups; i++ {
		gid := C.group_at(C.int(i), (*C.gid_t)(groups))
		if g.Gid == strconv.Itoa(int(gid)) {
			return true, nil
		}
	}
	return false, nil
}
