package schema

import "github.com/saphalpdyl/gcms/internals/models"

type ISchemaRepository interface {
	SchemaExists() bool

	InitializeEmptySchema() error
	LoadSchema() error

	GetGroupSchema(string) (*models.Schema, error)
	UpdateGroupSchema(*models.Schema, string) error

	CreateGroupSchema(string, []models.SchemaFormItem) error

	// Private methods
	saveCurrentSchemaToFS() error
}
