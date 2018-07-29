package setup

import (
	"github.com/daskioff/jessica/configs"
	"github.com/daskioff/jessica/utils/question"
	"github.com/spf13/viper"
)

func globalSection(config *viper.Viper) {
	username := question.AskQuestion("Ваше имя (для глобального): ", true)
	config.Set(configs.KeyUserName, username)
}
