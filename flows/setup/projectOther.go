package setup

import (
	"github.com/daskioff/jessica/configs/models"
	"github.com/daskioff/jessica/utils/question"
)

func (flow *SetupFlow) projectOther(config *models.ConfigOther, isForce bool) {
	if !config.HasProjectFolderName() || isForce {
		codeProjectFolderName := question.AskQuestionWithChooseFolderAnswer("\nВыберите папку с кодом проекта: ")
		if codeProjectFolderName == "" {
			codeProjectFolderName = "."
		}
		config.SetProjectFolderName(codeProjectFolderName)
	}
}
