package cmd

import (
	"fmt"
	"log"

	"github.com/go-git/go-git/v5"
	"github.com/saphalpdyl/gcms/internals/defaults"
	"github.com/saphalpdyl/gcms/internals/repository/github"
	"github.com/saphalpdyl/gcms/internals/styles"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var detachCommand = &cobra.Command{
	Use:   "detach [flags]",
	Short: "Detach the local repository with the remote",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		isHardDetachFlagValue, _ := cmd.Flags().GetBool("hard")

		// Open the repository
		r, err := git.PlainOpen(repoFolderPath)
		if err != nil {
			log.Fatalf("fatal %s", styles.RenderDiff(
				"couldn't open the local git repository",
				false,
				"",
			))
		}

		configRemoteOwnerName := viper.GetString(defaults.ConfigGithubRemoteUserName)
		configRemoteRepositoryName := viper.GetString(defaults.ConfigGithubRemoteRepoName)

		// HARD DETACH SECTION
		if isHardDetachFlagValue {

			url := fmt.Sprintf("https://api.github.com/repos/%s/%s", configRemoteOwnerName, configRemoteRepositoryName)

			var deleteConfirmationAnswer string

			// Confirmation message
			fmt.Printf(
				"\n%s repository with link: %s\n Are you sure you want to proceed? [y/N]",
				styles.RenderDiff("Deleting", false, ""),
				styles.RenderBold(url),
			)
			fmt.Scan(&deleteConfirmationAnswer)

			if deleteConfirmationAnswer != "y" && deleteConfirmationAnswer != "Y" {
				return
			}

			err = github.DeleteRepository(url)

			if err != nil {
				log.Fatal("fatal couldn't delete the remote repository: ", err)
			}

			fmt.Println("\nSuccessfully delete remote...")
		}

		err = r.DeleteRemote("origin")
		if err != nil {
			fmt.Print("fatal couldn't remove remote origin from the local repository: ", err)
		}

		// Remove the configuration remote
		viper.Set(defaults.ConfigGithubRemoteURL, defaults.MISSING_VALUE)
		viper.Set(defaults.ConfigGithubRemoteFullName, defaults.MISSING_VALUE)
		viper.Set(defaults.ConfigGithubRemoteRepoName, defaults.MISSING_VALUE)
		viper.Set(defaults.ConfigGithubRemoteUserName, defaults.MISSING_VALUE)
		viper.WriteConfig()
	},
}

func init() {
	detachCommand.PersistentFlags().Bool("hard", false, "Remove the remote in git and delete the remote repository if given perms")

	rootCmd.AddCommand(detachCommand)
}
