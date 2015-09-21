package apt

import (
	"regexp"

	"github.com/milosgajdos83/servpeek/utils"
)

const (
	QueryCmd = "dpkg-query"
)

var (
	// cli arguments passed to dpkg-query
	ListPkgsArgs  = []string{"-l"}
	QueryPkgsArgs = []string{"-W", "-f '${Status} ${Version}'"}
	// apt parseHints
	ParseHints = &utils.ParseHints{
		ListFilter:  regexp.MustCompile(`^ii`),
		ListMatch:   regexp.MustCompile(`^ii\s+(?P<name>\w+)\s+(?P<version>\S+)\s+.*`),
		QueryFilter: regexp.MustCompile(`^[A-Za-z]`),
		QueryMatch:  regexp.MustCompile(`^\w+\s+\w+\s+\w+\s+(\S+)$`),
	}
)
