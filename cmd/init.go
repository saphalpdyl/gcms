package cmd

import (
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/saphalpdyl/gcms/internals/defaults"
	"github.com/saphalpdyl/gcms/internals/repository/github"
	"github.com/saphalpdyl/gcms/internals/styles"
	"github.com/saphalpdyl/gcms/internals/validator"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
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

			// Create the remote repository
			var repoNameAnswer string
			fmt.Print(styles.RenderBold("Name of repo (default: gcms) - "))
			fmt.Scan(&repoNameAnswer)

			if repoNameAnswer == "" {
				repoNameAnswer = "gcms"
			}

			response, err := github.CreateNewRepository(repoNameAnswer)
			if err != nil {
				fmt.Print(
					styles.RenderDiff(
						"Failed to create remote repository: Already Exists\n",
						false,
						"",
					),
				)

				os.RemoveAll(repoFolderPath)
				viper.Set(defaults.ConfigGithubRemoteURL, defaults.MISSING_VALUE)
				viper.Set(defaults.ConfigGithubRemoteFullName, defaults.MISSING_VALUE)
				viper.Set(defaults.ConfigGithubRemoteRepoName, defaults.MISSING_VALUE)
				viper.Set(defaults.ConfigGithubRemoteUserName, defaults.MISSING_VALUE)
			}

			// Set remote in config as responseURL
			viper.Set(defaults.ConfigGithubRemoteURL, response.URL)
			viper.Set(defaults.ConfigGithubRemoteFullName, response.RepositoryFullName)
			viper.Set(defaults.ConfigGithubRemoteRepoName, response.RepositoryName)
			viper.Set(defaults.ConfigGithubRemoteUserName, response.RepositoryOwner.RepositoryOwnerName)
			viper.WriteConfig()
			return
		}

	},
}

func init() {
	initCommmand.PersistentFlags().BoolP("empty", "e", false, "Create a empty repository")
	initCommmand.PersistentFlags().String("from", "", "Fork an existing repository")

	rootCmd.AddCommand(initCommmand)
}
