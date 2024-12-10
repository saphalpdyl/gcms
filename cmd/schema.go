package cmd

import (
	"fmt"

	"github.com/saphalpdyl/gcms/handlers"
	"github.com/saphalpdyl/gcms/helpers"
	"github.com/spf13/cobra"
)

var schemaCommand = &cobra.Command{
	Use: "schema <sub-command> [flags] [args]",
}

var schemaCreateCommand = &cobra.Command{
	Use:   "create <group-name> <semicolon-separated-KVs>",
	Short: "Create a form schema for a group",
	Args: func(cmd *cobra.Command, args []string) error {
		// Validation of SSV
		if len(args) != 2 {
			return fmt.Errorf("%v", "Not enough arguments")
		}

		_, err := helpers.ParseStringFromSSV(args[1])
		if err != nil {
			return fmt.Errorf("%v", "invalid string seperated values passed in arguments")
		}

		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		handler.SchemaCreateNewGroup(handlers.SchemaCreateNewHandlerParams{})
	},
}

func init() {
	schemaCommand.AddCommand(schemaCreateCommand)

	rootCmd.AddCommand(schemaCommand)
}
