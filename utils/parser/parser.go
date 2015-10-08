// package parser impplements package manager
// command output parsers
package parser

import (
	"fmt"

	"github.com/milosgajdos83/servpeek/resource"
	"github.com/milosgajdos83/servpeek/utils/commander"
)

// PkgParser parses Commander package commands
type Parser interface {
	ParseList(out *commander.Out) ([]*resource.Pkg, error)
	ParseQuery(out *commander.Out) ([]*resource.Pkg, error)
}

// NewPkgParser returns PkgParser
func NewParser(pkgType string) (Parser, error) {
	switch pkgType {
	case "apt", "dpkg":
		return NewAptParser(), nil
	case "rpm", "yum":
		return NewYumParser(), nil
	case "pip":
		return NewPipParser(), nil
	case "gem":
		return NewGemParser(), nil
	}
	return nil, fmt.Errorf("Can't create Parser: Unsupported package type")
}
