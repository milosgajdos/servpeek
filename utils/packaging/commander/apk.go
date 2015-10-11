package commander

const (
	apk = "apk"
)

var (
	// cli arguments passed to dpkg-query
	apkListPkgsArgs = []string{"info", "-v"}
	apkQueryPkgArgs = []string{"info"}
)

// ApkCommander provides apk command manager commands
type ApkCommander struct {
	*PkgCommander
}

// NewApkCommander returns aptitude command manager
func NewApkCommander() *PkgCommander {
	return &PkgCommander{
		ListPkgs: BuildCmd(apk, apkListPkgsArgs...),
		QueryPkg: BuildCmd(apk, apkQueryPkgArgs...),
	}
}
