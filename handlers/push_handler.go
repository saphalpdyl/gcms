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

	if params.HasMetaData {
		var err error

		metaDataKeyValuePairs, err = helpers.ParseStringFromSSV(params.Metadata)

		if err != nil {
			log.Fatal("fatal couldn't parse metadata")
		}
	}

	// Extract file name
	baseFileName := filepath.Base(params.Filepath)
	newPathFile := filepath.Join(params.RepositoryFilePath, baseFileName)
	absoluteFilePath, err := filepath.Abs(params.Filepath)

	if err != nil {
		log.Fatal("fatal couldn't find absolute path of the file")
	}

	originalFile, err := os.Open(absoluteFilePath)
	if err != nil {
		log.Fatal("fatal couldn't open file")
	}

	defer originalFile.Close()

	moveFile, err := os.Create(newPathFile)
	if err != nil {
		log.Fatal("fatal couldn't move file to repository: ", err)
	}

	defer moveFile.Close()

	_, err = io.Copy(moveFile, originalFile)

	if err != nil {
		panic(err)
	}

	fmt.Println(metaDataKeyValuePairs)

}
