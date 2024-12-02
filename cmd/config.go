package cmd

import (
	"fmt"
	"log"

	"github.com/saphalpdyl/gcms/handlers"
	"github.com/saphalpdyl/gcms/internals/defaults"
	"github.com/saphalpdyl/gcms/utils"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	k_allowedConfigSetValues = []string{
		defaults.ConfigGithubPATToken,
		defaults.ConfigGithubRemoteURL,
	}
)

var configCommand = &cobra.Command{
	Use:   "config [set|get] <config_key>",
	Short: "Subcommands related to GCMS configs",
}

var configSetCommand = &cobra.Command{
	Use: "set <key> <value>",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) != 2 {
			return fmt.Errorf("set missing required number of arguments: expected 2 got %d", len(args))
		}

		k := args[0]

		if !utils.StringInStringList(k, k_allowedConfigSetValues) {
			return fmt.Errorf("set invalid key as arguments; accepted keys are %s", utils.GenerateDSVFromStringList(k_allowedConfigSetValues))
		}

		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		err := viper.ReadInConfig()
		if err != nil {
			log.Fatalf("fatal reading config %v", err)
		}

		k, v := args[0], args[1]

		handler.ConfigSet(handlers.ConfigSetHandlerParams{
			K: k,
			V: v,
		})
	},
}

var configGetCommand = &cobra.Command{
	Use:  "get <key>",
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		err := viper.ReadInConfig()
		if err != nil {
			log.Fatalf("fatal reading config %v", err)
		}

		k := args[0]

		handler.ConfigGet(handlers.ConfigGetHandlerParams{
			K: k,
		})
	},
}

func init() {
	configCommand.AddCommand(configSetCommand)
	configCommand.AddCommand(configGetCommand)
}
