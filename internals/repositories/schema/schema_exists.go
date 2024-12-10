package schema

import (
	"path"

	"github.com/saphalpdyl/gcms/utils"
)

func (s *SchemaRepository) SchemaExists() bool {
	return utils.PathExists(path.Join(s.schemaFolderPath, s.schemaFileName))
}
