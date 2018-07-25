package setup

import (
	"github.com/daskioff/jessica/configs"
	"github.com/daskioff/jessica/utils"
	"github.com/spf13/viper"
)

func commonSection(config *viper.Viper) {
	companyName := utils.AskQuestion("Your company name (for project): ", false)
	config.Set(configs.KeyCompanyName, companyName)

	projectType := utils.AskQuestionWithAnswers("Project type is [iOS, other]: ", []string{"iOS", "other"})
	config.Set(configs.KeyProjectType, projectType)
}
