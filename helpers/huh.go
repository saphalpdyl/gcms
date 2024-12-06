package helpers

import (
	"fmt"

	"github.com/charmbracelet/huh"
)

func GenerateSelectFormItemFromStringList(result *[]string, itemList []string, title string, prefix string) *huh.MultiSelect[string] {
	var optionsList []huh.Option[string]

	for _, item := range itemList {
		optionsList = append(optionsList, huh.NewOption[string](fmt.Sprintf("%s%s", prefix, item), item))
	}

	return huh.NewMultiSelect[string]().
		Title(title).
		Options(optionsList...).
		Value(result)
}
