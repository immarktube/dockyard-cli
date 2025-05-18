package executor

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type FakeExecutor struct {
	Output string
	Err    error
}

func (f *FakeExecutor) RunCommand(dir string, name string, args ...string) (string, error) {
	return f.Output, f.Err
}

func TestFakeExecutor(t *testing.T) {
	exec := &FakeExecutor{
		Output: "git pull output",
		Err:    nil,
	}
	out, err := exec.RunCommand("/repo", "git", "pull")
	assert.NoError(t, err)
	assert.Equal(t, "git pull output", out)
}
