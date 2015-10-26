// package file implements various functions that provide helpers
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

func withOsFile(path string, fn func(file *os.File) error) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()
	return fn(file)
}

func withFileInfo(path string, fn func(fi os.FileInfo) error) error {
	fi, err := os.Stat(path)
	if err != nil {
		return err
	}
	return fn(fi)
}

// IsRegular checks if provided file is a regular file
// It returns error if os.Stat returns error
func IsRegular(f *resource.File) error {
	return withFileInfo(f.Path, func(fi os.FileInfo) error {
		if fi.Mode().IsRegular() {
			return nil
		}
		return fmt.Errorf("%s file is not a regular file", f)
	})
}

// IsDirectory checks if provided file is a directory
// It returns error if os.Stat returns error
func IsDirectory(f *resource.File) error {
	return withFileInfo(f.Path, func(fi os.FileInfo) error {
		if fi.IsDir() {
			return nil
		}
		return fmt.Errorf("%s file is not a directory", f)
	})
}

// IsBlockDevice checks if provided file is a block device file
// It returns error if os.Stat returns error
func IsBlockDevice(f *resource.File) error {
	return withFileInfo(f.Path, func(fi os.FileInfo) error {
		if fi.Mode()&os.ModeDevice != 0 {
			return nil
		}
		return fmt.Errorf("%s file is not a block device", f)
	})
}

// IsCharDevice checks if provided file is a character device file
// It returns error if os.Stat returns error
func IsCharDevice(f *resource.File) error {
	return withFileInfo(f.Path, func(fi os.FileInfo) error {
		if fi.Mode()&os.ModeCharDevice != 0 {
			return nil
		}
		return fmt.Errorf("%s file is not a character device", f)
	})
}

// IsPipe checks if provided file is a named pipe file
// It returns error if os.Stat returns error
func IsPipe(f *resource.File) error {
	return withFileInfo(f.Path, func(fi os.FileInfo) error {
		if fi.Mode()&os.ModeNamedPipe != 0 {
			return nil
		}
		return fmt.Errorf("%s file is not a named pipe", f)
	})
}

// IsSocket checks if provided file is a socket
// It returns error if os.Stat returns error
func IsSocket(f *resource.File) error {
	return withFileInfo(f.Path, func(fi os.FileInfo) error {
		if fi.Mode()&os.ModeSocket != 0 {
			return nil
		}
		return fmt.Errorf("%s file is not a socket", f)
	})
}

// IsSymlink checks if provided file is a symbolic link
// It returns error if os.Stat returns error
func IsSymlink(f *resource.File) error {
	return withFileInfo(f.Path, func(fi os.FileInfo) error {
		if fi.Mode()&os.ModeSymlink != 0 {
			return nil
		}
		return fmt.Errorf("%s file is not a symlink", f)
	})
}

// IsMode checks if the provided file has the same mode as the one passed in via paramter
// It returns error if os.Stat returns error
func IsMode(f *resource.File, mode os.FileMode) error {
	return withFileInfo(f.Path, func(fi os.FileInfo) error {
		if fi.Mode()&mode != 0 {
			return nil
		}
		return fmt.Errorf("%s file is not mode %v", f, mode)
	})
}

// IsOwnedBy checks if the provided file is owned by username user
// It returns an error if os.Stat returns error
func IsOwnedBy(f *resource.File, username string) error {
	return withFileInfo(f.Path, func(fi os.FileInfo) error {
		uid, err := utils.RoleToId("user", username)
		if err != nil {
			return err
		}
		if fi.Sys().(*syscall.Stat_t).Uid == uint32(uid) {
			return nil
		}
		return fmt.Errorf("%s file is not owned by %s", f, username)
	})
}

// IsGrupedInto checks if the provided file is owned by groupname group
// It returns an error if os.Stat returns error
func IsGrupedInto(f *resource.File, groupname string) error {
	return withFileInfo(f.Path, func(fi os.FileInfo) error {
		gid, err := utils.RoleToId("group", groupname)
		if err != nil {
			return err
		}
		if fi.Sys().(*syscall.Stat_t).Gid == uint32(gid) {
			return nil
		}
		return fmt.Errorf("%s file is not grouped to %s", f, groupname)
	})
}

// LinksTo checks if the provided file is a symlink which links to path
// It returs error if the link can't be read
func LinksTo(f *resource.File, path string) error {
	dst, err := os.Readlink(path)
	if err != nil {
		return err
	}
	if dst == f.Path {
		return nil
	}
	return fmt.Errorf("%s does not link to %s", f, path)
}

// Md5 checks if the provided file md5 checksum is the same as the one passed in as paramter
// It returs error if the provided file can't be open
func Md5(f *resource.File, sum string) error {
	return withOsFile(f.Path, func(file *os.File) error {
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
func Sha256(f *resource.File, sum string) error {
	return withOsFile(f.Path, func(file *os.File) error {
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
func Size(f *resource.File, size int64) error {
	return withFileInfo(f.Path, func(fi os.FileInfo) error {
		if fi.Size() == size {
			return nil
		}
		return fmt.Errorf("%s size is different from %d", f, size)
	})
}

// ModTimeAfter checks if the provided file modification time is older than the one passed in as paramter
// It returns error if os.Stat returns error
func ModTimeAfter(f *resource.File, mtime time.Time) error {
	return withFileInfo(f.Path, func(fi os.FileInfo) error {
		if fi.ModTime().After(mtime) {
			return nil
		}
		return fmt.Errorf("%s modification time is bigger than %s", f, mtime)
	})
}

// Contains checks if the provided file content can be matched with the regexp passed in as paramter
// It returs error if the provided file can't be open
func Contains(f *resource.File, content *regexp.Regexp) error {
	return withOsFile(f.Path, func(file *os.File) error {
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			if content.Match(scanner.Bytes()) {
				return nil
			}
		}
		return nil
	})
}
