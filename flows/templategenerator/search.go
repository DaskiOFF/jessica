package templategenerator

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/daskioff/jessica/utils/files"
)

func (flow *TemplateGeneratorFlow) templatesRootPath() string {
	templatesRoot, err := os.Getwd()
	if err != nil {
		fmt.Println(err.Error())
		return ""
	}

	templatesFolderName := flow.projectConfig.GetTemplatesFolderName()
	return filepath.Join(templatesRoot, templatesFolderName)
}

func (flow *TemplateGeneratorFlow) searchTemplates() []string {
	templatesRoot := flow.templatesRootPath()

	folders := flow.folders(templatesRoot)
	templatesFolders := make([]string, 0)
	for _, folder := range folders {
		path := filepath.Join(templatesRoot, folder, TemplateDescriptionFileName)
		if files.IsFileExist(path) {
			templatesFolders = append(templatesFolders, folder)
		}
	}

	return templatesFolders
}

func (flow *TemplateGeneratorFlow) folders(root string) []string {
	files, err := ioutil.ReadDir(root)
	if err != nil {
		fmt.Println(err)
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
