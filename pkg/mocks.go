package pkg

type MockPkg struct {
	manager *MockPkgManager
	name    string
	version string
}

func (m *MockPkg) Manager() PkgManager { return m.manager }
func (m *MockPkg) Name() string        { return m.name }
func (m *MockPkg) Version() string     { return m.version }

type MockPkgManager struct {
	pkgType   string
	listPkgs  []Pkg
	queryPkgs []Pkg
}

func (m *MockPkgManager) Type() string                           { return m.pkgType }
func (m *MockPkgManager) ListPkgs() ([]Pkg, error)               { return m.listPkgs, nil }
func (m *MockPkgManager) QueryPkg(pkgName string) ([]Pkg, error) { return m.queryPkgs, nil }
