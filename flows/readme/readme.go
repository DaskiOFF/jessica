package readme

import (
	"github.com/daskioff/jessica/configs/models"
	"github.com/daskioff/jessica/flows"
)

type ReadmeFlow struct {
	projectConfig *models.ConfigProject
	iosConfig     *models.ConfigIOS
}

func (flow *ReadmeFlow) Start(args []string) {
	flow.checkFiles()

	flow.updateREADME()
}

func (flow *ReadmeFlow) Setup() {}

func (flow *ReadmeFlow) Description() string {
	return `
	--------------------------------------------------------------------------------
	Поддержка актуальности README.md файла. Генерируется по шаблону.

	Переменные шаблона:
		xcodeVersion        – Версия xcode из файла
		swiftVersion        – Версия swift из файла
		gemFileDependencies – Список зависимостей Gemfile
		podFileDependencies – Список зависимостей проекта Podfile
		projectName         – Имя проекта
	--------------------------------------------------------------------------------
	`
}

// ----------------------------------------------------------------------------
func NewFlow(projectConfig *models.ConfigProject, iosConfig *models.ConfigIOS) flows.Flow {
	flow := ReadmeFlow{}
	flow.projectConfig = projectConfig
	flow.iosConfig = iosConfig

	return &flow
}

func (flow *ReadmeFlow) checkFiles() {
	if flow.projectConfig.GetProjectType() == "iOS" {
		flow.checkXcodeVersionFile()
		flow.checkSwiftVersionFile()
		if flow.iosConfig.GetGemfileUse() {
			flow.checkGemfile()
		}
		if flow.iosConfig.GetPodfileUse() {
			flow.checkPodfile()
		}
		flow.checkReadmeTpl()
	}
}
