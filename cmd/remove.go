package cmd

import (
	"log"

	"github.com/saphalpdyl/gcms/handlers"
	"github.com/spf13/cobra"
)

var removeCommand = &cobra.Command{
	Use:   "remove <filename>",
	Short: "Removes the file from the system including metadata",
	Args: func(cmd *cobra.Command, args []string) error {
		noUIFlag, err := cmd.Flags().GetBool("nui")

		if err != nil {
			log.Fatal(err)
		}

		if noUIFlag && len(args) != 1 {
			// If NoUI is selected and there isn't a file argument
			log.Fatal("error no UI mode should get one argument as the filename")
		}

		if !noUIFlag && len(args) > 0 {
			// If UI mode is selected, but there are more than one arguments
			log.Fatal("error unnecessary arguments passed to remove")
		}

		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		removeFilePath := args[0]

		handler.Remove(handlers.RemoveHandlerParams{
			RepositoryFolderPath: repoFolderPath,
			FilePathToRemove:     removeFilePath,
		})

	},
}

func init() {
	removeCommand.PersistentFlags().Bool("nui", true, "Prevent UI from popping up")

	rootCmd.AddCommand(removeCommand)
}
