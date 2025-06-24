package cmd

import (
	"fmt"
	"github.com/immarktube/dockyard-cli/utils"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "dockyard",
	Short: "Dockyard is a multi-repo Git management tool",
	Long:  `Dockyard helps you perform batch Git operations like pull, push, status, or running scripts across multiple repositories from a central CLI interface.`,
	Run: func(cmd *cobra.Command, args []string) {
		utils.SafePrint("Welcome to Dockyard! Use `dockyard --help` to get started.")
	},
}

var maxConcurrency int

// Execute is the CLI entry point
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(syncCmd)
	rootCmd.AddCommand(statusCmd)
	rootCmd.AddCommand(pushCmd)
	rootCmd.AddCommand(cloneCmd)
	rootCmd.AddCommand(execCmd)
	rootCmd.AddCommand(checkoutCmd)
	rootCmd.AddCommand(runCmd)
	rootCmd.PersistentFlags().IntVar(&maxConcurrency, "concurrency", 5, "Global max concurrency (overridden by command-level settings if set)")
}
