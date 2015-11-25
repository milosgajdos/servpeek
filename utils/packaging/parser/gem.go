package parser

import (
	"regexp"

	"github.com/milosgajdos83/servpeek/resource"
	"github.com/milosgajdos83/servpeek/utils/command"
)

type gemParser struct {
	hinter *baseHinter
}

// NewGemParser returns PkgOutParser that parses gem PkgManager commands outputs
func NewGemParser() PkgOutParser {
	return &gemParser{
		hinter: &baseHinter{
			list: &hints{
				filter:  regexp.MustCompile(`^[A-Za-z]`),
				matcher: regexp.MustCompile(`^(\S+)\s+\((\S+)\)$`),
			},
			query: &hints{
				filter:  regexp.MustCompile(`^[A-Za-z]`),
				matcher: regexp.MustCompile(`^\S+\s+\((\S+)\)$`),
			},
		},
	}
}

// ParseList parses output of "gem list --local" command
// It returns slice of installed packages or error
func (gp *gemParser) ParseList(out command.Outer) ([]*resource.Pkg, error) {
	return parseStream(out, parseListOut, gp.hinter.list, "gem")
}

// ParseQuery parses output of "gem list --local" command
// It returns slice of queried packages or error
func (gp *gemParser) ParseQuery(out command.Outer) ([]*resource.Pkg, error) {
	return parseStream(out, parseQueryOut, gp.hinter.query, "gem")
}
