package command

import (
	"fmt"
	"github.com/immarktube/dockyard-cli/config"
	"github.com/immarktube/dockyard-cli/executor"
	"os"
	"os/exec"
	"strings"
	"sync"
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

var (
	mu          sync.Mutex
	failedRepos []string
)

// RunGit executes a git command and tracks failures.
func RunGit(repo config.Repository, exec executor.Executor, args ...string) {
	fmt.Printf("==> %s: git %s\n", repo.Path, strings.Join(args, " "))
	cmd := GitCommand{Repo: repo, Executor: exec, Args: args}
	if err := cmd.Run(); err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Error running git %s in %s: %v\n", strings.Join(args, " "), repo.Path, err)
		mu.Lock()
		failedRepos = append(failedRepos, repo.Path)
		mu.Unlock()
	}
}

// GetFailedRepos returns the list of failed repos
func GetFailedRepos() []string {
	mu.Lock()
	defer mu.Unlock()
	return append([]string{}, failedRepos...) // 返回副本避免外部修改
}

// ClearFailedRepos resets the failure list
func ClearFailedRepos() {
	mu.Lock()
	defer mu.Unlock()
	failedRepos = nil
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
