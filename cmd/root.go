package cmd

import (
	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:   "ci",
		Short: "A CLI tool for managing CI/CD pipelines",
		Long: `ci is a command line tool that provides a unified interface
for managing CI/CD pipelines across different providers like GitLab, GitHub Actions, and CircleCI.`,
	}
)

// Execute adds all child commands to the root command and sets flags appropriately.
func Execute() error {
	return rootCmd.Execute()
}

func init() {
	// Remove the initConfig call since .env is loaded when config.Load() is called
}
