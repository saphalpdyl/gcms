package helpers

import (
	"io/fs"
	"log"
	"os"

	"github.com/saphalpdyl/gcms/utils"
)

func GetFilesFromRepositoryDir(repositoryFolderPath string, exclusionList []string) []fs.DirEntry {
	entries, err := os.ReadDir(repositoryFolderPath)
	finalFilesList := []fs.DirEntry{}

	if err != nil {
		log.Fatal("fatal cannot read repository path: ", err)
	}

	for _, e := range entries {
		if utils.StringInStringList(e.Name(), exclusionList) {
			// Ignore excluded folders
			continue
		}

		finalFilesList = append(finalFilesList, e)
	}

	return finalFilesList
}
