package cmd

import (
	"fmt"

	utils "github.com/saphalpdyl/gcms/internals"
	"github.com/spf13/cobra"
)

var (
	k_allowedConfigSetValues = []string{"github.api_key"}
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
	},
}

func init() {
	configCommand.AddCommand(configSetCommand)
}
