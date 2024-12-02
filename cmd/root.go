/*
Copyright © 2024 Saphal Poudyal saphalpdyl@gmail.com

Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the “Software”), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED “AS IS”, WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
*/
package cmd

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/saphalpdyl/gcms/handlers"
	"github.com/saphalpdyl/gcms/internals/defaults"
	"github.com/saphalpdyl/gcms/internals/repositories/github"
	github_service "github.com/saphalpdyl/gcms/internals/services/github"
	"github.com/saphalpdyl/gcms/utils"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// Global variables
var (
	homePath         string
	repoFolderPath   string
	repositoryExists bool
)

// Repositories and Services
var (
	githubService    github_service.IGithubService
	githubRepository github.IGithubRepository
	handler          handlers.IHandler
)

var rootCmd = &cobra.Command{
	Use:   "gcms",
	Short: "A free Github-based Content Management System",
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	// Configuration
	rootCmd.AddCommand(configCommand)

	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Configuration directory and file
	homeDirectoryPath, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}

	// Create the base .gcms folder
	homePath = filepath.Join(homeDirectoryPath, ".gcms")
	repoFolderPath = filepath.Join(homePath, "repo")

	err = os.MkdirAll(homePath, os.ModePerm)
	if err != nil {
		log.Fatalf("cannot create .gcms directory in USER_HOME: %v", err)
	}

	// Setup Viper Configuration
	configFileName := "gcms.config.yml"

	viper.SetConfigType("yaml")
	viper.SetConfigFile(filepath.Join(homePath, configFileName))

	viper.SetDefault(defaults.ConfigGithubPATToken, defaults.MISSING_VALUE)
	viper.SetDefault(defaults.ConfigGithubRemoteURL, defaults.MISSING_VALUE)
	viper.SetDefault(defaults.ConfigGithubRemoteFullName, defaults.MISSING_VALUE)
	viper.SetDefault(defaults.ConfigGithubRemoteRepoName, defaults.MISSING_VALUE)
	viper.SetDefault(defaults.ConfigGithubRemoteUserName, defaults.MISSING_VALUE)

	err = viper.ReadInConfig()
	if err != nil {
		fmt.Println("Missing configuration files. Creating a new configuration...")

		viper.WriteConfig()
	}

	// Check for repository status
	repositoryExists = utils.PathExists(repoFolderPath)

	// Dependency injection
	if viper.GetString(defaults.ConfigGithubPATToken) == defaults.MISSING_VALUE {
		log.Fatal("fatal cannot proceed without Personal Access Token configured in settings. Visit https://docs.github.com/en/authentication/keeping-your-account-and-data-secure/managing-your-personal-access-tokens to learn more.")
	}

	githubRepository = github.NewRepository(viper.GetString(defaults.ConfigGithubPATToken), repoFolderPath)
	githubService = github_service.NewService(githubRepository)

	handler = handlers.NewHandler(githubService)
}
