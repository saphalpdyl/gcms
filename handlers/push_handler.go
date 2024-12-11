package handlers

import (
	"fmt"
	"log"
	"path/filepath"
	"time"

	"github.com/charmbracelet/huh"
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
	NoUIMode    bool

	RepositoryFilePath string
}

func handleUIPush(params PushHandlerParams, h *Handler) map[string]string {
	// Returns the metadata as SSV if valid
	schema, err := h.schemaRepository.GetGroupSchema(params.Group)

	if err != nil {
		log.Fatalf("fatal %v", err)
	}

	formTitleValuePtrPairs := make(map[string]*string)
	huhFields := make([]huh.Field, 0)

	for _, formItem := range schema.Schema {
		emptyValue := ""
		formTitleValuePtrPairs[formItem.Title] = &emptyValue

		answerPtr := formTitleValuePtrPairs[formItem.Title]

		var field huh.Field

		if formItem.ElementType == "INPUT" {
			field = huh.
				NewInput().
				Value(answerPtr).
				Title(formItem.Title)
		} else if formItem.ElementType == "TEXTAREA" {
			field = huh.
				NewText().
				Value(answerPtr).
				Title(formItem.Title).
				Lines(4)
		}

		huhFields = append(huhFields, field)
	}

	// Show charm cli UI here
	form := huh.NewForm(
		huh.NewGroup(
			huhFields...,
		),
	)

	err = form.Run()
	if err != nil {
		log.Fatal(err)
	}

	// Convert the pointers back to values map
	formTitleValuePairs := make(map[string]string)
	for k, v := range formTitleValuePtrPairs {
		formTitleValuePairs[k] = *v
	}

	fmt.Println(formTitleValuePairs)

	return make(map[string]string)
}

func (h *Handler) Push(params PushHandlerParams) {
	metaDataKeyValuePairs := make(map[string]string)

	// Validate path exists
	if !utils.PathExists(params.Filepath) {
		fmt.Printf("Error: invalid filepath %s. Item now Found\n", params.Filepath)
		return
	}

	if !params.NoUIMode && params.HasGroup {
		// If it isn't explicity mentioned not to show GUI
		// and the push is associated with a group
		metaDataKeyValuePairs = handleUIPush(params, h)
	} else {
		if params.HasMetaData {
			// Validate and compute metadata values
			var err error

			metaDataKeyValuePairs, err = helpers.ParseStringFromSSV(params.Metadata)

			if err != nil {
				log.Fatal("fatal couldn't parse metadata")
			}
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

	h.githubRepostiory.UpdateRepository()

	fmt.Print(helpers.RenderBold("File added successfully..."))
}
