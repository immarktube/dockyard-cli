package config

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadConfig(t *testing.T) {
	// 准备模拟文件内容
	yamlContent := `
repositories:
  - path: "/path/to/repo1"
  - path: "/path/to/repo2"
hook:
  pre: "echo pre-hook"
  post: "echo post-hook"
commands:
  - name: "hello"
    run: "echo Hello"
    args: []
`
	_ = os.WriteFile(".dockyard.yaml", []byte(yamlContent), 0644)
	_ = os.WriteFile(".env", []byte("FOO=bar"), 0644)
	defer os.Remove(".dockyard.yaml")
	defer os.Remove(".env")

	cfg, err := LoadConfig()
	assert.NoError(t, err)
	assert.Equal(t, 2, len(cfg.Repositories))
	assert.Equal(t, "bar", cfg.Env["FOO"])
	assert.Equal(t, "echo pre-hook", cfg.Hooks.Pre)
}
