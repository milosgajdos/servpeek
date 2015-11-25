// build linux

package parser

import (
	"regexp"

	"github.com/milosgajdos83/servpeek/resource"
	"github.com/milosgajdos83/servpeek/utils/command"
)

type apkParser struct {
	hinter *baseHinter
}

// NewApkParser returs PkgOutParser that parses apk PkgManager commands outputs
func NewApkParser() PkgOutParser {
	return &apkParser{
		hinter: &baseHinter{
			list: &hints{
				filter:  regexp.MustCompile(`^[^W].*`),
				matcher: regexp.MustCompile(`^(\S+)-(\d\S+)$`),
			},
			query: &hints{
				filter:  regexp.MustCompile(`.*description:$`),
				matcher: regexp.MustCompile(`^\S+-(\d\S+)\s+description:$`),
			},
		},
	}
}

// ParseList parses output of apk info -v command
// It returns slice of installed packages or error
func (ap *apkParser) ParseList(out command.Outer) ([]*resource.Pkg, error) {
	return parseStream(out, parseListOut, ap.hinter.list, "apk")
}

// ParseQuery parses output of apk info pkg_name command
// It returns slice of packages or error
func (ap *apkParser) ParseQuery(out command.Outer) ([]*resource.Pkg, error) {
	return parseStream(out, parseQueryOut, ap.hinter.query, "apk")
}
