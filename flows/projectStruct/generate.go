package projectstruct

import (
	"os"
	"path/filepath"

	"github.com/daskioff/jessica/configs/keys"
	"github.com/daskioff/jessica/utils/print"
)

func (flow *ProjectStructFlow) generateProjectStruct() {
	if !useCustomStruct || !hasCustomStruct {
		print.PrintlnAttentionMessage("Необходима конфигурация с помощью команды `struct setup`")
		return
	}

	projectName := flow.iosConfig.GetFolderNameCode()
	if len(projectName) == 0 {
		print.PrintlnAttentionMessage("Пропущен шаг создания структуры проекта. Название папки с проектом не указано. В конфигурации ключ " + keys.KeyIOSFolderNameCode)
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
