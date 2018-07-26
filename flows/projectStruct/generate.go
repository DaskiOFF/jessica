package projectstruct

import (
	"os"
	"path/filepath"

	"github.com/daskioff/jessica/configs"
	"github.com/daskioff/jessica/utils"
)

func generateProjectStruct() {
	if !useCustomStruct || !hasCustomStruct {
		utils.PrintlnAttentionMessage("Необходима конфигурация с помощью команды `struct setup`")
		return
	}

	projectName := configs.ProjectConfig.GetString(configs.KeyIOSFolderNameCode)
	if len(projectName) == 0 {
		utils.PrintlnAttentionMessage("Пропущен шаг создания структуры проекта. Название папки с проектом не указано. В конфигурации ключ " + configs.KeyIOSFolderNameCode)
		return
	}

	generateProjectStructInFolder(projectName)

	utils.PrintlnSuccessMessage("Структура проекта создана")
}

func generateProjectStructInFolder(root string) {
	paths := projectPaths()
	for _, path := range paths {
		resultPath := filepath.Join(root, path)

		os.MkdirAll(resultPath, os.ModePerm)
		utils.PrintlnInfoMessage(resultPath)
	}
}
