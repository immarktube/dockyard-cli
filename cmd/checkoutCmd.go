package cmd

import (
	"fmt"
	"github.com/immarktube/dockyard-cli/command"
	"github.com/immarktube/dockyard-cli/config"
	"github.com/immarktube/dockyard-cli/executor"
	"os"

	"github.com/spf13/cobra"
)

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
			fmt.Printf("\n==> Checking out branch '%s' in %s\n", branch, repo.Path)

			if err := command.RunWithHooks(cfg, exec, repo, []string{"checkout", branch}); err != nil {
				_, _ = fmt.Fprintf(os.Stderr, "Error checking out branch in %s: %v\n", repo.Path, err)
			}
		}
	},
}
