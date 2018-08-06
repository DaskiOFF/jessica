package global

import (
	"github.com/daskioff/jessica/configs/models"
	"github.com/daskioff/jessica/utils/question"
)

func Setup(config *models.ConfigGlobal, isForce bool) {
	err := config.Validate()
	if err == nil && !isForce {
		return
	}

	if !config.HasUsername() || isForce {
		username := question.AskQuestion("Ваше имя (для глобального): ", true)
		config.SetUsername(username)
	}
}
