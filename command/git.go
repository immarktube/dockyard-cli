package command

import (
	"fmt"
	"github.com/immarktube/dockyard-cli/config"
	"github.com/immarktube/dockyard-cli/executor"
	"github.com/immarktube/dockyard-cli/utils"
	"os"
	"os/exec"
	"path/filepath"
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
	utils.SafePrint("Running git command: %s", output)
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

// CloneRepo clones a Git repository into the specified path.
func CloneRepo(remoteURL string, repo config.Repository, globalConfig config.GlobalConfig) {

	// 确保目标目录的父路径存在
	parentDir := filepath.Dir(repo.Path)
	if err := os.MkdirAll(parentDir, 0755); err != nil {
		utils.SafeError("❌ Failed to create parent directory %s: %v\n", parentDir, err)
		return
	}

	if globalConfig.AuthToken != "" {
		// 如果配置了全局认证令牌，则注入到远程 URL 中
		remoteURL = InjectToken(remoteURL, globalConfig.AuthToken)
	}

	// 执行 git clone
	cmd := exec.Command("git", "clone", remoteURL, repo.Path)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	utils.SafePrint("🚀 Cloning %s into %s...\n", remoteURL, repo.Path)
	if err := cmd.Run(); err != nil {
		utils.SafeError("❌ Failed to clone %s: %v\n", repo.Path, err)
	}
}

func InjectToken(cloneURL, token string) string {
	if strings.HasPrefix(cloneURL, "https://") {
		return strings.Replace(cloneURL, "https://", fmt.Sprintf("https://%s@", token), 1)
	}
	return cloneURL
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

	if utils.NoHookFlag || cfg.Global.NoHook {
		git := GitCommand{Repo: repo, Executor: exec, Args: args}
		if err := git.Run(); err != nil {
			return err
		}
		return nil
	}

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
