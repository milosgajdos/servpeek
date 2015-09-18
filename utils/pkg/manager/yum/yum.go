package yum

import "github.com/milosgajdos83/servpeek/utils"

const (
	QueryCmd = "rpm"
)

var (
	// cli arguments passed to rpm
	ListPkgsArgs  = []string{"-qa --qf '%{NAME}%20{VERSION}-%{RELEASE}\n'"}
	QueryPkgsArgs = []string{"-qi"}
	// yum parse hints
	ParseHints = &utils.ParseHints{
		ListPrefix:      "",
		ListMinFields:   2,
		ListVersionIdx:  1,
		QueryPrefix:     "Version",
		QueryMinFields:  3,
		QueryVersionIdx: 2,
	}
)
