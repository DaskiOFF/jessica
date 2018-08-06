package readme

import (
	"github.com/daskioff/jessica/configs/models"
	"github.com/daskioff/jessica/flows"
)

type ReadmeFlow struct {
	projectConfig *models.ConfigProject
	iosConfig     *models.ConfigIOS
	otherConfig   *models.ConfigOther
}

func (flow *ReadmeFlow) Start(args []string) {
	flow.checkFiles()

	flow.updateREADME()
}

func (flow *ReadmeFlow) Description() string {
	return `--------------------------------------------------------------------------------
Поддержка актуальности README.md файла. Генерируется по шаблону.
  Переменные шаблона:
    projectName           – Имя проекта

    Для iOS проекта:
      xcodeVersion        – Версия xcode из файла
      swiftVersion        – Версия swift из файла
      gemFileDependencies – Список зависимостей Gemfile
      podFileDependencies – Список зависимостей проекта Podfile
--------------------------------------------------------------------------------`
}

// ----------------------------------------------------------------------------
func NewFlow(projectConfig *models.ConfigProject, iosConfig *models.ConfigIOS, otherConfig *models.ConfigOther) flows.Flow {
	flow := ReadmeFlow{}
	flow.projectConfig = projectConfig
	flow.iosConfig = iosConfig
	flow.otherConfig = otherConfig

	return &flow
}

func (flow *ReadmeFlow) checkFiles() {
	if flow.projectConfig.GetProjectType() == "iOS" {
		flow.checkXcodeVersionFile()
		flow.checkSwiftVersionFile()
		flow.checkReadmeTplIOS()
	} else {
		flow.checkReadmeTplOther()
	}
}
