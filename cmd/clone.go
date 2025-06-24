package cmd

import (
	"fmt"
	"github.com/immarktube/dockyard-cli/command"
	"github.com/immarktube/dockyard-cli/config"
	"github.com/immarktube/dockyard-cli/utils"
	"github.com/spf13/cobra"
	"os"
)

var cloneCmd = &cobra.Command{
	Use:   "clone",
	Short: "Run 'git clone' across all repositories",
	Run: func(cmd *cobra.Command, args []string) {
		cfg, err := config.LoadConfig()
		if err != nil {
			utils.SafeError("Error loading config: %v\n", err)
			os.Exit(1)
		}

		maxConcurrency := utils.GetConcurrency(maxConcurrency, cfg)
		utils.ForEachRepoConcurrently(cfg.Repositories, func(repo config.Repository) {
			repoPath := repo.Path
			if _, err := os.Stat(repoPath); os.IsNotExist(err) {
				fmt.Printf("Cloning %s...\n", repoPath)
				remoteUrl := utils.BuildRemoteURL(repo, cfg.Global)
				command.CloneRepo(remoteUrl, repo, cfg.Global)
			} else {
				fmt.Printf("Repository %s already exists, skipping.\n", repoPath)
			}
		}, maxConcurrency)
	},
}
