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
	ParseHints    = &utils.ParseHints{}
)

func init() {
	ParseHints.ListFilter = regexp.MustCompile("")
	ParseHints.ListMatch = regexp.MustCompile("")
	ParseHints.QueryFilter = regexp.MustCompile("")
	ParseHints.QueryMatch = regexp.MustCompile("")
}
