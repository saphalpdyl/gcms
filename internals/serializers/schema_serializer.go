package serializers

import "github.com/saphalpdyl/gcms/internals/models"

type SchemaSerializer struct{}

func NewSchemaSerializer() ISerializer[models.SchemaMap] {
	return &SchemaSerializer{}
}

func (s *SchemaSerializer) Deserialize(data []byte, schemaMap *models.SchemaMap) error {
	//Convert from string to schemaMap

	return nil
}

func (s *SchemaSerializer) Serialize(schemaMap models.SchemaMap) []byte {
	// Convert from schemaMap to string

	return []byte{}
}
