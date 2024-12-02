package cmd

import (
	"github.com/saphalpdyl/gcms/handlers"
	"github.com/spf13/cobra"
)

var pushCommand = &cobra.Command{
	Use:   "push <filepath> [flags]",
	Short: "Push a new file into the system with metadata(-m) if required",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		metadataFlagValue, _ := cmd.Flags().GetString("metadata")

		handler.Push(handlers.PushHandlerParams{
			Metadata:    metadataFlagValue,
			HasMetaData: metadataFlagValue != "",
			Filepath:    args[0],
		})
	},
}

func init() {
	pushCommand.PersistentFlags().StringP("metadata", "m", "", "Provide a metadata to be stored in metadata.json to be picked up by other applications")

	rootCmd.AddCommand(pushCommand)
}
