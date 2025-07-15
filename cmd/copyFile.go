package cmd

import (
	"os"
	"path/filepath"

	"github.com/immarktube/dockyard-cli/config"
	"github.com/immarktube/dockyard-cli/utils"
	"github.com/spf13/cobra"
)

var (
	sourcePath string
	targetPath string
)

var copyFileCmd = &cobra.Command{
	Use:   "copyFile",
	Short: "Copy a file from one path to another inside each repository",
	RunE: func(cmd *cobra.Command, args []string) error {
		cfg, err := config.LoadConfig()
		if err != nil {
			return err
		}
		maxConcurrency := utils.GetConcurrency(utils.MaxConcurrency, cfg)
		utils.ForEachRepoConcurrently(cfg.Repositories, func(repo config.Repository) {
			src := filepath.Join(repo.Path, sourcePath)
			dst := filepath.Join(repo.Path, targetPath)

			content, err := os.ReadFile(src)
			if err != nil {
				utils.SafeError("❌ Failed to read from %s: %v\n", src, err)
				return
			}

			err = os.MkdirAll(filepath.Dir(dst), os.ModePerm)
			if err != nil {
				utils.SafeError("❌ Failed to create target directory %s: %v\n", filepath.Dir(dst), err)
				return
			}

			err = os.WriteFile(dst, content, 0644)
			if err != nil {
				utils.SafeError("❌ Failed to write to %s: %v\n", dst, err)
				return
			}

			utils.SafePrint("✅ Copied %s → %s\n", src, dst)
		}, maxConcurrency)

		return nil
	},
}

func init() {
	rootCmd.AddCommand(copyFileCmd)
	copyFileCmd.Flags().StringVar(&sourcePath, "source", "", "Source relative file path (required)")
	copyFileCmd.Flags().StringVar(&targetPath, "target", "", "Target relative file path (required)")
	copyFileCmd.MarkFlagRequired("source")
	copyFileCmd.MarkFlagRequired("target")
}
