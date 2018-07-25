package templategenerator

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/daskioff/jessica/flows/projectstruct"
	"github.com/daskioff/jessica/utils"
)

func templatesRootPath() string {
	templatesRoot, err := os.Getwd()
	if err != nil {
		fmt.Println(err.Error())
		return ""
	}
	return filepath.Join(templatesRoot, projectstruct.TemplatesFolderName)
}

func searchTemplates() []string {
	templatesRoot := templatesRootPath()

	folders := folders(templatesRoot)
	templatesFolders := make([]string, 0)
	for _, folder := range folders {
		path := filepath.Join(templatesRoot, folder, TemplateDescriptionFileName)
		if utils.IsFileExist(path) {
			templatesFolders = append(templatesFolders, folder)
		}
	}

	return templatesFolders
}

func folders(root string) []string {
	files, err := ioutil.ReadDir(root)
	if err != nil {
		return []string{}
	}

	folders := make([]string, 0)
	for _, file := range files {
		if file.IsDir() {
			folders = append(folders, file.Name())
		}
	}

	return folders
}