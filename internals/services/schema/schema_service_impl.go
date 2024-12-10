package schema_service

import "github.com/saphalpdyl/gcms/internals/repositories/schema"

type SchemaService struct {
	schemaRepository schema.ISchemaRepository
}

func NewSchemaService(schemaRepository schema.ISchemaRepository) SchemaService {
	return SchemaService{
		schemaRepository: schemaRepository,
	}
}
