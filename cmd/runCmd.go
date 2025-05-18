package cmd

import (
	"fmt"
	"github.com/immarktube/dockyard-cli/config"
	"os"
	"os/exec"
	"strings"

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

		for _, repo := range cfg.Repositories {
			fmt.Printf("\n==> Running command in %s\n", repo.Path)
			c := exec.Command("sh", "-c", commandStr)
			c.Dir = repo.Path
			c.Stdout = os.Stdout
			c.Stderr = os.Stderr
			if err := c.Run(); err != nil {
				_, _ = fmt.Fprintf(os.Stderr, "Error running command in %s: %v\n", repo.Path, err)
			}
		}
	},
}
