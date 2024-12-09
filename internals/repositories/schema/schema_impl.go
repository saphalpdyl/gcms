package schema

import "github.com/saphalpdyl/gcms/internals/models"

type SchemaRepository struct {
	schemaFolderPath string
	schemaFileName   string
	data             *models.SchemaMap
}

func (s *SchemaRepository) NewAndLoad() ISchemaRepository {
	return &SchemaRepository{}
}
