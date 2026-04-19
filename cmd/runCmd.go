package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/immarktube/dockyard-cli/config"
	"github.com/immarktube/dockyard-cli/utils"

	"github.com/spf13/cobra"
)

var runCmd = &cobra.Command{
	Use:   "run [shell command]",
	Short: "Run arbitrary shell command in all repositories",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		commandStr := strings.Join(args, " ")
		cfg, err := config.LoadConfig()
		if err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "Error loading config: %v\n", err)
			os.Exit(1)
		}

		maxConcurrency := utils.GetConcurrency(utils.MaxConcurrency, cfg)
		script := args[0]

		if !filepath.IsAbs(script) {
			abs, err := filepath.Abs(script)
			if err == nil {
				script = abs
			}
		}
		fmt.Printf("\n==> Execute File Path in %s\n", script)
		utils.ForEachRepoConcurrently(cfg.Repositories, func(repo config.Repository) {
			fmt.Printf("\n==> Running command in %s\n", repo.Path)
			var c *exec.Cmd
			if strings.HasSuffix(args[0], ".sh") {
				c = exec.Command("sh", script)
			} else if strings.HasSuffix(args[0], ".py") {
				c = exec.Command("python", args[0])
			} else {
				c = exec.Command("sh", "-c", commandStr)
			}
			c.Dir = repo.Path
			c.Stdout = os.Stdout
			c.Stderr = os.Stderr
			if err := c.Run(); err != nil {
				_, _ = fmt.Fprintf(os.Stderr, "Error running command in %s: %v\n", repo.Path, err)
			}
		}, maxConcurrency)
	},
}
