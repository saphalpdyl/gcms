package schema

import (
	"fmt"

	"github.com/pkg/errors"
	"github.com/saphalpdyl/gcms/internals/models"
)

func (s *SchemaRepository) GetGroupSchema(groupName string) (*models.Schema, error) {
	if s.data == nil {
		return nil, errors.New("Data is empty")
	}

	for k, v := range *s.data {
		if k != groupName {
			continue
		}

		return &v, nil
	}

	return nil, fmt.Errorf("group %s not found", groupName)
}
