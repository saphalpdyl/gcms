package helpers

import (
	"errors"

	"github.com/saphalpdyl/gcms/utils"
)

func ConvertAbrrToElements(abbr string, conversionMap map[string][]string) (string, error) {
	for element, abbrList := range conversionMap {
		if utils.StringInStringList(abbr, abbrList) {
			return element, nil
		}
	}

	return "", errors.New("didn't find the respective element")
}
