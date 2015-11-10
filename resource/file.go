package resource

import (
	"fmt"
	"os"
)

type Filer interface {
	Path() string
	String() string
	Info() (os.FileInfo, error)
}

// File is an operating system file
type File struct {
	// Path to the physical file
	path string
}

func NewFile(path string) *File {
	return &File{path: path}
}

func (f *File) Path() string {
	return f.path
}

// Implement Stringer interface
func (f *File) String() string {
	return fmt.Sprintf("[File] %s", f.path)
}

func (f *File) Info() (os.FileInfo, error) {
	fi, err := os.Stat(f.path)
	if err != nil {
		return nil, err
	}
	return fi, nil
}
