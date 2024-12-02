package cmd

import (
	"fmt"
	"log"

	"github.com/saphalpdyl/gcms/handlers"
	"github.com/saphalpdyl/gcms/helpers"
	"github.com/spf13/cobra"
)

var initCommmand = &cobra.Command{
	Use:   "init [flags] [arguments]",
	Short: "Initialize the repository for GCMS",
	Args: func(cmd *cobra.Command, args []string) error {
		emptyFlag, _ := cmd.Flags().GetBool("empty")
		fromFlag, _ := cmd.Flags().GetString("from")

		if !emptyFlag && fromFlag == "" {
			return fmt.Errorf("init should have flags --empty or --from <link>")
		}

		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		emptyFlag, _ := cmd.Flags().GetBool("empty")
		// fromFlag, _ := cmd.Flags().GetString("from")

		helpers.ValidatePATExists()

		// Check for previous initialization
		if repositoryExists {
			log.Fatalf("fatal %s", helpers.RenderDiff("Repository already exists. Remove the existing repository first.", false, ""))
			return
		}

		handler.Init(handlers.InitHandlerParams{
			FromEmpty:            emptyFlag,
			FromRemote:           false,
			RepositoryFolderPath: repoFolderPath,
		})

	},
}

func init() {
	initCommmand.PersistentFlags().BoolP("empty", "e", false, "Create a empty repository")
	initCommmand.PersistentFlags().String("from", "", "Fork an existing repository")

	rootCmd.AddCommand(initCommmand)
}
