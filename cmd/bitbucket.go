package cmd

import "github.com/spf13/cobra"

var bitbucketCmd = &cobra.Command{
	Use:   "bitbucket",
	Short: "Manage Bitbucket pipelines",
	Long:  `Commands for managing Bitbucket pipelines, including viewing builds, triggering jobs, and managing projects.`,
}

var bitbucketListCmd = &cobra.Command{
	Use:   "list",
	Short: "List Bitbucket objects",
	Long:  `List all Bitbucket objects for the current repository or specified repository.`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Println("Bitbucket objects command - not implemented yet")
	},
}

func init() {
	rootCmd.AddCommand(bitbucketCmd)
	bitbucketCmd.AddCommand(bitbucketListCmd)
}
