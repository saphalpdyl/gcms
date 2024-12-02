package cmd

import (
	"fmt"
	"log"

	"github.com/saphalpdyl/gcms/helpers"
	"github.com/spf13/cobra"
)

var pushCommand = &cobra.Command{
	Use:   "push <filepath> [flags]",
	Short: "Push a new file into the system with metadata(-m) if required",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		metadataFlagValue, _ := cmd.Flags().GetString("metadata")

		var metaDataKeyValuePairs [][]string

		if metadataFlagValue != "" {
			var err error

			metaDataKeyValuePairs, err = helpers.ParseStringFromSSV(metadataFlagValue)

			if err != nil {
				log.Fatal("fatal couldn't parse metadata")
			}
		}

		fmt.Println(metaDataKeyValuePairs)
	},
}

func init() {
	pushCommand.PersistentFlags().StringP("metadata", "m", "", "Provide a metadata to be stored in metadata.json to be picked up by other applications")

	rootCmd.AddCommand(pushCommand)
}
