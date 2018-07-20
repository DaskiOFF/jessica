package configs

import (
	"os"

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
