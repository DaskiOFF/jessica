package projectstruct

import (
	"os"
	"path/filepath"

	"github.com/daskioff/jessica/configs/keys"
	"github.com/daskioff/jessica/configs/models"
	"github.com/daskioff/jessica/utils/print"
)

func (flow *ProjectStructFlow) generateProjectStruct() {
	if !useCustomStruct {
		print.PrintlnErrorMessage("В конфигурационном файле отключена генерация структуры проекта")
		return
	}
	if !hasCustomStruct {
		print.PrintlnErrorMessage("В конфигурационном файле не описана структура проекта")
		return
	}

	projectFolderPath := ""
	keyNameForProjectName := ""
	switch flow.projectConfig.GetProjectType() {
	case models.ConfigProjectTypeIOS:
		projectFolderPath = flow.iosConfig.GetFolderNameCode()
		keyNameForProjectName = keys.KeyIOSFolderNameCode
	case models.ConfigProjectTypeOther:
		projectFolderPath = flow.otherConfig.GetProjectFolderName()
		keyNameForProjectName = keys.KeyOtherProjectFolderName
	}

	if len(projectFolderPath) == 0 {
		print.PrintlnErrorMessage("Пропущен шаг создания структуры проекта. Название папки с проектом не указано. В конфигурации ключ " + keyNameForProjectName)
		return
	}

	flow.generateProjectStructInFolder(projectFolderPath)

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
