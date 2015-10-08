package parser

import (
	"regexp"

	"github.com/milosgajdos83/servpeek/resource"
	"github.com/milosgajdos83/servpeek/utils/commander"
)

type aptParser struct {
	hinter *baseHinter
}

func NewAptParser() Parser {
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

func (ap *aptParser) ParseList(out *commander.Out) ([]*resource.Pkg, error) {
	return parseStream(out, parseListOut, ap.hinter.list, "apt")
}

func (ap *aptParser) ParseQuery(out *commander.Out) ([]*resource.Pkg, error) {
	return parseStream(out, parseQueryOut, ap.hinter.query, "apt")
}
