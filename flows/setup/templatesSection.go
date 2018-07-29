package setup

import (
	"os"

	"github.com/daskioff/jessica/configs"
	"github.com/daskioff/jessica/utils/files"
	"github.com/daskioff/jessica/utils/question"
	"github.com/spf13/viper"
)

func templatesSection(config *viper.Viper) {
	templatesFolderName := "TemplatesJessica"
	config.Set(configs.KeyTemplatesFolderName, templatesFolderName)

	answer := question.AskQuestionWithBoolAnswer("Использовать шаблоны для генерации?")
	config.Set(configs.KeyTemplatesUse, answer)

	if answer && !files.IsFileExist(templatesFolderName) {
		os.Mkdir(templatesFolderName, os.ModePerm)
	}
}
