package utils

import (
	"fmt"
	"os"
)

func StringInStringList(searchItem string, list []string) bool {
	for _, item := range list {
		if item == searchItem {
			return true
		}
	}

	return false
}

func GenerateDSVFromStringList(list []string) string {
	var finalString = ""

	if len(list) <= 0 {
		return finalString
	}

	for _, item := range list {
		finalString += fmt.Sprintf("%s|", item)
	}

	return finalString[:len(finalString)-1]
}

func PathExists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return false
}
