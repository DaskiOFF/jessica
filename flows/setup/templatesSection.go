package setup

import (
	"os"

	"github.com/daskioff/jessica/configs/models"
	"github.com/daskioff/jessica/utils/files"
	"github.com/daskioff/jessica/utils/question"
)

func (flow *SetupFlow) templatesSection(config *models.ConfigProject) {
	templatesFolderName := "TemplatesJessica"
	config.SetTemplatesFolderName(templatesFolderName)

	answer := question.AskQuestionWithBoolAnswer("Использовать шаблоны для генерации?")
	config.SetTemplatesUse(answer)

	if answer && !files.IsFileExist(templatesFolderName) {
		os.Mkdir(templatesFolderName, os.ModePerm)
	}
}
