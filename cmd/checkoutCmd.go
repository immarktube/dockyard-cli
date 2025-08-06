package cmd

import (
	"fmt"
	"github.com/immarktube/dockyard-cli/command"
	"github.com/immarktube/dockyard-cli/config"
	"github.com/immarktube/dockyard-cli/executor"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

func isLikelyRemoteBranch(exec executor.Executor, repo config.Repository, base string) bool {
	out, err := exec.RunCommand(repo.Path, "git", "ls-remote", "--heads", "origin", base)
	return err == nil && strings.Contains(out, base)
}

func checkoutBranch(repo config.Repository, exec executor.Executor, newBranch string) error {
	fmt.Printf("==> %s: checking out branch '%s' from '%s'\n", newBranch, repo.Path)

	cmd := command.GitCommand{
		Repo: repo, Executor: exec,
		Args: []string{"switch", newBranch},
	}
	if err := cmd.Run(); err == nil {
		return nil
	}

	base := strings.TrimSpace(repo.BaseRef)
	if base == "" {
		base = "master"
	}

	isBranch := isLikelyRemoteBranch(exec, repo, base)
	var args []string
	if isBranch {
		args = []string{"switch", "-c", newBranch, "--track", "origin/" + base}
	} else {
		args = []string{"switch", "-c", newBranch, base}
	}

	createCmd := command.GitCommand{
		Repo: repo, Executor: exec,
		Args: args,
	}
	if err := createCmd.Run(); err != nil {
		return fmt.Errorf("❌ failed to create/switch to branch '%s' from '%s': %w", newBranch, base, err)
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
			if err := checkoutBranch(repo, exec, branch); err != nil {
				_, _ = fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			}
		}

	},
}
