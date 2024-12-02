package handlers

import (
	"fmt"
	"path/filepath"

	"github.com/go-git/go-git/v5"
	"github.com/saphalpdyl/gcms/helpers"
	"github.com/saphalpdyl/gcms/internals/defaults"
	"github.com/saphalpdyl/gcms/utils"
	"github.com/spf13/viper"
)

type DoctorHandlerParams struct {
	RootFolderPath       string
	RepositoryFolderPath string
	RepositoryExists     bool
	Viper                *viper.Viper
}

func (h *Handler) Doctor(params DoctorHandlerParams) {
	viper := params.Viper
	rootPathExists := utils.PathExists(params.RootFolderPath)

	fmt.Println(helpers.RenderBold("GCMS Doctor"))
	fmt.Println(helpers.RenderBold("==========="))
	fmt.Print("\nRunning Diagnostics...\n\n")

	if rootPathExists {
		fmt.Println(helpers.RenderDoctorResult(".gcms folder exists", true, ""))
	} else {
		fmt.Println(helpers.RenderDoctorResult(".gcms folder not found in typical folders", false, "Try reinitializing GCMS using 'gcms init'"))
	}

	// Proceed only if root path exists
	if rootPathExists {
		configPathExists := utils.PathExists(filepath.Join(params.RootFolderPath, "gcms.config.yml"))

		if configPathExists {
			fmt.Println(helpers.RenderDoctorResult("Configuration exists", true, ""))
		} else {
			fmt.Println(helpers.RenderDoctorResult("Configuration file missing", false, "Try deleting the .gcms folder and run the CLI again"))
		}

		// Proceed only if config exists
		if configPathExists {
			viper.SetConfigType("yaml")
			viper.SetConfigFile(filepath.Join(params.RootFolderPath, "gcms.config.yml"))
			viper.ReadInConfig()

			if err := viper.ReadInConfig(); err != nil {
				fmt.Println(helpers.RenderDoctorResult("Configuration file corrupted", false, "Try deleting the configuration file and reinitializing GCMS"))
			} else {
				fmt.Println(helpers.RenderDoctorResult("Configuration file verified", true, ""))
				patTokenMissing := viper.GetString(defaults.ConfigGithubPATToken) == defaults.MISSING_VALUE

				if patTokenMissing {
					fmt.Println(helpers.RenderDoctorResult("Personal Access Token missing", false, "Cannot proceed without Personal Access Token configured in settings. Visit https://docs.github.com/en/authentication/keeping-your-account-and-data-secure/managing-your-personal-access-tokens to learn more. Set your GitHub PAT using 'gcms config set github.pat <token>'"))
				} else {
					fmt.Println(helpers.RenderDoctorResult("Personal Access Token verified", true, ""))
				}

				// Proceed even if pat might be missing
				// Check for local repository
				if !params.RepositoryExists {
					fmt.Println(helpers.RenderDoctorResult("Local Repository missing", false, "Create a new one using 'gcms init --empty' or 'gcms init --from <remote-url>'"))
				} else {
					fmt.Println(helpers.RenderDoctorResult("Local Repository found", true, ""))

					// Proceed only if git folder is not missing
					// Check for git initialization status in the folder
					_, err := git.PlainOpen(params.RepositoryFolderPath)

					if err != nil {
						fmt.Println(helpers.RenderDoctorResult("Git not initialized in repository", false, "Please reinitialize GCMS"))
					} else {
						fmt.Println(helpers.RenderDoctorResult("Git initialized in repository", true, ""))

						// Proceed only if repository was found
						remoteURL := viper.GetString(defaults.ConfigGithubRemoteURL)
						remoteSettingsMissing := remoteURL == defaults.MISSING_VALUE

						if remoteSettingsMissing {
							fmt.Println(helpers.RenderDoctorResult("Remote connection missing", false, "Connect to remote using 'gcms remote set <url>'"))
						} else {
							fmt.Println(helpers.RenderDoctorResult(fmt.Sprintf("Connected to remote at %v", remoteURL), true, ""))
						}
					}
				}
			}
		}
	}
}
