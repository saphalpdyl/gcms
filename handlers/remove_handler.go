package handlers

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/saphalpdyl/gcms/helpers"
)

type RemoveHandlerParams struct {
	RepositoryFolderPath string
	FilePathToRemove     string
}

func (h *Handler) Remove(params RemoveHandlerParams) {
	metadata, err := helpers.ReadMetadata(params.FilePathToRemove)

	if err != nil {
		log.Fatal(err)
	}

	if !helpers.MetadataFilePathExists(metadata, params.FilePathToRemove) {
		log.Fatalf("fatal file doesn't exist: %s", params.FilePathToRemove)
	}

	helpers.MetadataRemoveFilePath(metadata, params.FilePathToRemove)

	err = helpers.WriteMetadata(params.RepositoryFolderPath, metadata)

	if err != nil {
		log.Fatal(err)
	}

	err = os.Remove(filepath.Join(params.RepositoryFolderPath, params.FilePathToRemove))

	if err != nil {
		log.Fatal("fatal couldn't delete file from the local repository. Metadata.json might be corrupted. Please manually delete the correspoding metadata and push to remote if it exists.")
	}

	err = helpers.CommitCurrentChanges(params.RepositoryFolderPath,
		fmt.Sprintf("removed %s", params.FilePathToRemove),
	)

	if err != nil {
		log.Fatal(err)
	}

	h.githubService.UpdateRepository()

	fmt.Println(helpers.RenderBold("Removed Successfully..."))
}
