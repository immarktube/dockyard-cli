package cmd

import (
	"io"
	"os"
	"path/filepath"

	"github.com/immarktube/dockyard-cli/config"
	"github.com/immarktube/dockyard-cli/utils"
	"github.com/spf13/cobra"
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

		sourcePath, _ := cmd.Flags().GetString("source")
		targetPath, _ := cmd.Flags().GetString("target")
		copyDryRun, _ := cmd.Flags().GetBool("dry-run")

		utils.ForEachRepoConcurrently(cfg.Repositories, func(repo config.Repository) {
			src := filepath.Join(repo.Path, sourcePath)
			dst := filepath.Join(repo.Path, targetPath)

			if copyDryRun {
				utils.SafePrint("📝 Dry-run: Would copy %s → %s\n", src, dst)
				return
			}

			err = os.MkdirAll(filepath.Dir(dst), os.ModePerm)
			if err != nil {
				utils.SafeError("❌ Failed to create target directory %s: %v\n", filepath.Dir(dst), err)
				return
			}

			if err := copyFile(src, dst); err != nil {
				utils.SafeError("❌ Failed to copy %s to %s: %v\n", src, dst, err)
				return
			}

			utils.SafePrint("✅%s: Copied %s → %s\n", repo.Path, src, dst)
		}, maxConcurrency)

		return nil
	},
}

func init() {
	rootCmd.AddCommand(copyFileCmd)
	copyFileCmd.Flags().String("source", "", "Source relative file path (required)")
	copyFileCmd.Flags().String("target", "", "Target relative file path (required)")
	copyFileCmd.Flags().Bool("dry-run", false, "Preview the copy and commit without making changes")

	err := copyFileCmd.MarkFlagRequired("source")
	err = copyFileCmd.MarkFlagRequired("target")
	err = copyFileCmd.MarkFlagRequired("message")
	if err != nil {
		return
	}
}

func copyFile(src, dst string) error {
	in, err := os.Open(src)
	if err != nil {
		return err
	}
	defer in.Close()

	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer out.Close()

	if _, err := io.Copy(out, in); err != nil {
		return err
	}
	return out.Sync()
}
