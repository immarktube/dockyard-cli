package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/immarktube/dockyard-cli/config"
	"github.com/immarktube/dockyard-cli/executor"
	"github.com/immarktube/dockyard-cli/utils"
	"github.com/spf13/cobra"
)

var (
	applyMessage string
	includeStr   string
	excludeStr   string
	applyAll     bool
	applyDryRun  bool
)

var applyCmd = &cobra.Command{
	Use:   "apply",
	Short: "Stage selected files and commit across all repositories",
	Run: func(cmd *cobra.Command, args []string) {

		if includeStr == "" && !applyAll {
			fmt.Println("❌ You must specify --include or --all")
			return
		}

		cfg, err := config.LoadConfig()
		if err != nil {
			_, err := fmt.Fprintf(os.Stderr, "Error loading config: %v\n", err)
			if err != nil {
				return
			}
			os.Exit(1)
		}

		exec := &executor.RealExecutor{Env: cfg.Env}
		maxConcurrency := utils.GetConcurrency(utils.MaxConcurrency, cfg)

		includes := splitPatterns(includeStr)
		excludes := splitPatterns(excludeStr)

		utils.ForEachRepoConcurrently(cfg.Repositories, func(repo config.Repository) {

			fmt.Printf("\n==> Applying in %s\n", repo.Path)

			var filesToAdd []string

			// 1️⃣ resolve include
			if len(includes) > 0 {
				filesToAdd = resolveFiles(repo.Path, includes)
			} else if applyAll {
				filesToAdd = []string{"."}
			}

			if len(filesToAdd) == 0 {
				fmt.Printf("⚠️ No files matched in %s, skipping\n", repo.Path)
				return
			}

			// 2️⃣ dry-run 输出
			if applyDryRun {
				fmt.Printf("[DRY-RUN] %s -> add: %v\n", repo.Path, filesToAdd)
				return
			}

			fmt.Printf("DEBUG: filesToAdd=%v\n", filesToAdd)
			// 3️⃣ git add
			if _, err := exec.RunCommand(repo.Path, "git", append([]string{"add"}, filesToAdd...)...); err != nil {
				fmt.Printf("❌ Failed to add files in %s: %v\n", repo.Path, err)
				return
			}

			// 4️⃣ exclude
			if len(excludes) > 0 {
				excludeFiles := resolveFiles(repo.Path, excludes)
				if len(excludeFiles) > 0 {
					_, _ = exec.RunCommand(repo.Path, "git", append([]string{"reset"}, excludeFiles...)...)
				}
			}

			// 5️⃣ commit
			msg := applyMessage
			if msg == "" {
				msg = "dockyard: apply changes"
			}

			out, err := exec.RunCommand(repo.Path, "git", "commit", "-m", msg)
			if err != nil {
				if strings.Contains(out, "nothing to commit") {
					fmt.Printf("ℹ️ Nothing to commit in %s\n", repo.Path)
					return
				}
				fmt.Printf("❌ Commit failed in %s: %v\nOutput: %s\n", repo.Path, err, out)
				return
			}

			fmt.Printf("✅ Committed in %s\n", repo.Path)

		}, maxConcurrency)
	},
}

func splitPatterns(input string) []string {
	if input == "" {
		return nil
	}
	parts := strings.Split(input, ",")
	var result []string
	for _, p := range parts {
		p = strings.TrimSpace(p)
		if p != "" {
			result = append(result, p)
		}
	}
	return result
}

func resolveFiles(repoPath string, patterns []string) []string {
	var matched []string

	for _, pattern := range patterns {
		fullPattern := filepath.Join(repoPath, pattern)
		files, _ := filepath.Glob(fullPattern)

		for _, f := range files {
			rel, err := filepath.Rel(repoPath, f)
			if err == nil {
				matched = append(matched, rel)
			}
		}
	}

	return unique(matched)
}

func unique(input []string) []string {
	m := make(map[string]struct{})
	var result []string
	for _, v := range input {
		if _, ok := m[v]; !ok {
			m[v] = struct{}{}
			result = append(result, v)
		}
	}
	return result
}

func init() {
	rootCmd.AddCommand(applyCmd)

	applyCmd.Flags().StringVarP(&applyMessage, "message", "m", "", "Commit message")
	applyCmd.Flags().StringVar(&includeStr, "include", "", "Files or glob patterns to include (comma-separated)")
	applyCmd.Flags().StringVar(&excludeStr, "exclude", "", "Files or glob patterns to exclude (comma-separated)")
	applyCmd.Flags().BoolVar(&applyAll, "all", false, "Include all files (git add .)")
	applyCmd.Flags().BoolVar(&applyDryRun, "dry-run", false, "Preview changes without committing")
}
