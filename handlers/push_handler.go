package handlers

import (
	"fmt"
	"log"

	"github.com/saphalpdyl/gcms/helpers"
	"github.com/saphalpdyl/gcms/utils"
)

type PushHandlerParams struct {
	HasMetaData bool
	Metadata    string
	Filepath    string
	HasGroup    bool
	Group       string
}

func (h *Handler) Push(params PushHandlerParams) {
	var metaDataKeyValuePairs [][]string

	if !utils.PathExists(params.Filepath) {
		fmt.Printf("Error: invalid filepath %s. Item now Found\n", params.Filepath)
		return
	}

	if params.HasMetaData {
		var err error

		metaDataKeyValuePairs, err = helpers.ParseStringFromSSV(params.Metadata)

		if err != nil {
			log.Fatal("fatal couldn't parse metadata")
		}
	}

	fmt.Println(metaDataKeyValuePairs)

}
