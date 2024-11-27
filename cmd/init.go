package cmd

import (
	"fmt"
	"log"
	"os/exec"
	"path/filepath"

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

		if emptyFlag {
			// Initialize the repository
			createRepoCmd := exec.Command("git", "init", filepath.Join(homePath, "repo"))
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
