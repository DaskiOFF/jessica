package utils

import (
	"fmt"
	"os"
	"path/filepath"
)

func TemplatesRootPath(templatesFolderName string) string {
	templatesRoot, err := os.Getwd()
	if err != nil {
		fmt.Println(err.Error())
		return ""
	}

	return filepath.Join(templatesRoot, templatesFolderName)
}
