package cmd

import (
	"fmt"
	"github.com/immarktube/dockyard-cli/command"
	"github.com/immarktube/dockyard-cli/config"
	"github.com/immarktube/dockyard-cli/executor"
	"os"

	"github.com/spf13/cobra"
)

var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "Run 'git status' across all repositories",
	Run: func(cmd *cobra.Command, args []string) {
		cfg, err := config.LoadConfig()
		if err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "Error loading config: %v\n", err)
			os.Exit(1)
		}

		exec := &executor.RealExecutor{Env: cfg.Env}
		for _, repo := range cfg.Repositories {
			fmt.Printf("\n==> Checking status for %s\n", repo.Path)

			if err := command.RunWithHooks(cfg, exec, repo, []string{"status"}); err != nil {
				_, _ = fmt.Fprintf(os.Stderr, "Error checking status for %s: %v\n", repo.Path, err)
			}

		}
	},
}
