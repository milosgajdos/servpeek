package pkg

import (
	"bufio"
	"strings"

	"github.com/milosgajdos83/servpeek/utils/command"
)

// mockPkgCmdOut implements command.Output interface
type mockPkgCmdOut struct {
	scanner *bufio.Scanner
}

func (m *mockPkgCmdOut) Next() bool   { return m.scanner.Scan() }
func (m *mockPkgCmdOut) Text() string { return m.scanner.Text() }
func (m *mockPkgCmdOut) Err() error   { return m.scanner.Err() }
func (m *mockPkgCmdOut) Close() error { return nil }

// mockPkgCommand implements command.Command interface
type mockPkgCommand struct {
	cmdOut string
}

func (m *mockPkgCommand) Run() command.Output {
	return &mockPkgCmdOut{scanner: bufio.NewScanner(strings.NewReader(m.cmdOut))}
}

func (m *mockPkgCommand) RunCombined() (string, error) { return "", nil }
func (m *mockPkgCommand) AppendArgs(args ...string)    {}

// mockPkgManager implements pkg.Manager interface
type mockPkgManager struct {
	listCmd  *mockPkgCommand
	queryCmd *mockPkgCommand
	parser   CmdOutParser
	pkgType  string
}

func (m *mockPkgManager) Type() string { return m.pkgType }
func (m *mockPkgManager) ListPkgs() ([]Pkg, error) {
	return m.parser.ParseListPkgsOut(m.listCmd.Run())
}
func (m *mockPkgManager) QueryPkg(pkgName string) ([]Pkg, error) {
	return m.parser.ParseQueryPkgOut(m.queryCmd.Run())
}

type mockPkg struct {
	manager *mockPkgManager
	name    string
	version string
}

func (m *mockPkg) Manager() Manager { return m.manager }
func (m *mockPkg) Name() string     { return m.name }
func (m *mockPkg) Version() string  { return m.version }
