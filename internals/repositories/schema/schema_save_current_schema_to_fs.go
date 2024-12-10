package schema

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

func (s *SchemaRepository) saveCurrentSchemaToFS() error {
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
