package setup

import (
	"github.com/daskioff/jessica/configs/models"
)

func (flow *SetupFlow) readmeSection(config *models.ConfigProject) {
	readmeFilename := ".readme.tpl.md"
	config.SetReadmeTemplateFilename(readmeFilename)
}
