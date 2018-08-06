package projectstruct

import (
	"os"
	"path/filepath"

	"github.com/daskioff/jessica/configs/keys"
	"github.com/daskioff/jessica/utils/print"
)

func (flow *ProjectStructFlow) generateProjectStruct() {
	if !useCustomStruct || !hasCustomStruct {
		print.PrintlnErrorMessage("Необходима конфигурация с помощью команды `struct setup`")
		return
	}

	projectName := ""
	keyNameForProjectName := ""
	if flow.projectConfig.GetProjectType() == "iOS" {
		projectName = flow.iosConfig.GetFolderNameCode()
		keyNameForProjectName = keys.KeyIOSFolderNameCode
	} else {
		projectName = flow.otherConfig.GetProjectFolderName()
		keyNameForProjectName = keys.KeyOtherProjectFolderName
	}

	if len(projectName) == 0 {
		print.PrintlnErrorMessage("Пропущен шаг создания структуры проекта. Название папки с проектом не указано. В конфигурации ключ " + keyNameForProjectName)
		return
	}

	flow.generateProjectStructInFolder(projectName)

	print.PrintlnSuccessMessage("Структура проекта создана")
}

func (flow *ProjectStructFlow) generateProjectStructInFolder(root string) {
	paths := flow.projectPaths()
	for _, path := range paths {
		resultPath := filepath.Join(root, path)

		os.MkdirAll(resultPath, os.ModePerm)
		print.PrintlnInfoMessage(resultPath)
	}
}
