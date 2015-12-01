package service

type mockService struct {
	name    string
	sysinit SysInit
}

func (m *mockService) Name() string     { return m.name }
func (m *mockService) SysInit() SysInit { return m.sysinit }

type mockSysInit struct {
	start  error
	stop   error
	status Status
	err    error
}

func (m *mockSysInit) Type() string                          { return "mock" }
func (m *mockSysInit) Start(svcname string) error            { return m.start }
func (m *mockSysInit) Stop(svcName string) error             { return m.stop }
func (m *mockSysInit) Status(svcName string) (Status, error) { return m.status, m.err }
