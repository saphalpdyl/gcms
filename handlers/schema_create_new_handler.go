package handlers

import (
	"log"

	"github.com/saphalpdyl/gcms/helpers"
	"github.com/saphalpdyl/gcms/internals/models"
	"github.com/saphalpdyl/gcms/utils"
)

type SchemaCreateNewHandlerParams struct {
	GroupName string
	FormData  string
}

var ALLOWED_ELEMENTS = []string{"txta", "textarea", "in", "input"}
var ABBR_TO_ELEMENTS = map[string][]string{
	"TEXTAREA": {"txta", "textarea"},
	"INPUT":    {"in", "input"},
}

func (h *Handler) SchemaCreateNewGroup(params SchemaCreateNewHandlerParams) {
	// Validate
	formItemsRaw, err := helpers.ParseStringFromSSV(params.FormData)

	if err != nil {
		log.Fatalf("fatal couldn't parse string sepearated form items values: %v", err)
	}

	formItems := make([]models.SchemaFormItem, 0)

	for title, elementRaw := range formItemsRaw {
		// Validate element input
		if !utils.StringInStringList(elementRaw, ALLOWED_ELEMENTS) {
			log.Fatalf(
				"fatal element %s is not one of [%s]",
				helpers.RenderBold(elementRaw),
				utils.GenerateDSVFromStringList(ALLOWED_ELEMENTS),
			)
		}

		// Converted elementType to capitalized and managed forms
		element, err := helpers.ConvertAbrrToElements(elementRaw, ABBR_TO_ELEMENTS)
		if err != nil {
			log.Fatalf("fatal %v", err)
		}

		formItems = append(formItems, models.SchemaFormItem{
			Title:       title,
			ElementType: element,
		})
	}

	if err := h.schemaRepository.CreateGroupSchema(params.GroupName, formItems); err != nil {
		log.Fatalf("fatal couldn't create group: %v", err)
	}
}
