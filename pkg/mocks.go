package pkg

import (
	"bufio"
	"strings"

	"github.com/milosgajdos83/servpeek/utils/command"
)

type mockOuter struct {
	scanner *bufio.Scanner
}

func (m *mockOuter) Next() bool   { return m.scanner.Scan() }
func (m *mockOuter) Text() string { return m.scanner.Text() }
func (m *mockOuter) Err() error   { return m.scanner.Err() }
func (m *mockOuter) Close() error { return nil }

type mockCommander struct {
	cmdOut string
}

func (m *mockCommander) Run() command.Outer {
	return &mockOuter{scanner: bufio.NewScanner(strings.NewReader(m.cmdOut))}
}

func (m *mockCommander) AppendArgs(args ...string) {}

type mockManager struct {
	listCmd  *mockCommander
	queryCmd *mockCommander
	parser   CmdOutParser
	pkgType  string
}

func (m *mockManager) Type() string { return m.pkgType }
func (m *mockManager) ListPkgs() ([]Pkg, error) {
	return m.parser.ParseListPkgsOut(m.listCmd.Run())
}
func (m *mockManager) QueryPkg(pkgName string) ([]Pkg, error) {
	return m.parser.ParseQueryPkgOut(m.queryCmd.Run())
}

type mockPkg struct {
	manager *mockManager
	name    string
	version string
}

func (m *mockPkg) Manager() Manager { return m.manager }
func (m *mockPkg) Name() string     { return m.name }
func (m *mockPkg) Version() string  { return m.version }
