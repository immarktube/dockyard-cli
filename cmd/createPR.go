package cmd

import (
	"fmt"
	"github.com/immarktube/dockyard-cli/config"
	"github.com/immarktube/dockyard-cli/utils"
	"os"

	"github.com/spf13/cobra"
)

var prTitle string
var prBody string

var prCmd = &cobra.Command{
	Use:   "createPR",
	Short: "Create pull requests for all modified repositories",
	RunE: func(cmd *cobra.Command, args []string) error {
		cfg, err := config.LoadConfig()
		if err != nil {
			return err
		}

		utils.ForEachRepoConcurrently(cfg.Repositories, func(repo config.Repository) {
			if repo.AuthToken == "" || repo.Owner == "" || repo.Name == "" {
				fmt.Fprintf(os.Stderr, "❌ Missing AuthToken, Owner, or Name for %s\n", repo.Path)
				return
			}

			err := utils.CreatePullRequest(repo, prTitle, prBody)
			if err != nil {
				fmt.Fprintf(os.Stderr, "❌ Failed to create PR for %s: %v\n", repo.Path, err)
			}
		}, 5)

		return nil
	},
}

func init() {
	rootCmd.AddCommand(prCmd)
	prCmd.Flags().StringVar(&prTitle, "title", "chore: automated patch", "Title of the pull request")
	prCmd.Flags().StringVar(&prBody, "body", "This PR was created automatically by Dockyard.", "Body of the pull request")
}
