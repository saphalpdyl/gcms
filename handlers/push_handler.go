package handlers

import (
	"fmt"
	"log"

	"github.com/saphalpdyl/gcms/helpers"
)

type PushHandlerParams struct {
	HasMetaData bool
	Metadata    string
	Filepath    string
}

func (h *Handler) Push(params PushHandlerParams) {
	var metaDataKeyValuePairs [][]string

	if params.HasMetaData {
		var err error

		metaDataKeyValuePairs, err = helpers.ParseStringFromSSV(params.Metadata)

		if err != nil {
			log.Fatal("fatal couldn't parse metadata")
		}
	}

	fmt.Println(metaDataKeyValuePairs)

}
