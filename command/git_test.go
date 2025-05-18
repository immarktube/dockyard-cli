package command

import (
	"github.com/immarktube/dockyard-cli/config"
	"testing"

	"github.com/stretchr/testify/assert"
)

type MockExecutor struct {
	CalledWith []string
}

func (m *MockExecutor) RunCommand(dir, name string, args ...string) (string, error) {
	m.CalledWith = append(m.CalledWith, name+" "+args[0])
	return "ok", nil
}

func TestRunWithHooks(t *testing.T) {
	mock := &MockExecutor{}
	cfg := &config.Config{
		Repositories: []config.Repository{
			{Path: "."},
		},
		Hooks: config.Hook{
			Pre:  "echo pre",
			Post: "echo post",
		},
		Env: map[string]string{},
	}
	err := RunWithHooks(cfg, mock, cfg.Repositories[0], []string{"status"})
	assert.NoError(t, err)
	assert.Contains(t, mock.CalledWith[0], "git status")
}
