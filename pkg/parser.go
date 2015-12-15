package pkg

import (
	"fmt"
	"regexp"

	"github.com/milosgajdos83/servpeek/utils/command"
)

// CmdOutParser provides an interface to parse output of package manager commands
type CmdOutParser interface {
	// ParseListOut parses parses output from list package command
	ParseListPkgsOut(command.Outer) ([]Pkg, error)
	// ParseQueryOut parses output from query package command
	ParseQueryPkgOut(command.Outer) ([]Pkg, error)
}

// ParsePkgOutFunc parses command output and returns a slice of packages.
// It returns error if the command output can not be parsed
type ParsePkgOutFunc func(command.Outer) ([]Pkg, error)

// NewCmdOutParser creates new package command output parser
func NewCmdOutParser(pkgMgrType string) (CmdOutParser, error) {
	switch pkgMgrType {
	case "apt":
		return NewAptParser(), nil
	case "yum":
		return NewYumParser(), nil
	case "apk":
		return NewApkParser(), nil
	case "pip":
		return NewPipParser(), nil
	case "gem":
		return NewGemParser(), nil
	}
	return nil, fmt.Errorf("Could not create CmdOutParser Unsupported package manager: %s", pkgMgrType)
}

type baseCmdOutParser struct {
	// ParseListOutFunc parses output of list package command
	ParseListOutFunc ParsePkgOutFunc
	// ParseQueryOutFunc parses output of query package command
	ParseQueryOutFunc ParsePkgOutFunc
}

// ParseListPkgsOut parses output from list packages command manager command
func (b *baseCmdOutParser) ParseListPkgsOut(out command.Outer) ([]Pkg, error) {
	return b.ParseListOutFunc(out)
}

// ParseQueryPkgOut parses output from query package command manager command
func (b *baseCmdOutParser) ParseQueryPkgOut(out command.Outer) ([]Pkg, error) {
	return b.ParseQueryOutFunc(out)
}

// GenParsePkgOutFunc generates function that can parse output of execute package manager command
// It returns error if unsupported package type is requested
func genParsePkgOutFunc(pkgType, cmdType string, h *hints) ParsePkgOutFunc {
	return func(out command.Outer) ([]Pkg, error) {
		switch cmdType {
		case "list":
			return parseStream(out, parseListOut, h, pkgType)
		case "query":
			return parseStream(out, parseQueryOut, h, pkgType)
		}
		return nil, fmt.Errorf("Unsupported parse command type: %s", cmdType)
	}
}

type lineParseFunc func(string, string, *regexp.Regexp) (Pkg, error)

type hints struct {
	filter  *regexp.Regexp
	matcher *regexp.Regexp
}

func (h *hints) Filter() *regexp.Regexp {
	return h.filter
}

func (h *hints) Matcher() *regexp.Regexp {
	return h.matcher
}

func parseStream(out command.Outer, fn lineParseFunc, h *hints, pkgType string) ([]Pkg, error) {
	var pkgs []Pkg
	for out.Next() {
		line := out.Text()
		if h.Filter().MatchString(line) {
			p, err := fn(pkgType, line, h.Matcher())
			if err != nil {
				return nil, err
			}
			pkgs = append(pkgs, p)
		}
	}
	return pkgs, out.Err()
}

func parseListOut(pkgType, line string, re *regexp.Regexp) (Pkg, error) {
	match := re.FindStringSubmatch(line)
	if match == nil || len(match) < 3 {
		return nil, fmt.Errorf("Unable to parse List package info")
	}
	return NewSwPkg(pkgType, match[1], match[2])
}

func parseQueryOut(pkgType, line string, re *regexp.Regexp) (Pkg, error) {
	match := re.FindStringSubmatch(line)
	if match == nil || len(match) < 2 {
		return nil, fmt.Errorf("Unable to parse Query package info")
	}
	return NewSwPkg(pkgType, "foo", match[1])
}
