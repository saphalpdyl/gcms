package cmd

import (
	"fmt"
	"log"

	"github.com/saphalpdyl/gcms/internals/styles"
	utils "github.com/saphalpdyl/gcms/internals/utils"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
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
		err := viper.ReadInConfig()
		if err != nil {
			log.Fatalf("fatal reading config %v", err)
		}

		k, _ := args[0], args[1]

		previousValue := viper.GetString(k)

		if previousValue != MISSING_VALUE {
			// Ask for confirmation

			confirmationMessage := fmt.Sprintf(
				"Value exists: %s\n Are you sure you want to replace it? [y/N]", styles.RenderDanger(previousValue))
			confirmationRenderedMesage := styles.RenderBold(confirmationMessage)
			var confirmationAnswer string

			fmt.Print(confirmationRenderedMesage)
			fmt.Scan(&confirmationAnswer)
		}
	},
}

func init() {
	configCommand.AddCommand(configSetCommand)
}
