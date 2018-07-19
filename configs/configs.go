package configs

import (
	"github.com/spf13/viper"
)

const configFileName = ".jessica.yml"

var ProjectConfig *viper.Viper

func init() {
	ProjectConfig = viper.New()
	ProjectConfig.SetConfigFile(configFileName)

	if err := ProjectConfig.ReadInConfig(); err != nil {
		return
	}
}
