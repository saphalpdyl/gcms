package serializers

import "github.com/saphalpdyl/gcms/internals/models"

type SchemaSerializer struct{}

func New() ISerializer[models.SchemaMap] {
	return &SchemaSerializer{}
}

func (s *SchemaSerializer) Serialize(data string) (models.SchemaMap, error) {

	return nil, nil
}

func (s *SchemaSerializer) Deserialize(schemaMap models.SchemaMap) string {
	return ""
}
