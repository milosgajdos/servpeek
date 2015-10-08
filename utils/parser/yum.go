package parser

import (
	"regexp"

	"github.com/milosgajdos83/servpeek/resource"
	"github.com/milosgajdos83/servpeek/utils/commander"
)

type yumParser struct {
	hinter *baseHinter
}

func NewYumParser() Parser {
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

func (yp *yumParser) ParseList(out *commander.Out) ([]*resource.Pkg, error) {
	return parseStream(out, parseListOut, yp.hinter.list, "yum")
}

func (yp *yumParser) ParseQuery(out *commander.Out) ([]*resource.Pkg, error) {
	return parseStream(out, parseQueryOut, yp.hinter.query, "yum")
}
