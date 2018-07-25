package projectstruct

import (
	"os"
	"path/filepath"

	"github.com/daskioff/jessica/configs"
	"github.com/daskioff/jessica/utils"
)

func generateProjectStruct() {
	if !useCustomStruct || !hasCustomStruct {
		utils.PrintlnAttentionMessage("Необходимо сначала сконфигурировать с помощью команды `struct setup`")
		return
	}

	projectName := configs.ProjectConfig.GetString(configs.KeyProjectName)
	if len(projectName) == 0 {
		utils.PrintlnAttentionMessage("Skipped the creation of the project structure. Project name is empty")
		return
	}

	generateProjectStructInFolder(projectName)

	utils.PrintlnSuccessMessage("Project structure created")
}

func generateProjectStructInFolder(root string) {
	paths := projectPaths()
	for _, path := range paths {
		resultPath := filepath.Join(root, path)

		os.MkdirAll(resultPath, os.ModePerm)
		utils.PrintlnInfoMessage(resultPath)
	}
}
