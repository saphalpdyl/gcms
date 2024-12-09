package schema

import "github.com/saphalpdyl/gcms/internals/models"

type ISchemaRepository interface {
	InitializeEmptySchema()
	LoadSchema()

	GetGroupSchema(string) (*models.Schema, error)
	UpdateGroupSchema(*models.Schema, string) error
}
