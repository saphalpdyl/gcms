package schema

import (
	"os"
	"path"

	"github.com/saphalpdyl/gcms/internals/models"
)

func (s *SchemaRepository) LoadSchema() error {
	configAbsolutePath := path.Join(s.schemaFolderPath, s.schemaFileName)

	data, err := os.ReadFile(configAbsolutePath)

	if err != nil {
		// Handle errors when opening the file (e.g., permission denied)
		return err
	}

	var schema models.SchemaMap

	err = s.serializer.Deserialize(data, &schema)

	if err != nil {
		// Handle errors during deserialization
		return err
	}

	s.data = &schema

	return nil
}
