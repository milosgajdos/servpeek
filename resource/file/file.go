package file

import (
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
	"io"
	"os"
	"regexp"
	"syscall"
	"time"

	"github.com/milosgajdos83/servpeek/utils"
)

type File struct {
	fi os.FileInfo
}

func NewFile(path string) (*File, error) {
	fi, err := os.Stat(path)
	if err != nil {
		return nil, err
	}

	return &File{
		fi: fi,
	}, nil
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

func (f *File) Contains(content regexp.Regexp) bool {
	return false
}

func (f *File) IsMode(mode os.FileMode) bool {
	return f.fi.Mode()&mode != 0
}

func (f *File) IsOwnedBy(user string) (bool, error) {
	uid, err := utils.Uid(user)
	if err != nil {
		return false, err
	}
	return f.fi.Sys().(*syscall.Stat_t).Uid == uid, nil
}

// FIXME: Golang can't do groupname lookups :-/
// https://github.com/golang/go/issues/2617
func (f *File) IsGrupedInto(gid uint32) bool {
	return f.fi.Sys().(*syscall.Stat_t).Gid == gid
}

func (f *File) IsReadableBy(username string) bool {
	return false
}

func (f *File) IsWritableBy(username string) bool {
	return false
}

func (f *File) IsExecutableBy(username string) bool {
	return false
}

func (f *File) LinksTo(path string) (bool, error) {
	_, err := os.Readlink(path)
	if err != nil {
		return false, err
	}
	return true, nil
}

func withFileContext(path string, fn func(file *os.File) (bool, error)) (bool, error) {
	file, err := os.Open(path)
	if err != nil {
		return false, err
	}
	defer file.Close()
	return fn(file)
}

func (f *File) Md5(sum string) (bool, error) {
	return withFileContext(f.fi.Name(), func(file *os.File) (bool, error) {
		hasher := md5.New()
		if _, err := io.Copy(hasher, file); err != nil {
			return false, nil
		}
		return sum == hex.EncodeToString(hasher.Sum(nil)), nil
	})
}

func (f *File) Sha256(sum string) (bool, error) {
	return withFileContext(f.fi.Name(), func(file *os.File) (bool, error) {
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
