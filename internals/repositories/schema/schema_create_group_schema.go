package schema

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/saphalpdyl/gcms/internals/models"
)

func (s *SchemaRepository) CreateGroupSchema(groupName string, formItems []models.SchemaFormItem) error {
	// Check that the schema doesn't already exists
	schema, err := s.GetGroupSchema(groupName)

	if err != nil {
		return fmt.Errorf("couldn't validate group existence: %v", err)
	}

	if schema != nil {
		return fmt.Errorf("group %s already exists in the schema", groupName)
	}

	newSchema := &models.Schema{}
	newSchema.Schema = formItems

	(*s.data)[groupName] = *newSchema

	err = s.saveCurrentSchemaToFS()

	if err != nil {
		return fmt.Errorf("couldn't save schema: %v", err)
	}

	file, err := os.Create(filepath.Join(s.schemaFolderPath, s.schemaFileName))
	if err != nil {
		return fmt.Errorf("%v", err)
	}

	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", " ")

	if err = encoder.Encode((*s.data)); err != nil {
		return fmt.Errorf("couldn't encode json")
	}

	return nil
}
