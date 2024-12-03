package handlers

import (
	"fmt"
	"log"
	"path/filepath"
	"time"

	"github.com/saphalpdyl/gcms/helpers"
	"github.com/saphalpdyl/gcms/internals/models"
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
	metaDataKeyValuePairs := make(map[string]string)

	// Validate path exists
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

	groupName := params.Group
	if !params.HasGroup {
		// Assign `global` group to the file
		groupName = "global"
	}

	fileExistsInRepo := utils.PathExists(newPathFile)

	// Move the file to the repository
	if err := utils.CopyToPath(absoluteFilePath, newPathFile); err != nil {
		log.Fatal("fatal couldn't copy file to repository: ", err)
	}

	metadata, err := helpers.ReadMetadata(params.RepositoryFilePath)
	if err != nil {
		log.Fatal(err)
	}

	var group *models.GroupData

	if !helpers.MetadataGroupExists(metadata, groupName) {
		// Group doesn't exists
		// Creating a new group and adding it to the final metadata
		group = &models.GroupData{
			Group: groupName,
			Files: make([]*models.FileMetadata, 0),
		}

		metadata.Data = append(metadata.Data, group)
	} else {
		group = helpers.MetadataGetGroup(metadata, groupName)
	}

	// Getting the filepath relative to the repository
	relativePath, err := filepath.Rel(params.RepositoryFilePath, newPathFile)

	if err != nil {
		log.Fatal("fatal couldn't calculate the relative path of the file")
	}

	if !fileExistsInRepo {
		// Add file to the group object
		group.Files = append(group.Files, &models.FileMetadata{
			FilePath: relativePath,
			Metadata: metaDataKeyValuePairs,
		})
	}

	// Updated the last updated property
	metadata.LastUpdated = time.Now().UnixMilli()

	// Save metadata
	err = helpers.WriteMetadata(params.RepositoryFilePath, metadata)

	if err != nil {
		log.Fatalf("fatal %v", err)
	}

	err = helpers.CommitCurrentChanges(
		params.RepositoryFilePath,
		fmt.Sprintf("last updated at %d", time.Now().UnixMilli()),
	)

	if err != nil {
		log.Fatal(err)
	}

	h.githubService.UpdateRepository()

	fmt.Print(helpers.RenderBold("File added successfully..."))
}
