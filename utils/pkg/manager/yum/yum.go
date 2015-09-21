package yum

import (
	"regexp"

	"github.com/milosgajdos83/servpeek/utils"
)

const (
	QueryCmd = "rpm"
)

var (
	// cli arguments passed to rpm
	ListPkgsArgs  = []string{"-qa --qf '%{NAME}%20{VERSION}-%{RELEASE}\n'"}
	QueryPkgsArgs = []string{"-qi"}
	// yum parse hints
	ParseHints = &utils.ParseHints{
		ListFilter:  regexp.MustCompile(`^[A-Za-z]`),
		ListMatch:   regexp.MustCompile(`^(?P<name>\S+)\s+(?P<version>\S+).*`),
		QueryFilter: regexp.MustCompile(`^Version`),
		QueryMatch:  regexp.MustCompile(`^Version\s+:\s+(?P<version>\S+).*`),
	}
)
