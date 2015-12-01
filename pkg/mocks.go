package pkg

type mockPkg struct {
	manager *mockManager
	name    string
	version string
}

func (m *mockPkg) Manager() Manager { return m.manager }
func (m *mockPkg) Name() string     { return m.name }
func (m *mockPkg) Version() string  { return m.version }

type mockManager struct {
	pkgType   string
	listPkgs  []Pkg
	queryPkgs []Pkg
}

func (m *mockManager) Type() string                           { return m.pkgType }
func (m *mockManager) ListPkgs() ([]Pkg, error)               { return m.listPkgs, nil }
func (m *mockManager) QueryPkg(pkgName string) ([]Pkg, error) { return m.queryPkgs, nil }
