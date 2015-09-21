package pip

import (
	"regexp"

	"github.com/milosgajdos83/servpeek/utils"
)

const (
	QueryCmd = "pip"
)

var (
	ListPkgsArgs  = []string{"list"}
	QueryPkgsArgs = []string{"show"}
	ParseHints    = &utils.ParseHints{
		ListFilter:  regexp.MustCompile(`^[A-Za-z]`),
		ListMatch:   regexp.MustCompile(`^(?P<name>\S+)\s+\((?P<version>\S+)\)$`),
		QueryFilter: regexp.MustCompile(`^Version`),
		QueryMatch:  regexp.MustCompile(`^Version\s?:\s+(?P<version>\S+).*`),
	}
)
