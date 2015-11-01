// build linux

package parser

import (
	"regexp"

	"github.com/milosgajdos83/servpeek/resource"
	"github.com/milosgajdos83/servpeek/utils/command"
)

type yumParser struct {
	hinter *baseHinter
}

// NewYumParser returs PkgParser that parses yum PkgManager commands outputs
func NewYumParser() PkgParser {
	return &yumParser{
		hinter: &baseHinter{
			list: &hints{
				filter:  regexp.MustCompile(`^[A-Za-z]`),
				matcher: regexp.MustCompile(`^(\S+)\s+(\S+).*`),
			},
			query: &hints{
				filter:  regexp.MustCompile(`^Version`),
				matcher: regexp.MustCompile(`^Version\s+:\s+(\S+).*`),
			},
		},
	}
}

// ParseList parses output of rpm -qa --qf %{NAME}%20{VERSION}-%{RELEASE}
// It returns slice of installed packages or error
func (yp *yumParser) ParseList(out *command.Out) ([]*resource.Pkg, error) {
	return parseStream(out, parseListOut, yp.hinter.list, "yum")
}

// ParseQuery parses output of rpm -qi command
// It returns slice of packages or error
func (yp *yumParser) ParseQuery(out *command.Out) ([]*resource.Pkg, error) {
	return parseStream(out, parseQueryOut, yp.hinter.query, "yum")
}
