package handlers

import (
	"fmt"
	"log"

	"github.com/go-git/go-git/v5"
	"github.com/saphalpdyl/gcms/helpers"
	"github.com/saphalpdyl/gcms/internals/defaults"
	"github.com/spf13/viper"
)

type DetachHandlerParams struct {
	Viper                *viper.Viper
	RepositoryFolderPath string
	IsHardDetach         bool
}

func (h *Handler) Detach(params DetachHandlerParams) {
	repoFolderPath := params.RepositoryFolderPath
	viper := params.Viper

	r, err := git.PlainOpen(repoFolderPath)
	if err != nil {
		log.Fatalf("fatal %s", helpers.RenderDiff(
			"couldn't open the local git repository",
			false,
			"",
		))
	}

	configRemoteOwnerName := viper.GetString(defaults.ConfigGithubRemoteUserName)
	configRemoteRepositoryName := viper.GetString(defaults.ConfigGithubRemoteRepoName)

	// HARD DETACH SECTION
	if params.IsHardDetach {

		url := fmt.Sprintf("https://api.github.com/repos/%s/%s", configRemoteOwnerName, configRemoteRepositoryName)

		var deleteConfirmationAnswer string

		// Confirmation message
		fmt.Printf(
			"\n%s repository with link: %s\n Are you sure you want to proceed? [y/N]",
			helpers.RenderDiff("Deleting", false, ""),
			helpers.RenderBold(url),
		)
		fmt.Scan(&deleteConfirmationAnswer)

		if deleteConfirmationAnswer != "y" && deleteConfirmationAnswer != "Y" {
			return
		}

		err = h.githubRepostiory.DeleteRepository(url)

		if err != nil {
			log.Fatal("fatal couldn't delete the remote repository: ", err)
		}

		fmt.Println("\nSuccessfully deleted remote...")
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
}
