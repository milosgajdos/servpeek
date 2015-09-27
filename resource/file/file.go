package file

import (
	"bufio"
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"os"
	"os/user"
	"regexp"
	"strconv"
	"syscall"
	"time"

	"github.com/milosgajdos83/servpeek/utils/group"
)

func withOsFileContext(path string, fn func(file *os.File) (bool, error)) (bool, error) {
	file, err := os.Open(path)
	if err != nil {
		return false, err
	}
	defer file.Close()
	return fn(file)
}

type File struct {
	fi os.FileInfo
	// fi.Name() returns base path :-/
	// so it's better to keep the ref to provided path
	path string
}

func NewFile(path string) (*File, error) {
	fi, err := os.Stat(path)
	if err != nil {
		return nil, err
	}

	return &File{
		fi:   fi,
		path: path,
	}, nil
}

func (f *File) String() string {
	return fmt.Sprintf("%s", f.path)
}

func (f *File) IsRegular() bool {
	return f.fi.Mode().IsRegular()
}

func (f *File) IsDirectory() bool {
	return f.fi.IsDir()
}

func (f *File) IsBlockDevice() bool {
	return f.fi.Mode()&os.ModeDevice != 0
}

func (f *File) IsCharDevice() bool {
	return f.fi.Mode()&os.ModeCharDevice != 0
}

func (f *File) IsPipe() bool {
	return f.fi.Mode()&os.ModeNamedPipe != 0
}

func (f *File) IsSocket() bool {
	return f.fi.Mode()&os.ModeSocket != 0
}

func (f *File) IsSymlink() bool {
	return f.fi.Mode()&os.ModeSymlink != 0
}

func (f *File) IsMode(mode os.FileMode) bool {
	return f.fi.Mode()&mode != 0
}

func (f *File) IsOwnedBy(username string) (bool, error) {
	u, err := user.Lookup(username)
	if err != nil {
		return false, err
	}
	uid, err := strconv.ParseUint(u.Uid, 10, 32)
	if err != nil {
		return false, err
	}
	return f.fi.Sys().(*syscall.Stat_t).Uid == uint32(uid), nil
}

func (f *File) IsGrupedInto(groupname string) (bool, error) {
	g, err := group.Lookup(groupname)
	if err != nil {
		return false, err
	}
	gid, err := strconv.ParseUint(g.Gid, 10, 32)
	if err != nil {
		return false, err
	}
	return f.fi.Sys().(*syscall.Stat_t).Gid == uint32(gid), nil
}

func (f *File) LinksTo(path string) (bool, error) {
	_, err := os.Readlink(path)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (f *File) Md5(sum string) (bool, error) {
	return withOsFileContext(f.path, func(file *os.File) (bool, error) {
		hasher := md5.New()
		if _, err := io.Copy(hasher, file); err != nil {
			return false, nil
		}
		return sum == hex.EncodeToString(hasher.Sum(nil)), nil
	})
}

func (f *File) Sha256(sum string) (bool, error) {
	return withOsFileContext(f.path, func(file *os.File) (bool, error) {
		hasher := sha256.New()
		if _, err := io.Copy(hasher, file); err != nil {
			return false, nil
		}
		return sum == hex.EncodeToString(hasher.Sum(nil)), nil
	})
}

func (f *File) Size(size int64) bool {
	return f.fi.Size() == size
}

func (f *File) ModTimeAfter(mtime time.Time) bool {
	return f.fi.ModTime().After(mtime)
}

func (f *File) Contains(content regexp.Regexp) (bool, error) {
	return withOsFileContext(f.fi.Name(), func(file *os.File) (bool, error) {
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			if content.Match(scanner.Bytes()) {
				return true, nil
			}
		}
		return false, nil
	})
}
