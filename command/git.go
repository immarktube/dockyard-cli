package command

import (
	"fmt"
	"github.com/immarktube/dockyard-cli/config"
	"github.com/immarktube/dockyard-cli/executor"
	"os"
	"os/exec"
	"strings"
)

type GitCommand struct {
	Repo     config.Repository
	Executor executor.Executor
	Args     []string
}

func (g *GitCommand) Run() error {
	output, err := g.Executor.RunCommand(g.Repo.Path, "git", g.Args...)
	fmt.Print(output)
	return err
}

func runShellScript(dir, script string, env map[string]string) error {
	if strings.TrimSpace(script) == "" {
		return nil
	}
	cmd := exec.Command("sh", "-c", script)
	cmd.Dir = dir
	if env != nil {
		var envList []string
		for k, v := range env {
			envList = append(envList, fmt.Sprintf("%s=%s", k, v))
		}
		cmd.Env = append(envList, cmd.Env...)
	}
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func RunWithHooks(cfg *config.Config, exec executor.Executor, repo config.Repository, args []string) error {
	hooks := config.GetHooksForRepo(cfg, repo) // 使用新函数

	if err := runShellScript(repo.Path, hooks.Pre, cfg.Env); err != nil {
		return fmt.Errorf("pre-hook failed: %w", err)
	}

	git := GitCommand{Repo: repo, Executor: exec, Args: args}
	if err := git.Run(); err != nil {
		return err
	}

	if err := runShellScript(repo.Path, hooks.Post, cfg.Env); err != nil {
		return fmt.Errorf("post-hook failed: %w", err)
	}

	return nil
}
