package cmd

import (
	"fmt"
	"log"

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
			return fmt.Errorf("not enough arguments, got %d", len(args))
		}

		_, err := helpers.ParseStringFromSSV(args[1])
		if err != nil {
			return fmt.Errorf("%v", "invalid string seperated values passed in arguments")
		}

		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		groupName := args[0]
		if groupName == " " {
			log.Fatal("fatal invalid groupName: ", helpers.RenderBold(groupName))
		}

		handler.SchemaCreateNewGroup(handlers.SchemaCreateNewHandlerParams{
			GroupName: groupName,
			FormData:  args[1],
		})
	},
}

func init() {
	schemaCommand.AddCommand(schemaCreateCommand)

	rootCmd.AddCommand(schemaCommand)
}
