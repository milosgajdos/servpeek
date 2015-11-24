package resource

import (
	"fmt"
	"os"
)

// Filer interface defines API interface to File resource
type Filer interface {
	// Path returns path to physical file on OS
	Path() string
	// Info returns info about underlying File or fails with error
	Info() (os.FileInfo, error)
	// String implements Stringer interface
	String() string
}

// File is an operating system file
type File struct {
	// Path to the physical file on filesystem
	path string
}

// NewFile rerturns pointer to File
func NewFile(path string) *File {
	return &File{path: path}
}

// Path returns path to physical path on OS
func (f *File) Path() string {
	return f.path
}

// Info returns os.FileInfo or error if os.Stat called on
// the underlying physical file fails with error
func (f *File) Info() (os.FileInfo, error) {
	fi, err := os.Stat(f.path)
	if err != nil {
		return nil, err
	}
	return fi, nil
}

// Implements Stringer interface
func (f *File) String() string {
	return fmt.Sprintf("[File] %s", f.path)
}
