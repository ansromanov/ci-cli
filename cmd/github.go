package cmd

import (
	"ci/internal/config"
	"ci/internal/providers/github"
	"context"
	"fmt"
	"time"

	"github.com/spf13/cobra"
)

var githubCmd = &cobra.Command{
	Use:   "github",
	Short: "Manage GitHub Actions workflows",
	Long:  `Commands for managing GitHub Actions workflows, including viewing runs, triggering workflows, and managing repositories.`,
}

var githubLoginCmd = &cobra.Command{
	Use:   "login",
	Short: "Login to GitHub",
	Long:  `Login to GitHub using a personal access token. The token can be provided via environment variable GITHUB_TOKEN or interactively.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return gitHubLogin(cmd)
	},
}

func gitHubLogin(cmd *cobra.Command) error {
	// Load configuration
	cfg := config.Load()

	// Check if token is already set
	if cfg.Providers.GitHub.Token != "" {
		cmd.Println("GitHub token found in configuration")

		// Validate the token
		client := github.NewClient(cfg.Providers.GitHub.Token)
		if client == nil {
			return fmt.Errorf("failed to create GitHub client")
		}

		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		if err := client.ValidateToken(ctx); err != nil {
			return fmt.Errorf("invalid GitHub token: %w", err)
		}

		user, err := client.GetAuthenticatedUser(ctx)
		if err != nil {
			return fmt.Errorf("failed to get user information: %w", err)
		}

		cmd.Printf("✅ Successfully authenticated as: %s\n", *user.Login)
		return nil
	}

	// Interactive token input
	cmd.Println("GitHub token not found in configuration")
	cmd.Println("Please provide your GitHub Personal Access Token:")
	cmd.Println("1. Go to https://github.com/settings/tokens")
	cmd.Println("2. Generate a new token with appropriate permissions")
	cmd.Println("3. Set it in one of the following ways:")
	cmd.Println("   - Environment variable: GITHUB_TOKEN")
	cmd.Println("   - .env file: GITHUB_TOKEN=your_token")

	return fmt.Errorf("interactive token input not implemented yet. Please set GITHUB_TOKEN in environment or .env file")
}

var githubListCmd = &cobra.Command{
	Use:   "list",
	Short: "List GitHub Actions objects",
	Long:  `List all GitHub Actions objects for the current repository or specified repository.`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: Implement GitHub Actions workflows listing
		cmd.Println("GitHub Actions objects command - not implemented yet")
	},
}

var githubListRepositoriesCmd = &cobra.Command{
	Use:   "repositories",
	Short: "List GitHub repositories",
	Long:  `List all GitHub repositories for the current user or specified user.`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: Implement GitHub repositories listing
		cmd.Println("GitHub repositories command - not implemented yet")
	},
}

var githubListWorkflowsCmd = &cobra.Command{
	Use:   "workflows",
	Short: "List GitHub Actions workflows",
	Long:  `List all GitHub Actions workflows for the current repository or specified repository.`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: Implement GitHub Actions workflows listing
		cmd.Println("GitHub Actions workflows command - not implemented yet")
	},
}

var githubTriggerCmd = &cobra.Command{
	Use:   "trigger [workflow-name]",
	Short: "Trigger a GitHub Actions workflow",
	Long:  `Trigger a new GitHub Actions workflow with optional inputs.`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: Implement GitHub Actions workflow triggering
		cmd.Printf("Triggering GitHub Actions workflow: %s - not implemented yet\n", args[0])
	},
}

func init() {
	rootCmd.AddCommand(githubCmd)
	githubCmd.AddCommand(githubLoginCmd)
	// List
	githubCmd.AddCommand(githubListCmd)
	githubListCmd.AddCommand(githubListWorkflowsCmd)
	githubListCmd.AddCommand(githubListRepositoriesCmd)

	// Trigger
	githubCmd.AddCommand(githubTriggerCmd)
}
