package cmd

import (
	"github.com/saphalpdyl/gcms/handlers"
	"github.com/spf13/cobra"
)

var deleteLocalCommand = &cobra.Command{
	Use:   "delete-local",
	Short: "Delete the local repository [WARNING: Cannot be reversed]",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		handler.DeleteLocal(handlers.DeleteLocalHandlerParams{
			RepositoryFolderPath: repoFolderPath,
		})
	},
}

func init() {
	rootCmd.AddCommand(deleteLocalCommand)
}
