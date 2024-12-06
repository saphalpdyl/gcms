package handlers

import (
	"fmt"

	"github.com/saphalpdyl/gcms/helpers"
)

type ListHandlerParams struct {
	RepositoryFolderPath string
}

var (
	READDIR_EXCLUDE = []string{
		".git",
		"metadata.json",
	}
)

func (h *Handler) List(params ListHandlerParams) {
	files := helpers.GetFilesFromRepositoryDir(params.RepositoryFolderPath, READDIR_EXCLUDE)

	for _, file := range files {
		fmt.Printf("- %s\n", file.Name())
	}
}
