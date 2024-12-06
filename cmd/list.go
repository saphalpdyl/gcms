package cmd

import (
	"github.com/saphalpdyl/gcms/handlers"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Display all the files being tracked",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		handler.List(handlers.ListHandlerParams{
			RepositoryFolderPath: repoFolderPath,
		})
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
