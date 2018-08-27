package other

import (
	"github.com/daskioff/jessica/configs/models"
)

type ReadmeOtherFlow struct {
	projectConfig *models.ConfigProject
	otherConfig   *models.ConfigOther
}

func (flow *ReadmeOtherFlow) Start(args []string) {
	flow.checkFiles()

	flow.updateREADME()
}

func (flow *ReadmeOtherFlow) Description() string {
	return `--------------------------------------------------------------------------------
Поддержка актуальности README.md файла. Генерируется по шаблону.
  Переменные шаблона:
    projectName           – Имя проекта
--------------------------------------------------------------------------------`
}

// ----------------------------------------------------------------------------
func New(projectConfig *models.ConfigProject, otherConfig *models.ConfigOther) *ReadmeOtherFlow {
	flow := ReadmeOtherFlow{}
	flow.projectConfig = projectConfig
	flow.otherConfig = otherConfig

	return &flow
}

func (flow *ReadmeOtherFlow) checkFiles() {
	flow.checkReadmeTplOther()
}
