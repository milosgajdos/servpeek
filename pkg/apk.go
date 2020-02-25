// build linux

package pkg

import (
	"regexp"

	"github.com/milosgajdos/servpeek/utils/command"
)

const apk = "apk"

var (
	// cli arguments passed to dpkg-query
	apkListPkgsCmdArgs = []string{"info", "-v"}
	apkQueryPkgCmdArgs = []string{"info"}
	// apk package manager parser hints
	apkListPkgsOutHints = &hints{
		filter:  regexp.MustCompile(`^[^W].*`),
		matcher: regexp.MustCompile(`^(\S+)-(\d\S+)$`),
	}
	apkQueryPkgsOutHints = &hints{
		filter:  regexp.MustCompile(`.*description:$`),
		matcher: regexp.MustCompile(`^\S+-(\d\S+)\s+description:$`),
	}
)

// apkManager implements apk package manager
type apkManager struct {
	baseManager
}

// NewApkManager returns apk Manager or fails with error
func NewApkManager() Manager {
	return &apkManager{
		baseManager: baseManager{
			ListPkgsCmd: command.NewCommand(apk, apkListPkgsCmdArgs...),
			QueryPkgCmd: command.NewCommand(apk, apkQueryPkgCmdArgs...),
			parser:      NewApkParser(),
			pkgType:     "apk",
		},
	}
}

// apkParser provides parser for list and query package manager commands
type apkParser struct {
	baseCmdOutParser
}

// NewApkParser returs CmdOutParser that parses apk Manager commands outputs
func NewApkParser() CmdOutParser {
	return &apkParser{
		baseCmdOutParser: baseCmdOutParser{
			ParseListOutFunc:  genParsePkgOutFunc("apk", "list", apkListPkgsOutHints),
			ParseQueryOutFunc: genParsePkgOutFunc("apk", "query", apkQueryPkgsOutHints),
		},
	}
}
