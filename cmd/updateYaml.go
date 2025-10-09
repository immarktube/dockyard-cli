package cmd

import (
	"fmt"
	"github.com/immarktube/dockyard-cli/config"
	"github.com/immarktube/dockyard-cli/executor"
	"github.com/immarktube/dockyard-cli/utils"
	"github.com/spf13/cobra"
	"os"
	"path/filepath"
	"strings"
)

var updateYamlCmd = &cobra.Command{
	Use:     "updateYaml",
	Short:   "Modify a specific yaml file in all repositories and commit the change.",
	Example: `dockyard updateYaml --filePath env --nodePath 'a.b.c' --value 'exampleValue' --createIfAbsent --message 'Update database host'`,
	RunE: func(cmd *cobra.Command, args []string) error {
		cfg, err := config.LoadConfig()
		if err != nil {
			return err
		}

		exec := &executor.RealExecutor{Env: cfg.Env}
		maxConcurrency := utils.GetConcurrency(utils.MaxConcurrency, cfg)

		dryRun, _ := cmd.Flags().GetBool("dry-run")
		filePath, _ := cmd.Flags().GetString("filePath")
		nodePath, _ := cmd.Flags().GetString("nodePath")
		value, _ := cmd.Flags().GetString("value")
		createIfAbsent, _ := cmd.Flags().GetBool("createIfAbsent")
		commitMsg, _ := cmd.Flags().GetString("commitMsg")

		if !strings.HasSuffix(filePath, ".yml") && !strings.HasSuffix(filePath, ".yaml") {
			return fmt.Errorf("file path must end with .yml or .yaml")
		}
		utils.ForEachRepoConcurrently(cfg.Repositories, func(repo config.Repository) {
			targetFilepath := filepath.Join(repo.Path, filePath)

			println("the createIfAbsent is:", createIfAbsent)
			result := utils.UpdateYAMLFile(targetFilepath, map[string]interface{}{
				nodePath: value,
			}, createIfAbsent)

			if result != nil {
				utils.SafeError("❌ Failed to write %s: %v\n", targetFilepath, result)
			}

			if dryRun {
				utils.SafePrint("📝 Dry-run: Would modify %s\n", targetFilepath)
				return
			}

			utils.SafePrint("📦 Committing change in %s\n", repo.Path)

			exec.RunCommand(repo.Path, "git", "add", filePath)
			msg := commitMsg
			if msg == "" {
				msg = fmt.Sprintf("chore: patch file %s", filePath)
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
	rootCmd.AddCommand(updateYamlCmd)

	updateYamlCmd.Flags().String("filePath", os.Getenv("filePath"), "File path to update (relative to repo root)")
	updateYamlCmd.Flags().String("nodePath", os.Getenv("nodePath"), "Node path to update (relative to root node)")
	updateYamlCmd.Flags().String("value", os.Getenv("value"), "Replacement text")
	updateYamlCmd.Flags().Bool("dry-run", false, "Perform a dry run without modifying files")
	updateYamlCmd.Flags().Bool("createIfAbsent", false, "Create node when absent")
	updateYamlCmd.Flags().String("commitMsg", os.Getenv("commitMsg"), "Commit message to use (optional)")
}
