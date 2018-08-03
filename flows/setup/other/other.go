package other

import (
	"github.com/daskioff/jessica/configs/models"
	"github.com/daskioff/jessica/utils/question"
)

func Setup(config *models.ConfigOther, isForce bool) {
	if !config.HasProjectName() || isForce {
		codeProjectName := question.AskQuestion("\nВведите название проекта: ", true)
		config.SetProjectName(codeProjectName)
	}

	if !config.HasProjectFolderName() || isForce {
		codeProjectFolderName := question.AskQuestionWithChooseFolderAnswer("\nВыберите папку с кодом проекта: ")
		if codeProjectFolderName == "" {
			codeProjectFolderName = "."
		}
		config.SetProjectFolderName(codeProjectFolderName)
	}
}
