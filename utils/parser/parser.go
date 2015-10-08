// package parser impplements package manager
// command output parsers
package parser

import (
	"fmt"
	"regexp"

	"github.com/milosgajdos83/servpeek/resource/pkg"
	"github.com/milosgajdos83/servpeek/utils/commander"
)

// PkgParser parses Commander package commands
type Parser interface {
	ParseList(out *commander.Out) ([]*pkg.Pkg, error)
	ParseQuery(out *commander.Out) ([]*pkg.Pkg, error)
}

type Hinter interface {
	Filter() *regexp.Regexp
	Matcher() *regexp.Regexp
}

// NewPkgParser returns PkgParser
func NewPkgParser(pkgType string) (Parser, error) {
	switch pkgType {
	case "apt", "dpkg":
		return NewAptParser()
	case "rpm", "yum":
		return NewYumParser()
	case "pip":
		return NewPipParser()
	case "gem":
		return NewGemParser()
	}
	return nil, fmt.Errorf("Can't create Parser: Unsupported package type")
}
