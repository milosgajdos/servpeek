package apt

import "github.com/milosgajdos83/servpeek/utils"

const (
	QueryCmd = "dpkg-query"
)

var (
	// cli arguments passed to dpkg-query
	ListPkgsArgs  = []string{"-l"}
	QueryPkgsArgs = []string{"-W", "-f '${Status} ${Version}'"}
	// apt parseHints
	ParseHints = &utils.ParseHints{
		ListPrefix:      "ii",
		ListMinFields:   3,
		ListVersionIdx:  2,
		QueryPrefix:     "",
		QueryMinFields:  4,
		QueryVersionIdx: 3,
	}
)
