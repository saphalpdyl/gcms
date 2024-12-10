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

func NewRepository(
	schemaFolderPath string,
	schemaFileName string,
	serializer serializers.ISerializer[models.SchemaMap],
) ISchemaRepository {
	return &SchemaRepository{
		schemaFolderPath: schemaFolderPath,
		schemaFileName:   schemaFileName,
	}
}
