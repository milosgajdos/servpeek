package resource

import "fmt"

// File is an operating system file
type File struct {
	// Path to the physical file
	Path string
}

// Implement Stringer interface
func (f *File) String() string {
	return fmt.Sprintf("[File] %s", f.Path)
}
