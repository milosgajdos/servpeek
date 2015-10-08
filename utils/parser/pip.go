package parser

import (
	"regexp"

	"github.com/milosgajdos83/servpeek/resource"
	"github.com/milosgajdos83/servpeek/utils/commander"
)

type pipParser struct {
	hinter *baseHinter
}

func NewPipParser() Parser {
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

func (pp *pipParser) ParseList(out *commander.Out) ([]*resource.Pkg, error) {
	return parseStream(out, parseListOut, pp.hinter.list, "pip")
}

func (pp *pipParser) ParseQuery(out *commander.Out) ([]*resource.Pkg, error) {
	return parseStream(out, parseQueryOut, pp.hinter.query, "pip")
}
