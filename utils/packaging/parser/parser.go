// Package parser impplements package manager command output parsers
package parser

import (
	"fmt"

	"github.com/milosgajdos83/servpeek/resource"
	"github.com/milosgajdos83/servpeek/utils/command"
)

// PkgParser parses PkgCommander command output
type PkgParser interface {
	// ParseList parses output from ListPkgs command
	ParseList(out *command.Out) ([]*resource.Pkg, error)
	// ParseQuery parses output from QueryPkg command
	ParseQuery(out *command.Out) ([]*resource.Pkg, error)
}

// NewParser returns PkgParser
func NewParser(pkgType string) (PkgParser, error) {
	switch pkgType {
	case "apt", "dpkg":
		return NewAptParser(), nil
	case "rpm", "yum":
		return NewYumParser(), nil
	case "apk":
		return NewApkParser(), nil
	case "pip":
		return NewPipParser(), nil
	case "gem":
		return NewGemParser(), nil
	}
	return nil, fmt.Errorf("Unable to create PkgParser for %s: Unsupported package type", pkgType)
}
