// Package file implements various functions that provide helpers
// to query various aspects of operating system files
package file

import (
	"bufio"
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"os"
	"regexp"
	"syscall"
	"time"

	"github.com/milosgajdos83/servpeek/resource"
	"github.com/milosgajdos83/servpeek/utils"
)

const ModeRegular = ^os.ModeType

func withOsFile(path string, fn func(file *os.File) error) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()
	return fn(file)
}

func isTrueOrError(ok bool, err error) error {
	if ok {
		return nil
	}
	return err
}

// IsRegular checks if provided file is a regular file
// It returns error if os.Stat returns error
func IsRegular(f resource.Filer) error {
	return IsMode(f, ModeRegular)
}

// IsDirectory checks if provided file is a directory
// It returns error if os.Stat returns error
func IsDirectory(f resource.Filer) error {
	return IsMode(f, os.ModeDir)
}

// IsBlockDevice checks if provided file is a block device file
// It returns error if os.Stat returns error
func IsBlockDevice(f resource.Filer) error {
	return IsMode(f, os.ModeDevice)
}

// IsCharDevice checks if provided file is a character device file
// It returns error if os.Stat returns error
func IsCharDevice(f resource.Filer) error {
	return IsMode(f, os.ModeCharDevice)
}

// IsPipe checks if provided file is a named pipe file
// It returns error if os.Stat returns error
func IsPipe(f resource.Filer) error {
	return IsMode(f, os.ModeNamedPipe)
}

// IsSocket checks if provided file is a socket
// It returns error if os.Stat returns error
func IsSocket(f resource.Filer) error {
	return IsMode(f, os.ModeSocket)
}

// IsSymlink checks if provided file is a symbolic link
// It returns error if os.Stat returns error
func IsSymlink(f resource.Filer) error {
	return IsMode(f, os.ModeSymlink)
}

var verboseNames = map[os.FileMode]string{
	os.ModeSymlink:   "symlink",
	os.ModeSocket:    "socket",
	os.ModeDir:       "directory",
	os.ModeNamedPipe: "named pipe",
	ModeRegular:      "regular file",
}

// IsMode checks if the provided file has the same mode as the one passed in via paramter
// It returns error if os.Stat returns error
func IsMode(f resource.Filer, mode os.FileMode) error {
	fi, err := f.Info()
	if err != nil {
		return err
	}

	if verboseMode, ok := verboseNames[mode]; ok {
		err = fmt.Errorf("%s file is not a %s", f, verboseMode)
	} else {
		err = fmt.Errorf("%s file is not mode %v", f, mode)
	}
	return isTrueOrError(fi.Mode()&mode != 0, err)
}

// IsOwnedBy checks if the provided file is owned by username user
// It returns an error if os.Stat returns error
func IsOwnedBy(f resource.Filer, username string) error {
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

// IsGrupedInto checks if the provided file is owned by groupname group
// It returns an error if os.Stat returns error
func IsGrupedInto(f resource.Filer, groupname string) error {
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

// LinksTo checks if the provided file is a symlink which links to path
// It returs error if the link can't be read
func LinksTo(f resource.Filer, path string) error {
	dst, err := os.Readlink(path)
	if err != nil {
		return err
	}
	if dst == f.Path() {
		return nil
	}
	return fmt.Errorf("%s does not link to %s", f, path)
}

// Md5 checks if the provided file md5 checksum is the same as the one passed in as paramter
// It returs error if the provided file can't be open
func Md5(f resource.Filer, sum string) error {
	return withOsFile(f.Path(), func(file *os.File) error {
		hasher := md5.New()
		if _, err := io.Copy(hasher, file); err != nil {
			return nil
		}
		if sum == hex.EncodeToString(hasher.Sum(nil)) {
			return nil
		}
		return fmt.Errorf("%s md5sum does not equal to %s", f, sum)
	})
}

// Sha256 checks if the provided file sha256 checksum is the same as the one passed in as paramter
// It returs error if the provided file can't be open
func Sha256(f resource.Filer, sum string) error {
	return withOsFile(f.Path(), func(file *os.File) error {
		hasher := sha256.New()
		if _, err := io.Copy(hasher, file); err != nil {
			return nil
		}
		if sum == hex.EncodeToString(hasher.Sum(nil)) {
			return nil
		}
		return fmt.Errorf("%s sha256 does not equal to %s", f, sum)
	})
}

// Size checks if the provided file byte size is the same as the one passed in as paramter
// It returns error if os.Stat returns error
func Size(f resource.Filer, size int64) error {
	fi, err := f.Info()
	if err != nil {
		return err
	}
	if fi.Size() == size {
		return nil
	}
	return fmt.Errorf("%s size is different from %d", f, size)
}

// ModTimeAfter checks if the provided file modification time is older than the one passed in as paramter
// It returns error if os.Stat returns error
func ModTimeAfter(f resource.Filer, mtime time.Time) error {
	fi, err := f.Info()
	if err != nil {
		return err
	}
	if fi.ModTime().After(mtime) {
		return nil
	}
	return fmt.Errorf("%s modification time is bigger than %s", f, mtime)
}

// Contains checks if the provided file content can be matched with any of the regexps
// passed in as paramter. It returs error if the provided file can't be open
func Contains(f resource.Filer, contents ...*regexp.Regexp) error {
	return withOsFile(f.Path(), func(file *os.File) error {
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			for _, content := range contents {
				if content.Match(scanner.Bytes()) {
					return nil
				}
			}
		}
		return fmt.Errorf("%s does not match any provided regular expression", f)
	})
}
