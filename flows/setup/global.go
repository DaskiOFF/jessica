package setup

import (
	"github.com/daskioff/jessica/configs/models"
	"github.com/daskioff/jessica/utils/question"
)

func (flow *SetupFlow) setupGlobal(config *models.ConfigGlobal, isForce bool) {
	err := flow.globalConfig.Validate()
	if err == nil && !isForce {
		return
	}

	if config.GetUsername() == "" || isForce {
		username := question.AskQuestion("Ваше имя (для глобального): ", true)
		config.SetUsername(username)
	}
}
