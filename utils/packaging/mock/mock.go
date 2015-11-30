// Package mock provides various mock types useful for testing packaging packages
package mock

import "github.com/milosgajdos83/servpeek/utils/command"

type MockPkgCmd struct {
	cmdOut string
	args   string
}

func (mq *MockQueryPkgCmd) Run() command.Outer        { return MockOuter{cmdOut: mq.cmdOut} }
func (mq *MockQueryPkgCmd) AppendArgs(args ...string) { mq.args = append(mq.args, args) }

type MockOuter struct {
	cmdOut string
}

func (mo *MockOuter) Next() bool   {}
func (mo *MockOuter) Text() string {}
func (mo *MockOuter) Err() error   {}
func (mo *MockOuter) Close() error {}
