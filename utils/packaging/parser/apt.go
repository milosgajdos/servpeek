// build linux

package parser

import (
	"regexp"

	"github.com/milosgajdos83/servpeek/resource"
	"github.com/milosgajdos83/servpeek/utils/command"
)

type aptParser struct {
	hinter *baseHinter
}

// NewAptParser returs PkgParser that parses aptitude PkgManager commands outputs
func NewAptParser() PkgParser {
	return &aptParser{
		hinter: &baseHinter{
			list: &hints{
				filter:  regexp.MustCompile(`^ii`),
				matcher: regexp.MustCompile(`^ii\s+(\S+)\s+(\S+)\s+.*`),
			},
			query: &hints{
				filter:  regexp.MustCompile(`^\s+'[A-Za-z]`),
				matcher: regexp.MustCompile(`^\s+'\w+\s+\w+\s+\w+\s+(\S+)'$`),
			},
		},
	}
}

// ParseList parses output of dpkg-query -l command
// It returns slice of installed packages or error
func (ap *aptParser) ParseList(out *command.Out) ([]*resource.Pkg, error) {
	return parseStream(out, parseListOut, ap.hinter.list, "apt")
}

// ParseQuery parses output of dpkg-query -W -f '${Status} ${Version} command
// It returns slice of packages or error
func (ap *aptParser) ParseQuery(out *command.Out) ([]*resource.Pkg, error) {
	return parseStream(out, parseQueryOut, ap.hinter.query, "apt")
}
