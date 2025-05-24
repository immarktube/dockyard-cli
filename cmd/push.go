package cmd

import (
	"fmt"
	"github.com/immarktube/dockyard-cli/command"
	"github.com/immarktube/dockyard-cli/config"
	"github.com/immarktube/dockyard-cli/executor"
	"github.com/immarktube/dockyard-cli/utils"
	"os"

	"github.com/spf13/cobra"
)

var pushCmd = &cobra.Command{
	Use:   "push",
	Short: "Run 'git push' across all repositories",
	Run: func(cmd *cobra.Command, args []string) {
		cfg, err := config.LoadConfig()
		if err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "Error loading config: %v\n", err)
			os.Exit(1)
		}

		exec := &executor.RealExecutor{Env: cfg.Env}

		utils.ForEachRepoConcurrently(cfg.Repositories, func(repo config.Repository) {
			fmt.Printf("\n==> Pushing %s\n", repo.Path)
			command.RunGit(repo, exec, "push", ".", "HEAD")
			if len(command.GetFailedRepos()) > 0 {
				_, _ = fmt.Fprintf(os.Stderr, "Error pushing %s: %v\n", command.GetFailedRepos(), err)
				command.ClearFailedRepos()
			}
		})
	},
}
