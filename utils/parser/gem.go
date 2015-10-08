package parser

import (
	"regexp"

	"github.com/milosgajdos83/servpeek/resource/pkg"
	"github.com/milosgajdos83/servpeek/utils/commander"
)

type GemParser struct {
	ph *parseHints
}

func NewGemParser() Parser {
	return &GemParser{
		ph: &parseHints{
			listFilter:  regexp.MustCompile(`^[A-Za-z]`),
			listMatch:   regexp.MustCompile(`^(\S+)\s+\((\S+)\)$`),
			queryFilter: regexp.MustCompile(`^[A-Za-z]`),
			queryMatch:  regexp.MustCompile(`^\S+\s+\((\S+)\)$`),
		},
	}
}

func (gp *GemParser) ParseList(out *commander.Out) ([]*pkg.Pkg, error) {
	return parseStream(out, gemListFilterRE, gemListMatchRE, parseListOut)
}

func (gp *GemParser) ParseQuery(out *commander.Out) ([]*pkg.Pkg, error) {
	return parseStream(out, gemQueryFilterRE, gemQueryMatchRE, parseQueryOut)
}
