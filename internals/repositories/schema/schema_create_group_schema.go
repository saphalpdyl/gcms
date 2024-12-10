package schema

import (
	"fmt"

	"github.com/saphalpdyl/gcms/internals/models"
)

func (s *SchemaRepository) CreateGroupSchema(groupName string, formItems []models.SchemaFormItem) error {
	// Check that the schema doesn't already exists
	schema, _ := s.GetGroupSchema(groupName)

	if schema != nil {
		return fmt.Errorf("group %s already exists in the schema", groupName)
	}

	newSchema := &models.Schema{}
	newSchema.Schema = formItems

	(*s.data)[groupName] = *newSchema

	err := s.saveCurrentSchemaToFS()

	if err != nil {
		return fmt.Errorf("couldn't save schema: %v", err)
	}

	return nil
}
