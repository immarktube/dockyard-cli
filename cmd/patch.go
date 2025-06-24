package cmd

import (
	"fmt"
	"github.com/immarktube/dockyard-cli/config"
	"github.com/immarktube/dockyard-cli/executor"
	"github.com/immarktube/dockyard-cli/utils"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/spf13/cobra"
)

var (
	dryRun         bool
	patchFile      string
	patchOld       string
	patchNew       string
	patchRegex     bool
	patchCommitMsg string
)

var patchCmd = &cobra.Command{
	Use:   "patch",
	Short: "Modify a specific file in all repositories and commit the change.",
	Example: `dockyard patch --file .env --old 'DB_HOST=localhost' --new 'DB_HOST=
db.example.com' --message 'Update database host'`,
	RunE: func(cmd *cobra.Command, args []string) error {
		cfg, err := config.LoadConfig()
		if err != nil {
			return err
		}

		exec := &executor.RealExecutor{Env: cfg.Env}
		maxConcurrency := utils.GetConcurrency(utils.MaxConcurrency, cfg)
		utils.ForEachRepoConcurrently(cfg.Repositories, func(repo config.Repository) {
			targetFile := filepath.Join(repo.Path, patchFile)

			contentBytes, err := os.ReadFile(targetFile)
			if err != nil {
				utils.SafeError("❌ Cannot read %s: %v\n", targetFile, err)
				return
			}

			content := string(contentBytes)
			var modified string

			if patchRegex {
				re, err := regexp.Compile(patchOld)
				if err != nil {
					utils.SafeError("❌ Invalid regex '%s': %v\n", patchOld, err)
					return
				}
				modified = re.ReplaceAllString(content, patchNew)
			} else {
				modified = strings.ReplaceAll(content, patchOld, patchNew)
			}

			if content == modified {
				utils.SafePrint("✅ %s already up to date.\n", targetFile)
				return
			}

			if dryRun {
				utils.SafePrint("📝 Dry-run: Would modify %s\n", targetFile)
				return
			}

			if err := os.WriteFile(targetFile, []byte(modified), 0644); err != nil {
				utils.SafeError("❌ Failed to write %s: %v\n", targetFile, err)
				return
			}

			utils.SafePrint("📦 Committing change in %s\n", repo.Path)

			exec.RunCommand(repo.Path, "git", "add", patchFile)
			msg := patchCommitMsg
			if msg == "" {
				msg = fmt.Sprintf("chore: patch file %s", patchFile)
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
	rootCmd.AddCommand(patchCmd)

	patchCmd.Flags().StringVar(&patchFile, "file", os.Getenv("PATCH_FILE"), "File path to patch (relative to repo root)")
	patchCmd.Flags().StringVar(&patchOld, "old", os.Getenv("PATCH_OLD"), "Text or pattern to replace")
	patchCmd.Flags().StringVar(&patchNew, "new", os.Getenv("PATCH_NEW"), "Replacement text")
	patchCmd.Flags().BoolVar(&dryRun, "dry-run", false, "Perform a dry run without modifying files")
	patchCmd.Flags().BoolVar(&patchRegex, "regex", false, "Treat --old as regular expression")
	patchCmd.Flags().StringVar(&patchCommitMsg, "message", os.Getenv("PATCH_COMMIT_MSG"), "Commit message to use (optional)")
}
