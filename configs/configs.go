package configs

import (
	"os"

	"github.com/daskioff/jessica/utils/print"
	"github.com/spf13/viper"
)

const configFileName = ".jessica.yml"

var ProjectConfig *viper.Viper
var GlobalConfig *viper.Viper

func init() {
	ProjectConfig = viper.New()
	ProjectConfig.SetConfigFile(configFileName)

	if err := ProjectConfig.ReadInConfig(); err != nil {
	}

	GlobalConfig = viper.New()
	GlobalConfig.SetConfigFile(os.Getenv("HOME") + "/" + configFileName)

	if err := GlobalConfig.ReadInConfig(); err != nil {
	}
}

func WriteGlobal() {
	err := GlobalConfig.WriteConfig()
	if err != nil {
		print.PrintlnErrorMessage("Ошибка сохранения глобального файла конфигурации: " + err.Error())
	}
}

func WriteProject() {
	err := ProjectConfig.WriteConfig()
	if err != nil {
		print.PrintlnErrorMessage("Ошибка сохранения локального файла конфигурации: " + err.Error())
	}
}
