package cmd

import "github.com/spf13/cobra"

var updateCommand = &cobra.Command{
	Use:   "update",
	Short: "Pull from the remote repository",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		githubService.UpdateRepository()
	},
}

func init() {
	rootCmd.AddCommand(updateCommand)
}
