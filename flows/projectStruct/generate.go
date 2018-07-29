package projectstruct

import (
	"os"
	"path/filepath"

	"github.com/daskioff/jessica/configs"
	"github.com/daskioff/jessica/utils/print"
)

func generateProjectStruct() {
	if !useCustomStruct || !hasCustomStruct {
		print.PrintlnAttentionMessage("Необходима конфигурация с помощью команды `struct setup`")
		return
	}

	projectName := configs.ProjectConfig.GetString(configs.KeyIOSFolderNameCode)
	if len(projectName) == 0 {
		print.PrintlnAttentionMessage("Пропущен шаг создания структуры проекта. Название папки с проектом не указано. В конфигурации ключ " + configs.KeyIOSFolderNameCode)
		return
	}

	generateProjectStructInFolder(projectName)

	print.PrintlnSuccessMessage("Структура проекта создана")
}

func generateProjectStructInFolder(root string) {
	paths := projectPaths()
	for _, path := range paths {
		resultPath := filepath.Join(root, path)

		os.MkdirAll(resultPath, os.ModePerm)
		print.PrintlnInfoMessage(resultPath)
	}
}
