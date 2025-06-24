package cmd

import (
	"fmt"
	"github.com/immarktube/dockyard-cli/command"
	"github.com/immarktube/dockyard-cli/config"
	"github.com/immarktube/dockyard-cli/executor"
	"github.com/immarktube/dockyard-cli/utils"
	"github.com/spf13/cobra"
	"os"
)

var syncCmd = &cobra.Command{
	Use:   "sync",
	Short: "Run 'git fetch' and 'git pull' across all repositories",
	Run: func(cmd *cobra.Command, args []string) {
		cfg, err := config.LoadConfig()
		if err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "Error loading config: %v\n", err)
			os.Exit(1)
		}
		exec := &executor.RealExecutor{Env: cfg.Env}
		maxConcurrency := utils.GetConcurrency(utils.MaxConcurrency, cfg)
		utils.ForEachRepoConcurrently(cfg.Repositories, func(repo config.Repository) {
			fmt.Printf("\n==> Fetching and Pulling %s\n", repo.Path)
			command.RunGit(repo, exec, "fetch", "--all")
			command.RunGit(repo, exec, "pull", "--rebase", "--autostash")
			if len(command.GetFailedRepos()) > 0 {
				fmt.Println("Some repositories failed at stage fetching and pulling:\n", command.GetFailedRepos())
				command.ClearFailedRepos()
			}
		}, maxConcurrency)

	},
}
