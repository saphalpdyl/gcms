package cmd

import (
	"fmt"
	"os"

	"github.com/saphalpdyl/gcms/internals/styles"
	"github.com/spf13/cobra"
)

var removeCommand = &cobra.Command{
	Use:   "remove",
	Short: "Delete the local repository [WARNING: Cannot be reversed]",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		var deleteConfirmationAnswer string

		// Confirmation message
		fmt.Printf(
			"\n%s\n%s\n%s\n",
			styles.RenderDiff("Deleting the local repository", false, ""),
			"This action is irreversible.",
			styles.RenderBold("Are you sure you want to continue?[y/N] "),
		)
		fmt.Scan(&deleteConfirmationAnswer)

		if deleteConfirmationAnswer != "y" && deleteConfirmationAnswer != "Y" {
			return
		}

		os.RemoveAll(repoFolderPath)
	},
}

func init() {
	rootCmd.AddCommand(removeCommand)
}
