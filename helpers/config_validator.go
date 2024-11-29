package helpers

import (
	"fmt"
	"log"

	"github.com/saphalpdyl/gcms/internals/defaults"
	"github.com/spf13/viper"
)

func ValidatePATExists() {
	// Function that makes sure that the PAT is not missing
	viper.ReadInConfig()

	patValue := viper.GetString(defaults.ConfigGithubPATToken)

	if patValue == defaults.MISSING_VALUE {
		log.Fatalf(
			"fatal %s",
			RenderDiff(
				fmt.Sprintf("GitHub Personal Access Token is not configured. Configure using gcms config set %s <token>", defaults.ConfigGithubPATToken),
				false,
				"",
			),
		)
	}
}

func ValidateRemoteExists() {
	// Function that makes sure that the PAT is not missing
	viper.ReadInConfig()

	patValue := viper.GetString(defaults.ConfigGithubPATToken)

	if patValue == defaults.MISSING_VALUE {
		log.Fatalf(
			"fatal %s",
			RenderDiff(
				fmt.Sprintf("GitHub Remote is not configured. Configure using gcms config set %s <remote_link>", defaults.ConfigGithubRemoteURL),
				false,
				"",
			),
		)
	}
}
