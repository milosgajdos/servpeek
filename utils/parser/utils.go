package parser

import (
	"fmt"
	"regexp"

	"github.com/milosgajdos83/servpeek/resource/pkg"
	"github.com/milosgajdos83/servpeek/utils/commander"
)

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

type parseFunc func(string, *regexp.Regexp) ([]*pkg.Pkg, error)

func parseStream(out *commander.Out, fn parseFunc,
	h Hinter, pkgType string) ([]*pkg.Pkg, error) {
	pkgs := make([]*pkg.Pkg, 0)
	for out.Next() {
		line := out.Text()
		if h.Filter().MatchString(line) {
			pkg, err := fn(line, h.Matcher())
			if err != nil {
				return nil, err
			}
			if pkg != nil {
				pkg.Type = pkgType
				pkgs = append(pkgs, pkg)
			}
		}
	}
	return pkgs, nil
}

func parseListOut(line string, re *regexp.Regexp) (*pkg.Pkg, error) {
	match := re.FindStringSubmatch(line)
	if match == nil || len(match) < 3 {
		return nil, fmt.Errorf("Unable to parse package info")
	}
	return &pkg.Pkg{
		Version: match[2],
		Name:    match[1],
	}, nil
}

func parseQueryOut(line string, re *regexp.Regexp) (*pkg.Pkg, error) {
	match := re.FindStringSubmatch(line)
	if match == nil || len(match) < 2 {
		return nil, fmt.Errorf("Unable to parse package info")
	}
	return &pkg.Pkg{
		Version: match[1],
	}, nil
}
