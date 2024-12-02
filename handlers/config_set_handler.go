package handlers

import (
	"fmt"

	"github.com/saphalpdyl/gcms/helpers"
	"github.com/saphalpdyl/gcms/internals/defaults"
	"github.com/spf13/viper"
)

type ConfigSetHandlerParams struct {
	K string
	V string
}

func (h *Handler) ConfigSet(params ConfigSetHandlerParams) {
	k := params.K
	v := params.V

	previousValue := viper.GetString(k)

	if previousValue != defaults.MISSING_VALUE {
		// Ask for confirmation

		confirmationMessage := fmt.Sprintf(
			"Value exists: \n\t%s\n\t%s\n Are you sure you want to replace it? [y/N]",
			helpers.RenderDiff(previousValue, false, "- "),
			helpers.RenderDiff(v, true, "+ "),
		)

		confirmationRenderedMesage := helpers.RenderBold(confirmationMessage)
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
}
