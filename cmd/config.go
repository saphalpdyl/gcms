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
	k_allowedConfigSetValues = []string{
		configGithubPATToken,
		configGithubRemote,
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

		previousValue := viper.GetString(k)

		if previousValue != MISSING_VALUE {
			// Ask for confirmation

			confirmationMessage := fmt.Sprintf(
				"Value exists: \n\t%s\n\t%s\n Are you sure you want to replace it? [y/N]",
				styles.RenderDiff(previousValue, false, "- "),
				styles.RenderDiff(v, true, "+ "),
			)

			confirmationRenderedMesage := styles.RenderBold(confirmationMessage)
			var confirmationAnswer string

			fmt.Print(confirmationRenderedMesage)
			fmt.Scan(&confirmationAnswer)

			// Exit if other character except Y or y
			if confirmationAnswer != "Y" && confirmationAnswer != "y" {
				return
			}
		}

		viper.Set(k, v)
		viper.WriteConfig()
		fmt.Println("Configuration Saved...")
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

		value := viper.GetString(k)

		if value == "" {
			fmt.Printf("Key %s doesn't exist in configuration\n", styles.RenderBold(k))
			return
		}

		fmt.Printf("Result -> %s: %s\n", styles.RenderBold(k), styles.RenderBold(value))
	},
}

func init() {
	configCommand.AddCommand(configSetCommand)
	configCommand.AddCommand(configGetCommand)
}
