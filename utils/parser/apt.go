package parser

import (
	"regexp"

	"github.com/milosgajdos83/servpeek/resource/pkg"
	"github.com/milosgajdos83/servpeek/utils/commander"
)

type AptParser struct {
	listHinter  Hinter
	queryHinter Hinter
}

func NewAptParser() Parser {
	return &AptParser{
		listHinter: &hints{
			listFilter: regexp.MustCompile(`^ii`),
			listMatch:  regexp.MustCompile(`^ii\s+(\S+)\s+(\S+)\s+.*`),
		},
		queryHinter: &hints{
			queryFilter: regexp.MustCompile(`^\s+'[A-Za-z]`),
			queryMatch:  regexp.MustCompile(`^\s+'\w+\s+\w+\s+\w+\s+(\S+)'$`),
		},
	}
}

func (ap *AptParser) ParseList(out *commander.Out) ([]*pkg.Pkg, error) {
	return parseStream(out, parseListOut, ap.listHinter, "apt")
}

func (ap *AptParser) ParseQuery(out *commander.Out) ([]*pkg.Pkg, error) {
	return parseStream(out, parseQueryOut, ap.queryHinter, "apt")
}
