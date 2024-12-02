package cmd

import (
	"github.com/saphalpdyl/gcms/handlers"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var doctorCommmand = &cobra.Command{
	Use:   "doctor",
	Short: "Runs a status check on everything",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		handler.Doctor(handlers.DoctorHandlerParams{
			RootFolderPath:       homePath,
			RepositoryExists:     repositoryExists,
			RepositoryFolderPath: repoFolderPath,
			Viper:                viper.GetViper(),
		})
	},
}

func init() {
	rootCmd.AddCommand(doctorCommmand)
}
