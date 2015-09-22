package gem

import (
	"regexp"

	"github.com/milosgajdos83/servpeek/utils"
)

const (
	QueryCmd = "gem"
)

var (
	ListPkgsArgs  = []string{"list", "--local"}
	QueryPkgsArgs = []string{"list", "--local"}
	ParseHints    = &utils.ParseHints{
		ListFilter:  regexp.MustCompile(`^[A-Za-z]`),
		ListMatch:   regexp.MustCompile(`^(?P<name>\S+)\s+\((?P<version>\S+)\)$`),
		QueryFilter: regexp.MustCompile(`^[A-Za-z]`),
		QueryMatch:  regexp.MustCompile(`^\S+\s+\((?P<version>\S+)\)$`),
	}
)
