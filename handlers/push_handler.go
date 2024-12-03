package handlers

import (
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"

	"github.com/saphalpdyl/gcms/helpers"
	"github.com/saphalpdyl/gcms/utils"
)

type PushHandlerParams struct {
	HasMetaData bool
	Metadata    string
	Filepath    string
	HasGroup    bool
	Group       string

	RepositoryFilePath string
}

func (h *Handler) Push(params PushHandlerParams) {
	var metaDataKeyValuePairs [][]string

	if !utils.PathExists(params.Filepath) {
		fmt.Printf("Error: invalid filepath %s. Item now Found\n", params.Filepath)
		return
	}

	// Validate and compute metadata values
	if params.HasMetaData {
		var err error

		metaDataKeyValuePairs, err = helpers.ParseStringFromSSV(params.Metadata)

		if err != nil {
			log.Fatal("fatal couldn't parse metadata")
		}
	}

	// Extract file names
	baseFileName := filepath.Base(params.Filepath)
	newPathFile := filepath.Join(params.RepositoryFilePath, baseFileName)
	absoluteFilePath, err := filepath.Abs(params.Filepath)

	if err != nil {
		log.Fatal("fatal couldn't find absolute path of the file")
	}

	if err := utils.CopyToPath(absoluteFilePath, newPathFile); err != nil {
		log.Fatal("fatal couldn't copy file to repository: ", err)
		return
	}

	fmt.Println(metaDataKeyValuePairs)

}
