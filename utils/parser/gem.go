package parser

import (
	"regexp"

	"github.com/milosgajdos83/servpeek/resource"
	"github.com/milosgajdos83/servpeek/utils/commander"
)

type gemParser struct {
	hinter *baseHinter
}

func NewGemParser() Parser {
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

func (gp *gemParser) ParseList(out *commander.Out) ([]*resource.Pkg, error) {
	return parseStream(out, parseListOut, gp.hinter.list, "gem")
}

func (gp *gemParser) ParseQuery(out *commander.Out) ([]*resource.Pkg, error) {
	return parseStream(out, parseQueryOut, gp.hinter.query, "gem")
}
