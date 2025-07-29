package cmd

import (
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/immarktube/dockyard-cli/config"
	"github.com/immarktube/dockyard-cli/executor"
	"github.com/immarktube/dockyard-cli/utils"
	"github.com/spf13/cobra"
)

var (
	sourcePath    string
	targetPath    string
	copyCommitMsg string
	copyDryRun    bool
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
		exec := &executor.RealExecutor{Env: cfg.Env}

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

			utils.SafePrint("✅ Copied %s → %s\n", src, dst)

			exec.RunCommand(repo.Path, "git", "add", targetPath)
			msg := copyCommitMsg
			if msg == "" {
				msg = fmt.Sprintf("dockyard: copy file %s", targetPath)
			}
			out, err := exec.RunCommand(repo.Path, "git", "commit", "-m", msg)
			if err != nil {
				utils.SafeError("❌ Failed to commit in %s: %v\nOutput: %s\\n", repo.Path, err, out)
				return
			}
		}, maxConcurrency)

		return nil
	},
}

func init() {
	rootCmd.AddCommand(copyFileCmd)
	copyFileCmd.Flags().StringVar(&sourcePath, "source", "", "Source relative file path (required)")
	copyFileCmd.Flags().StringVar(&targetPath, "target", "", "Target relative file path (required)")
	copyFileCmd.Flags().StringVar(&copyCommitMsg, "message", "", "git commit message (required)")
	copyFileCmd.Flags().BoolVar(&copyDryRun, "dry-run", false, "Preview the copy and commit without making changes")

	copyFileCmd.MarkFlagRequired("source")
	copyFileCmd.MarkFlagRequired("target")
	copyFileCmd.MarkFlagRequired("message")
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
