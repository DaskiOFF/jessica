package setup

import (
	"github.com/daskioff/jessica/configs"
	"github.com/daskioff/jessica/utils"
	"github.com/spf13/viper"
)

func templatesSection(config *viper.Viper) {
	config.Set(configs.KeyTemplatesFolderName, "TeemplatesJessica")

	answer := utils.AskQuestionWithBoolAnswer("Use templates?")
	config.Set(configs.KeyTemplatesUse, answer)
}
