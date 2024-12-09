package models

type SchemaFormItem struct {
	Title       string `json:"title"`
	ElementType string `json:"element_type"`
	Description string `json:"description"`
}

type Schema struct {
	Schema []SchemaFormItem `json:"schema"`
}

type SchemaMap map[string]Schema
