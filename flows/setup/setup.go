package setup

import (
	"github.com/daskioff/jessica/configs"
	"github.com/daskioff/jessica/flows"
	"github.com/daskioff/jessica/utils"
)

type SetupFlow struct {
}

func (flow *SetupFlow) Start(args []string) {
	setup()
}

func (flow *SetupFlow) Setup() {
	setup()
}

func (flow *SetupFlow) Description() string {
	return `
	--------------------------------------------------------------------------------
	Первичная настройка файла конфигурации

	Имя проекта (Название xcodeproj файла)
	--------------------------------------------------------------------------------
	`
}

// ----------------------------------------------------------------------------
func NewFlow() flows.Flow {
	flow := SetupFlow{}
	return &flow
}

func setup() {
	err := configs.ValidateProjectConfig()
	if err == nil {
		utils.PrintlnSuccessMessage("Файл уже сконфигурирован")
		return
	}

	globalConfig := configs.GlobalConfig
	localConfig := configs.ProjectConfig

	// global config
	globalSection(globalConfig)

	// project config
	commonSection(localConfig)
	readmeSection(localConfig)
	customProjectStructSection(localConfig)
	templatesSection(localConfig)
	iosSection(localConfig)

	configs.WriteGlobal()
	configs.WriteProject()
}
