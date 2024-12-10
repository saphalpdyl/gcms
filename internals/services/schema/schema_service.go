package schema_service

import "github.com/saphalpdyl/gcms/internals/models"

type ISchemaService interface {
	InitializeEmptySchema() error
	LoadSchema() error

	GetGroupSchema(string) (*models.Schema, error)
	UpdateGroupSchema(*models.Schema, string) error
}
