package parser

import (
	"regexp"

	"github.com/milosgajdos83/servpeek/resource/pkg"
	"github.com/milosgajdos83/servpeek/utils/commander"
)

type PipParser struct {
	ph *parseHints
}

func NewPipParser() Parser {
	return &PipParser{
		ph: &parseHints{
			listFilter:  regexp.MustCompile(`^[A-Za-z]`),
			listMatch:   regexp.MustCompile(`^(\S+)\s+\((\S+)\)$`),
			queryFilter: regexp.MustCompile(`^Version`),
			queryMatch:  regexp.MustCompile(`^Version\s?:\s+(\S+).*`),
		},
	}
}

func (pp *PipParser) ParseList(out *commander.Out) ([]*pkg.Pkg, error) {
	return parseStream(out, pipListFilterRE, pipListMatchRE, parseListOut)
}

func (pp *PipParser) ParseQuery(out *commander.Out) ([]*pkg.Pkg, error) {
	return parseStream(out, pipQueryFilterRE, pipQueryMatchRE, parseQueryOut)
}
