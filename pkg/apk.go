// build linux

package pkg

import (
	"regexp"

	"github.com/milosgajdos83/servpeek/utils/command"
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
	basePkgManager
}

// NewApkManager returns apk PkgManager or fails with error
func NewApkManager() (PkgManager, error) {
	return &apkManager{
		basePkgManager: basePkgManager{
			PkgCommander: NewApkCommander(),
			pkgType:      "apk",
		},
	}, nil
}

// apkCommander provides apk command manager commands
type apkCommander struct {
	basePkgCommander
}

// NewApkCommander returns PkgCommander that provides apk package manager commands
func NewApkCommander() PkgCommander {
	return &apkCommander{
		basePkgCommander: basePkgCommander{
			ListPkgsCmd: command.NewCommand(apk, apkListPkgsCmdArgs...),
			QueryPkgCmd: command.NewCommand(apk, apkQueryPkgCmdArgs...),
		},
	}
}

// apkParser provides parser for list and query package manager commands
type apkParser struct {
	basePkgOutParser
}

// NewApkParser returs PkgOutParser that parses apk PkgManager commands outputs
func NewApkParser() PkgOutParser {
	return &apkParser{
		basePkgOutParser: basePkgOutParser{
			ParseListOutFunc:  genParsePkgOutFunc("apk", "list", apkListPkgsOutHints),
			ParseQueryOutFunc: genParsePkgOutFunc("apk", "query", apkQueryPkgsOutHints),
		},
	}
}
