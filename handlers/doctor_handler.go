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
	// Check for the .gcms folder
	// Check for the .gcms.config.yaml file
	// Check for PAT token
	// Check for the repo folder
	// Check for git in repo folder
	// Check for remote information in setings

	viper := params.Viper
	var statusMessages []string

	rootPathExists := utils.PathExists(params.RootFolderPath)

	if rootPathExists {
		statusMessages = append(statusMessages, helpers.RenderDiff("✅ .gcms folder exists", true, ""))
	} else {
		statusMessages = append(statusMessages, helpers.RenderDiff("❌ .gcms folder not found in typical folders.", false, ""))
	}

	// Proceed only if root path exists
	if rootPathExists {
		configPathExists := utils.PathExists(filepath.Join(params.RootFolderPath, "gcms.config.yml"))

		if configPathExists {
			statusMessages = append(statusMessages, helpers.RenderDiff("✅ configuration exists", true, ""))
		} else {
			statusMessages = append(statusMessages, helpers.RenderDiff("❌ .gcms.config.yml is missing. Try deleting the .gcms folder and run the CLI again.", false, ""))
		}

		// Proceed only if config exists
		if configPathExists {

			viper.SetConfigType("yaml")
			viper.SetConfigFile(filepath.Join(params.RootFolderPath, "gcms.config.yml"))
			viper.ReadInConfig()

			if err := viper.ReadInConfig(); err != nil {
				statusMessages = append(statusMessages, helpers.RenderDiff("❌ Couldn't read configuration file. The file might be corrupted.", false, ""))
			} else {
				// Proceed only if there was no error while reading the config file

				statusMessages = append(statusMessages, helpers.RenderDiff("✅ Configuration file ok", true, ""))
				patTokenMissing := viper.GetString(defaults.ConfigGithubPATToken) == defaults.MISSING_VALUE

				if patTokenMissing {
					statusMessages = append(statusMessages, helpers.RenderDiff("❌ Personal Access Token is missing.", false, ""))
				} else {
					statusMessages = append(statusMessages, helpers.RenderDiff("✅ Personal Access Token found in configuration", true, ""))
				}

				// Proceed even if pat might be missing
				// Check for local repository
				if !params.RepositoryExists {
					statusMessages = append(statusMessages, helpers.RenderDiff("❌ Local Repository is missing. Create a new one using gcms init --empty or gcms init --from <remote-url>", false, ""))
				} else {
					statusMessages = append(statusMessages, helpers.RenderDiff("✅ Local Repository was found", true, ""))

					// Proceed only if git folder is not missing
					// Check for git initialization status in the folder
					_, err := git.PlainOpen(params.RepositoryFolderPath)

					if err != nil {
						// Repository opening failed
						statusMessages = append(statusMessages, helpers.RenderDiff("❌ Git was not found in the local repository path. Please reinitialize GCMS.", false, ""))
					} else {
						statusMessages = append(statusMessages, helpers.RenderDiff("✅ Git was found in the local repository path.", true, ""))

						// Proceed only if repository was found
						remoteURL := viper.GetString(defaults.ConfigGithubRemoteURL)
						remoteSettingsMissing := remoteURL == defaults.MISSING_VALUE

						if remoteSettingsMissing {
							statusMessages = append(statusMessages, helpers.RenderDiff("❌ GCMS is not connected with remote.", false, ""))
						} else {
							statusMessages = append(statusMessages, helpers.RenderDiff(fmt.Sprintf("✅ GCMS is connected with remote at %v", remoteURL), true, ""))
						}
					}
				}
			}
		}
	}

	for _, msg := range statusMessages {
		fmt.Println(msg)
	}
}
