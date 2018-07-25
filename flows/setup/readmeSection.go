package setup

import (
	"github.com/daskioff/jessica/configs"
	"github.com/spf13/viper"
)

func readmeSection(config *viper.Viper) {
	config.Set(configs.KeyReadmeTemplateFilename, ".readme.tpl.md")
}
