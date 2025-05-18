package cmd

import (
	"fmt"
	"github.com/immarktube/dockyard-cli/command"
	"github.com/immarktube/dockyard-cli/config"
	"github.com/immarktube/dockyard-cli/executor"
	"os"

	"github.com/spf13/cobra"
)

var execCmd = &cobra.Command{
	Use:   "exec [git args]",
	Short: "Run arbitrary git command across all repositories",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		cfg, err := config.LoadConfig()
		if err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "Error loading config: %v\n", err)
			os.Exit(1)
		}
		exec := &executor.RealExecutor{Env: cfg.Env}

		for _, repo := range cfg.Repositories {
			fmt.Printf("\n==> Executing git %v in %s\n", args, repo.Path)

			if err := command.RunWithHooks(cfg, exec, repo, args); err != nil {
				_, _ = fmt.Fprintf(os.Stderr, "Error executing git %v in %s: %v\n", args, repo.Path, err)
			}
		}
	},
}
