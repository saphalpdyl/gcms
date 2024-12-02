package cmd

import (
	"github.com/saphalpdyl/gcms/handlers"
	"github.com/spf13/cobra"
)

var updateCommand = &cobra.Command{
	Use:   "update",
	Short: "Pull from the remote repository",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		handler.Update(handlers.UpdateHandlerParams{})
	},
}

func init() {
	rootCmd.AddCommand(updateCommand)
}
