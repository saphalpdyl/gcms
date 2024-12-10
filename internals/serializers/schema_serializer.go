package serializers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"

	"github.com/saphalpdyl/gcms/internals/models"
)

type SchemaSerializer struct{}

func NewSchemaSerializer() ISerializer[models.SchemaMap] {
	return &SchemaSerializer{}
}

func (s *SchemaSerializer) Deserialize(data []byte, schemaMap *models.SchemaMap) error {
	//Convert from string to schemaMap
	err := json.Unmarshal(data, schemaMap)
	if err != nil {
		// Log and return the error if unmarshalling fails
		return fmt.Errorf("failed to deserialize data into schemaMap: %v", err)
	}

	// Optionally return nil if there is no error
	return nil
}

func (s *SchemaSerializer) Serialize(schemaMap models.SchemaMap) ([]byte, error) {
	// Convert from schemaMap to string
	var buffer bytes.Buffer

	// Create a new JSON encoder and set the indentation for pretty-printing
	encoder := json.NewEncoder(&buffer)
	encoder.SetIndent("", "  ") // Use two spaces for indentation

	// Encode the schemaMap into the buffer
	err := encoder.Encode(schemaMap)
	if err != nil {
		// Log the error and return an empty byte slice in case of an error
		log.Printf("Error serializing schemaMap: %v", err)
		return []byte{}, err
	}

	// Return the buffer's contents as a byte slice
	return buffer.Bytes(), nil
}
