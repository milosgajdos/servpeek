package file

import (
	"bufio"
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
	"io"
	"os"
	"os/user"
	"regexp"
	"strconv"
	"syscall"
	"time"

	"github.com/milosgajdos83/servpeek/resource"
	"github.com/milosgajdos83/servpeek/utils/group"
)

func withOsFile(path string, fn func(file *os.File) (bool, error)) (bool, error) {
	file, err := os.Open(path)
	if err != nil {
		return false, err
	}
	defer file.Close()
	return fn(file)
}

func withFileInfo(path string, fn func(fi os.FileInfo) (bool, error)) (bool, error) {
	fi, err := os.Stat(path)
	if err != nil {
		return false, err
	}
	return fn(fi)
}

func IsRegular(f *resource.File) (bool, error) {
	return withFileInfo(f.Path, func(fi os.FileInfo) (bool, error) {
		return fi.Mode().IsRegular(), nil
	})
}

func IsDirectory(f *resource.File) (bool, error) {
	return withFileInfo(f.Path, func(fi os.FileInfo) (bool, error) {
		return fi.IsDir(), nil
	})
}

func IsBlockDevice(f *resource.File) (bool, error) {
	return withFileInfo(f.Path, func(fi os.FileInfo) (bool, error) {
		return fi.Mode()&os.ModeDevice != 0, nil
	})
}

func IsCharDevice(f *resource.File) (bool, error) {
	return withFileInfo(f.Path, func(fi os.FileInfo) (bool, error) {
		return fi.Mode()&os.ModeCharDevice != 0, nil
	})
}

func IsPipe(f *resource.File) (bool, error) {
	return withFileInfo(f.Path, func(fi os.FileInfo) (bool, error) {
		return fi.Mode()&os.ModeNamedPipe != 0, nil
	})
}

func IsSocket(f *resource.File) (bool, error) {
	return withFileInfo(f.Path, func(fi os.FileInfo) (bool, error) {
		return fi.Mode()&os.ModeSocket != 0, nil
	})
}

func IsSymlink(f *resource.File) (bool, error) {
	return withFileInfo(f.Path, func(fi os.FileInfo) (bool, error) {
		return fi.Mode()&os.ModeSymlink != 0, nil
	})
}

func IsMode(f *resource.File, mode os.FileMode) (bool, error) {
	return withFileInfo(f.Path, func(fi os.FileInfo) (bool, error) {
		return fi.Mode()&mode != 0, nil
	})
}

func IsOwnedBy(f *resource.File, username string) (bool, error) {
	return withFileInfo(f.Path, func(fi os.FileInfo) (bool, error) {
		u, err := user.Lookup(username)
		if err != nil {
			return false, err
		}
		uid, err := strconv.ParseUint(u.Uid, 10, 32)
		if err != nil {
			return false, err
		}
		return fi.Sys().(*syscall.Stat_t).Uid == uint32(uid), nil
	})
}

func IsGrupedInto(f *resource.File, groupname string) (bool, error) {
	return withFileInfo(f.Path, func(fi os.FileInfo) (bool, error) {
		g, err := group.Lookup(groupname)
		if err != nil {
			return false, err
		}
		gid, err := strconv.ParseUint(g.Gid, 10, 32)
		if err != nil {
			return false, err
		}
		return fi.Sys().(*syscall.Stat_t).Gid == uint32(gid), nil
	})
}

func LinksTo(f *resource.File, path string) (bool, error) {
	dst, err := os.Readlink(path)
	if err != nil {
		return false, err
	}
	return dst == f.Path, nil
}

func Md5(f *resource.File, sum string) (bool, error) {
	return withOsFile(f.Path, func(file *os.File) (bool, error) {
		hasher := md5.New()
		if _, err := io.Copy(hasher, file); err != nil {
			return false, nil
		}
		return sum == hex.EncodeToString(hasher.Sum(nil)), nil
	})
}

func Sha256(f *resource.File, sum string) (bool, error) {
	return withOsFile(f.Path, func(file *os.File) (bool, error) {
		hasher := sha256.New()
		if _, err := io.Copy(hasher, file); err != nil {
			return false, nil
		}
		return sum == hex.EncodeToString(hasher.Sum(nil)), nil
	})
}

func Size(f *resource.File, size int64) (bool, error) {
	return withFileInfo(f.Path, func(fi os.FileInfo) (bool, error) {
		return fi.Size() == size, nil
	})
}

func ModTimeAfter(f *resource.File, mtime time.Time) (bool, error) {
	return withFileInfo(f.Path, func(fi os.FileInfo) (bool, error) {
		return fi.ModTime().After(mtime), nil
	})
}

func Contains(f *resource.File, content regexp.Regexp) (bool, error) {
	return withOsFile(f.Path, func(file *os.File) (bool, error) {
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			if content.Match(scanner.Bytes()) {
				return true, nil
			}
		}
		return false, nil
	})
}
