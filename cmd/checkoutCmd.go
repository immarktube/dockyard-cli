package cmd

import (
	"bufio"
	"fmt"
	"github.com/immarktube/dockyard-cli/command"
	"github.com/immarktube/dockyard-cli/config"
	"github.com/immarktube/dockyard-cli/executor"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

func switchToBranch(repo config.Repository, exec executor.Executor, branch string) error {
	fmt.Printf("==> Checking out branch '%s' in %s\n", branch, repo.Path)

	// 尝试切换本地分支
	switchCmd := command.GitCommand{
		Repo: repo, Executor: exec, Args: []string{"switch", branch},
	}
	if err := switchCmd.Run(); err == nil {
		return nil
	}

	// fallback：提示用户输入远程分支（默认 origin/master）
	fmt.Printf("Failed to switch to '%s'.\n", branch)
	fmt.Printf("Please enter the remote upstream branch to track (default: master): ")

	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	upstreamBranch := strings.TrimSpace(input)
	if upstreamBranch == "" {
		upstreamBranch = "master"
	}

	// 创建并跟踪远程分支
	createCmd := command.GitCommand{
		Repo: repo, Executor: exec,
		Args: []string{"switch", "-c", branch, "--track", upstreamBranch},
	}
	if err := createCmd.Run(); err != nil {
		return fmt.Errorf("failed to create and switch to branch '%s' tracking '%s': %w",
			branch, upstreamBranch, err)
	}

	return nil
}

var checkoutCmd = &cobra.Command{
	Use:   "checkout [branch]",
	Short: "Batch checkout branch in all repositories",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		branch := args[0]
		cfg, err := config.LoadConfig()
		if err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "Error loading config: %v\n", err)
			os.Exit(1)
		}
		exec := &executor.RealExecutor{Env: cfg.Env}

		for _, repo := range cfg.Repositories {
			if err := switchToBranch(repo, exec, branch); err != nil {
				_, _ = fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			}
		}

	},
}
