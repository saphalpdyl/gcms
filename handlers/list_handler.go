package handlers

import (
	"fmt"
	"log"
	"os"

	"github.com/saphalpdyl/gcms/utils"
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
	entries, err := os.ReadDir(params.RepositoryFolderPath)

	if err != nil {
		log.Fatal("fatal cannot read repository path: ", err)
	}

	for _, e := range entries {
		if utils.StringInStringList(e.Name(), READDIR_EXCLUDE) {
			// Ignore excluded folders
			continue
		}

		fmt.Printf("- %s\n", e.Name())
	}

}
