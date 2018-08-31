package ios

import (
	"github.com/daskioff/jessica/configs/models"
)

type ReadmeIOSFlow struct {
	projectConfig *models.ConfigProject
	iosConfig     *models.ConfigIOS
}

func (flow *ReadmeIOSFlow) Start(args []string) {
	flow.checkFiles()

	flow.updateREADME()
}

func (flow *ReadmeIOSFlow) Description() string {
	return `--------------------------------------------------------------------------------
Поддержка актуальности README.md файла. Генерируется по шаблону.
  Переменные шаблона:
    projectName         – Имя проекта
	xcodeVersion        – Версия xcode из файла
	swiftVersion        – Версия swift из файла
	gemFileDependencies – Список зависимостей Gemfile
	podFileDependencies – Список зависимостей проекта Podfile
--------------------------------------------------------------------------------`
}

// ----------------------------------------------------------------------------
func New(projectConfig *models.ConfigProject, iosConfig *models.ConfigIOS) *ReadmeIOSFlow {
	flow := ReadmeIOSFlow{}
	flow.projectConfig = projectConfig
	flow.iosConfig = iosConfig

	return &flow
}

func (flow *ReadmeIOSFlow) checkFiles() {
	flow.checkXcodeVersionFile()
	flow.checkSwiftVersionFile()
	flow.checkReadmeTplIOS()
}
