package pkg

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"path"
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
	cmd     *mockPkgCommand
	parser  CmdOutParser
	pkgType string
}

func newMockPkgManager(pkgType, cmdType string) (*mockPkgManager, error) {
	parsers := map[string]CmdOutParser{
		"apt": NewAptParser(),
		"yum": NewYumParser(),
		"apk": NewApkParser(),
		"pip": NewPipParser(),
		"gem": NewGemParser(),
	}

	currentDir, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	fileName := fmt.Sprintf("%s-%s.out", pkgType, cmdType)
	fixturesPath := path.Join(currentDir, "test-fixtures", fileName)
	cmdOut, err := ioutil.ReadFile(fixturesPath)
	if err != nil {
		return nil, err
	}

	return &mockPkgManager{
		cmd: &mockPkgCommand{
			cmdOut: string(cmdOut),
		},
		parser:  parsers[pkgType],
		pkgType: pkgType,
	}, nil
}

func (m *mockPkgManager) Type() string { return m.pkgType }
func (m *mockPkgManager) ListPkgs() ([]Pkg, error) {
	return m.parser.ParseListPkgsOut(m.cmd.Run())
}
func (m *mockPkgManager) QueryPkg(pkgName string) ([]Pkg, error) {
	return m.parser.ParseQueryPkgOut(m.cmd.Run())
}

type mockPkg struct {
	manager  *mockPkgManager
	name     string
	versions []string
}

func newMockPkg(pkgType, pkgName, cmdType string, pkgVersion ...string) (*mockPkg, error) {
	pkgManager, err := newMockPkgManager(pkgType, cmdType)
	if err != nil {
		return nil, err
	}

	return &mockPkg{
		manager:  pkgManager,
		name:     pkgName,
		versions: pkgVersion,
	}, nil
}

func (m *mockPkg) Manager() Manager   { return m.manager }
func (m *mockPkg) Name() string       { return m.name }
func (m *mockPkg) Versions() []string { return m.versions }
