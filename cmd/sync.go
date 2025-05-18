package cmd

import (
	"fmt"
	"github.com/immarktube/dockyard-cli/command"
	"github.com/immarktube/dockyard-cli/config"
	"github.com/immarktube/dockyard-cli/executor"
	"os"

	"github.com/spf13/cobra"
)

var syncCmd = &cobra.Command{
	Use:   "sync",
	Short: "Run 'git pull' across all repositories",
	Run: func(cmd *cobra.Command, args []string) {
		cfg, err := config.LoadConfig()
		if err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "Error loading config: %v\n", err)
			os.Exit(1)
		}

		exec := &executor.RealExecutor{Env: cfg.Env}
		for _, repo := range cfg.Repositories {
			fmt.Printf("\n==> Pulling %s\n", repo.Path)
			cmd := &command.GitCommand{Repo: repo, Executor: exec, Args: []string{"pull"}}
			if err := cmd.Run(); err != nil {
				_, _ = fmt.Fprintf(os.Stderr, "Error pulling %s: %v\n", repo.Path, err)
			}
		}
	},
}
