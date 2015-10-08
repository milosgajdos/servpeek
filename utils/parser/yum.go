package parser

import (
	"regexp"

	"github.com/milosgajdos83/servpeek/resource/pkg"
	"github.com/milosgajdos83/servpeek/utils/commander"
)

type YumParser struct {
	ph *parseHints
}

func NewYumParser() Parser {
	return &YumParser{
		ph: &parseHints{
			listFilter:  regexp.MustCompile(`^[A-Za-z]`),
			listMatch:   regexp.MustCompile(`^(\S+)\s+(\S+).*`),
			queryFilter: regexp.MustCompile(`^Version`),
			queryMatch:  regexp.MustCompile(`^Version\s+:\s+(\S+).*`),
		},
	}
}

func (yp *YumParser) ParseList(out *commander.Out) ([]*pkg.Pkg, error) {
	return parseStream(out, yumListFilterRE, yumListMatchRE, parseListOut)
}

func (yp *YumParser) ParseQuery(out *commander.Out) ([]*pkg.Pkg, error) {
	return parseStream(out, yumQueryFilterRE, yumQueryMatchRE, parseQueryOut)
}
