package models

type FileMetadata struct {
	FilePath string            `json:"filepath"`
	Metadata map[string]string `json:"metadata"`
}

type GroupData struct {
	Group string          `json:"group"`
	Files []*FileMetadata `json:"files"`
}

type RootMetaData struct {
	LastUpdated int64 `json:"last_updated"`

	Data []*GroupData `json:"data"`
}
