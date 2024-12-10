package schema

import (
	"github.com/saphalpdyl/gcms/internals/models"
	"github.com/saphalpdyl/gcms/internals/serializers"
)

type SchemaRepository struct {
	schemaFolderPath string
	schemaFileName   string
	data             *models.SchemaMap

	serializer serializers.ISerializer[models.SchemaMap]
}

func New(
	schemaFolderPath string,
	schemaFileName string,
	data *models.SchemaMap,
) ISchemaRepository {
	return &SchemaRepository{
		schemaFolderPath: schemaFolderPath,
		schemaFileName:   schemaFileName,
		data:             data,
	}
}
