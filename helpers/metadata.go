package helpers

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/saphalpdyl/gcms/internals/models"
)

func ReadMetadata(repositoryFolderPath string) (*models.RootMetaData, error) {
	// Read from metadata.json
	jsonByte, err := os.ReadFile(
		filepath.Join(repositoryFolderPath, "metadata.json"),
	)

	if err != nil {
		return nil, fmt.Errorf("fatal couldn't read metadata.json %v", err)
	}

	var metadata models.RootMetaData

	if err = json.Unmarshal(jsonByte, &metadata); err != nil {
		return nil, fmt.Errorf("fatal couldn't unmarshal metadata.json ")
	}

	return &metadata, nil
}

func WriteMetadata(repositoryFolderPath string, metadata *models.RootMetaData) error {
	jsonByte, err := json.Marshal(metadata)
	if err != nil {
		return fmt.Errorf("couldn't stringify json into byte[]")
	}

	err = os.WriteFile(filepath.Join(repositoryFolderPath, "metadata.json"), jsonByte, os.ModePerm)

	if err != nil {
		return fmt.Errorf("couldn't write to metadata.json")
	}

	return nil
}

func MetadataGroupExists(metadata *models.RootMetaData, groupName string) bool {
	for _, groupItem := range metadata.Data {
		if groupItem.Group == groupName {
			return true
		}
	}

	return false
}

func MetadataGetGroup(metadata *models.RootMetaData, groupName string) *models.GroupData {
	for _, groupItem := range metadata.Data {
		if groupItem.Group == groupName {
			return groupItem
		}
	}

	return nil
}

func MetadataFilePathExists(metadata *models.RootMetaData, filePathToSearch string) bool {
	for _, groupItem := range metadata.Data {
		for _, fileItem := range groupItem.Files {
			if fileItem.FilePath == filePathToSearch {
				return true
			}
		}
	}

	return false
}

func MetadataRemoveFilePath(metadata *models.RootMetaData, filePathToSearch string) {

	newData := make([]*models.GroupData, 0)

	for _, groupItem := range metadata.Data {
		newGroup := &models.GroupData{
			Group: groupItem.Group,
			Files: make([]*models.FileMetadata, 0),
		}

		for _, fileItem := range groupItem.Files {
			if fileItem.FilePath == filePathToSearch {
				continue
			}

			newGroup.Files = append(newGroup.Files, &models.FileMetadata{
				FilePath: fileItem.FilePath,
				Metadata: fileItem.Metadata,
			})
		}

		newData = append(newData, newGroup)
	}

	metadata.Data = newData
}
