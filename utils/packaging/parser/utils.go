package parser

import (
	"fmt"
	"regexp"

	"github.com/milosgajdos83/servpeek/resource"
	"github.com/milosgajdos83/servpeek/utils/command"
)

type hinter interface {
	Filter() *regexp.Regexp
	Matcher() *regexp.Regexp
}

type baseHinter struct {
	list  hinter
	query hinter
}

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

type parseFunc func(string, *regexp.Regexp) (*resource.Pkg, error)

func parseStream(out *command.Out, fn parseFunc,
	h hinter, pkgType string) ([]*resource.Pkg, error) {
	pkgs := make([]*resource.Pkg, 0)
	for out.Next() {
		line := out.Text()
		if h.Filter().MatchString(line) {
			p, err := fn(line, h.Matcher())
			if err != nil {
				return nil, err
			}
			if p != nil {
				p.Type = pkgType
				pkgs = append(pkgs, p)
			}
		}
	}
	return pkgs, nil
}

func parseListOut(line string, re *regexp.Regexp) (*resource.Pkg, error) {
	match := re.FindStringSubmatch(line)
	if match == nil || len(match) < 3 {
		return nil, fmt.Errorf("Unable to parse package info")
	}
	return &resource.Pkg{
		Version: match[2],
		Name:    match[1],
	}, nil
}

func parseQueryOut(line string, re *regexp.Regexp) (*resource.Pkg, error) {
	match := re.FindStringSubmatch(line)
	if match == nil || len(match) < 2 {
		return nil, fmt.Errorf("Unable to parse package info")
	}
	return &resource.Pkg{
		Version: match[1],
	}, nil
}
