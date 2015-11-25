package parser

import (
	"regexp"

	"github.com/milosgajdos83/servpeek/resource"
	"github.com/milosgajdos83/servpeek/utils/command"
)

type pipParser struct {
	hinter *baseHinter
}

// NewPipParser returns PkgOutParser that parses pip PkgManager commands outputs
func NewPipParser() PkgOutParser {
	return &pipParser{
		hinter: &baseHinter{
			list: &hints{
				filter:  regexp.MustCompile(`^[A-Za-z]`),
				matcher: regexp.MustCompile(`^(\S+)\s+\((\S+)\)$`),
			},
			query: &hints{
				filter:  regexp.MustCompile(`^Version`),
				matcher: regexp.MustCompile(`^Version\s?:\s+(\S+).*`),
			},
		},
	}
}

// ParseList parses output of "pip list" command
// It returns slice of list packages or error
func (pp *pipParser) ParseList(out command.Outer) ([]*resource.Pkg, error) {
	return parseStream(out, parseListOut, pp.hinter.list, "pip")
}

// ParseQuery parses output of "pip show" command
// It returns slice of queried packages or error
func (pp *pipParser) ParseQuery(out command.Outer) ([]*resource.Pkg, error) {
	return parseStream(out, parseQueryOut, pp.hinter.query, "pip")
}
