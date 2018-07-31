package setup

import (
	"github.com/daskioff/jessica/configs/models"
	"github.com/daskioff/jessica/utils/question"
)

func (flow *SetupFlow) globalSection(config *models.ConfigGlobal) {
	username := question.AskQuestion("Ваше имя (для глобального): ", true)
	config.SetUsername(username)
}
