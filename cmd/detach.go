package cmd

import (
	"github.com/saphalpdyl/gcms/handlers"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var detachCommand = &cobra.Command{
	Use:   "detach [flags]",
	Short: "Detach the local repository with the remote",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {

		isHardDetachFlagValue, _ := cmd.Flags().GetBool("hard")

		handler.Detach(handlers.DetachHandlerParams{
			Viper:                viper.GetViper(),
			RepositoryFolderPath: repoFolderPath,
			IsHardDetach:         isHardDetachFlagValue,
		})
	},
}

func init() {
	detachCommand.PersistentFlags().Bool("hard", false, "Remove the remote in git and delete the remote repository if given perms")

	rootCmd.AddCommand(detachCommand)
}
