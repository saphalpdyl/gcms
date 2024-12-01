package handlers

import (
	"fmt"

	"github.com/saphalpdyl/gcms/helpers"
	"github.com/spf13/viper"
)

type ConfigGetHandlerParams struct {
	k string
}

func (h *Handler) ConfigGet(params ConfigGetHandlerParams) {
	k := params.k
	value := viper.GetString(k)

	if value == "" {
		fmt.Printf("Key %s doesn't exist in configuration\n", helpers.RenderBold(k))
		return
	}

	fmt.Printf("Result -> %s: %s\n", helpers.RenderBold(k), helpers.RenderBold(value))
}
