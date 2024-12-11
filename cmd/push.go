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
		groupFlagValue, _ := cmd.Flags().GetString("group")
		nuiFlagValue, _ := cmd.Flags().GetBool("nui")

		handler.Push(handlers.PushHandlerParams{
			Metadata:           metadataFlagValue,
			HasMetaData:        metadataFlagValue != "",
			Filepath:           args[0],
			HasGroup:           groupFlagValue != "",
			Group:              groupFlagValue,
			RepositoryFilePath: repoFolderPath,
			NoUIMode:           nuiFlagValue,
		})
	},
}

func init() {
	pushCommand.PersistentFlags().StringP("metadata", "m", "", "Provide a metadata to be stored in metadata.json to be picked up by other applications")
	pushCommand.PersistentFlags().StringP("group", "g", "", "Groups the file by a keyword in both the filename and metaname. File name will appear as example.group.html.")
	pushCommand.PersistentFlags().Bool("nui", false, "Run in CLI mode: GUI will not show up")

	rootCmd.AddCommand(pushCommand)
}
