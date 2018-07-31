package configs

import (
	"os"

	"github.com/daskioff/jessica/configs/models"
	"github.com/spf13/viper"
)

const configFileName = ".jessica.yml"

var projectConfig *viper.Viper
var globalConfig *viper.Viper

func init() {
	projectConfig = viper.New()
	projectConfig.SetConfigFile(configFileName)

	if err := projectConfig.ReadInConfig(); err != nil {
	}

	globalConfig = viper.New()
	globalConfig.SetConfigFile(os.Getenv("HOME") + "/" + configFileName)

	if err := globalConfig.ReadInConfig(); err != nil {
	}
}

func Global() *models.ConfigGlobal {
	return models.NewGlobal(globalConfig)
}

func Project() *models.ConfigProject {
	return models.NewProject(projectConfig)
}

func IOS() *models.ConfigIOS {
	return models.NewIOS(projectConfig)
}
