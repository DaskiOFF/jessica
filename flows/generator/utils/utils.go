package utils

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/daskioff/jessica/utils/files"
)

func TemplatesRootPath(templatesFolderName string) string {
	templatesRoot, err := os.Getwd()
	if err != nil {
		fmt.Println(err.Error())
		return ""
	}

	return filepath.Join(templatesRoot, templatesFolderName)
}

func SearchTemplates(templatesRoot string, templateDescriptionFileName string) []string {
	folders := files.Folders(templatesRoot)

	templatesFolders := make([]string, 0)
	for _, folder := range folders {
		path := filepath.Join(templatesRoot, folder, templateDescriptionFileName)
		if files.IsFileExist(path) {
			templatesFolders = append(templatesFolders, folder)
		}
	}

	return templatesFolders
}
