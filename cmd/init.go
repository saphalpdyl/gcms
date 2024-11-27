package cmd

import (
	"fmt"
	"log"
	"os/exec"

	"github.com/saphalpdyl/gcms/internals/styles"
	"github.com/saphalpdyl/gcms/internals/validator"
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

		validator.ValidatePATExists()

		// Check for previous initialization
		if repositoryExists {
			log.Fatalf("fatal %s", styles.RenderDiff("Repository already exists. Remove the existing repository first.", false, ""))
			return
		}

		if emptyFlag {
			// Initialize the repository

			createRepoCmd := exec.Command("git", "init", repoFolderPath)
			_, err := createRepoCmd.CombinedOutput()

			if err != nil {
				log.Fatalf("fatal cannot create empty repository %v", err)
			}

			return
		}

	},
}

func init() {
	initCommmand.PersistentFlags().BoolP("empty", "e", false, "Create a empty repository")
	initCommmand.PersistentFlags().String("from", "", "Fork an existing repository")

	rootCmd.AddCommand(initCommmand)
}
