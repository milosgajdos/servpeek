// +build darwin dragonfly freebsd !android,linux netbsd openbsd solaris
// +build cgo

package file

import (
	"fmt"
	"syscall"

	"github.com/milosgajdos83/servpeek/utils"
)

func isOwnedBy(f Filer, username string) error {
	fi, err := f.Info()
	if err != nil {
		return nil
	}
	uid, err := utils.RoleToID("user", username)
	if err != nil {
		return err
	}
	return isTrueOrError(fi.Sys().(*syscall.Stat_t).Uid == uint32(uid), fmt.Errorf("%s file is not owned by %s", f, username))
}

func isGrupedInto(f Filer, groupname string) error {
	fi, err := f.Info()
	if err != nil {
		return err
	}
	gid, err := utils.RoleToID("group", groupname)
	if err != nil {
		return err
	}
	return isTrueOrError(fi.Sys().(*syscall.Stat_t).Gid == uint32(gid), fmt.Errorf("%s file is not grouped to %s", f, groupname))
}
