package cmd

import (
	"fmt"

	"github.com/saphalpdyl/gcms/helpers"
	"github.com/spf13/cobra"
)

var versionCommand = &cobra.Command{
	Use:   "version",
	Short: "Gives the version of the installation",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(helpers.RenderBold("GCMS Version v1.2.1"))
	},
}

func init() {
	rootCmd.AddCommand(versionCommand)
}
