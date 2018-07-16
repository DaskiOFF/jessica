package readme

import (
	"github.com/daskioff/jessica/flows"
)

type ReadmeFlow struct {
}

func (flow *ReadmeFlow) Start(args []string) {
	checkFiles()

	updateREADME()
}

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
func NewFlow() flows.Flow {
	flow := ReadmeFlow{}
	return &flow
}

func checkFiles() {
	checkXcodeVersionFile()
	checkSwiftVersionFile()
	checkGemfile()
	checkPodfile()
	checkReadmeTpl()
}
