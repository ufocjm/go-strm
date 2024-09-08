package go_strm

import (
	"path/filepath"
)

func GetSuffixName(fileName string) string {
	return filepath.Ext(fileName)[1:]
}

func Contains(slice []string, value string) bool {
	for _, item := range slice {
		if item == value {
			return true
		}
	}
	return false
}
