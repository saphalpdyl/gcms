package helpers

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/saphalpdyl/gcms/internals/models"
)

func CreateEmptyJsonFile(folderPath string) error {
	file, err := os.Create("schema.json")
	if err != nil {
		return fmt.Errorf("%v", err)
	}

	defer file.Close()

	jsonData := make(models.SchemaMap)

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", " ")

	if err = encoder.Encode(jsonData); err != nil {
		return fmt.Errorf("couldn't encode json")
	}

	return nil
}
