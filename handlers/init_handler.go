package handlers

import (
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/saphalpdyl/gcms/helpers"
	"github.com/saphalpdyl/gcms/internals/defaults"
	"github.com/spf13/viper"
)

type InitHandlerParams struct {
	FromEmpty            bool
	FromRemote           bool
	RepositoryFolderPath string
}

func (h *Handler) Init(params InitHandlerParams) {
	if params.FromEmpty {
		// Initialize the repository
		createRepoCmd := exec.Command("git", "init", params.RepositoryFolderPath)
		_, err := createRepoCmd.CombinedOutput()

		if err != nil {
			log.Fatalf("fatal cannot create empty repository %v", err)
		}

		// Create the remote repository
		var repoNameAnswer string
		fmt.Print(helpers.RenderBold("Name of repo (default: gcms) - "))
		fmt.Scan(&repoNameAnswer)

		if repoNameAnswer == "" {
			repoNameAnswer = "gcms"
		}

		response, err := h.githubService.CreateNewRepository(repoNameAnswer)
		if err != nil {
			fmt.Print(
				helpers.RenderDiff(
					"Failed to create remote repository: Already Exists\n",
					false,
					"",
				),
			)

			os.RemoveAll(params.RepositoryFolderPath)
			viper.Set(defaults.ConfigGithubRemoteURL, defaults.MISSING_VALUE)
			viper.Set(defaults.ConfigGithubRemoteFullName, defaults.MISSING_VALUE)
			viper.Set(defaults.ConfigGithubRemoteRepoName, defaults.MISSING_VALUE)
			viper.Set(defaults.ConfigGithubRemoteUserName, defaults.MISSING_VALUE)

			return
		}

		// Set remote in config as responseURL
		viper.Set(defaults.ConfigGithubRemoteURL, response.URL)
		viper.Set(defaults.ConfigGithubRemoteFullName, response.RepositoryFullName)
		viper.Set(defaults.ConfigGithubRemoteRepoName, response.RepositoryName)
		viper.Set(defaults.ConfigGithubRemoteUserName, response.RepositoryOwner.RepositoryOwnerName)
		viper.WriteConfig()

		// Add the remote repository to the local
		h.githubService.LinkLocalToRemote(params.RepositoryFolderPath, response.RepositoryName, response.RepositoryOwner.RepositoryOwnerName)

		return
	}

}
