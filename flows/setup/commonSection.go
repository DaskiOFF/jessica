package setup

import (
	"github.com/daskioff/jessica/configs"
	"github.com/daskioff/jessica/utils/question"
	"github.com/spf13/viper"
)

func commonSection(config *viper.Viper) {
	companyName := question.AskQuestion("Название комании (для проекта): ", false)
	config.Set(configs.KeyCompanyName, companyName)

	projectType := question.AskQuestionWithAnswers("Введите тип проекта [iOS, other]: ", []string{"iOS", "other"})
	config.Set(configs.KeyProjectType, projectType)
}
