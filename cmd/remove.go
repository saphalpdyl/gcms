package cmd

import (
	"github.com/saphalpdyl/gcms/handlers"
	"github.com/spf13/cobra"
)

var removeCommand = &cobra.Command{
	Use:   "remove <filename>",
	Short: "Removes the file from the system including metadata",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		removeFilePath := args[0]

		handler.Remove(handlers.RemoveHandlerParams{
			RepositoryFolderPath: repoFolderPath,
			FilePathToRemove:     removeFilePath,
		})
	},
}

func init() {
	rootCmd.AddCommand(removeCommand)
}
