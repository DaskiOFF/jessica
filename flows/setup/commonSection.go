package setup

import (
	"github.com/daskioff/jessica/configs/models"
	"github.com/daskioff/jessica/utils/question"
)

func (flow *SetupFlow) commonSection(config *models.ConfigProject) {
	companyName := question.AskQuestion("Название комании (для проекта): ", false)
	config.SetCompanyName(companyName)

	projectType := question.AskQuestionWithAnswers("Введите тип проекта [iOS, other]: ", []string{"iOS", "other"})
	config.SetProjectType(projectType)
}
