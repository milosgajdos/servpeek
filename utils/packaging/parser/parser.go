// Package parser implements package manager command output parsers
package parser

import (
	"fmt"

	"github.com/milosgajdos83/servpeek/resource"
	"github.com/milosgajdos83/servpeek/utils/command"
)

// PkgOutParser parses PkgCommander command output
type PkgOutParser interface {
	// ParseList parses output from ListPkgs command
	ParseList(out command.Outer) ([]*resource.Pkg, error)
	// ParseQuery parses output from QueryPkg command
	ParseQuery(out command.Outer) ([]*resource.Pkg, error)
}

// NewParser returns PkgOutParser
func NewParser(pkgType string) (PkgOutParser, error) {
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
