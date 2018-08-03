package templategenerator

import (
	"path/filepath"

	"github.com/daskioff/jessica/flows/templategenerator/utils"
	"github.com/daskioff/jessica/utils/files"
)

func searchTemplates(templatesFolderName string, templateDescriptionFileName string) []string {
	templatesRoot := utils.TemplatesRootPath(templatesFolderName)
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
