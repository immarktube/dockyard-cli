package executor

import (
	"bytes"
	"fmt"
	"os/exec"
)

type Executor interface {
	RunCommand(dir string, name string, args ...string) (string, error)
}

type RealExecutor struct {
	Env map[string]string
}

func (e *RealExecutor) RunCommand(dir string, name string, args ...string) (string, error) {
	cmd := exec.Command(name, args...)
	cmd.Dir = dir

	if e.Env != nil {
		var env []string
		for k, v := range e.Env {
			env = append(env, fmt.Sprintf("%s=%s", k, v))
		}
		cmd.Env = append(env, cmd.Env...)
	}

	var out bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &out

	err := cmd.Run()
	return out.String(), err
}
