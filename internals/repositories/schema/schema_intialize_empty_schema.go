package schema

import "github.com/saphalpdyl/gcms/helpers"

func (s *SchemaRepository) InitializeEmptySchema() error {
	// Create an empty JSON file
	return helpers.CreateEmptyJsonFile(s.schemaFolderPath)
}
