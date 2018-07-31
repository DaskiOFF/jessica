package setup

import (
	"github.com/daskioff/jessica/configs/models"
	"github.com/daskioff/jessica/flows"
	"github.com/daskioff/jessica/utils/print"
)

type SetupFlow struct {
	globalConfig  *models.ConfigGlobal
	projectConfig *models.ConfigProject
	iosConfig     *models.ConfigIOS
}

func (flow *SetupFlow) Start(args []string) {
	flow.setup()
}

func (flow *SetupFlow) Setup() {
	flow.setup()
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
func NewFlow(globalConfig *models.ConfigGlobal, projectConfig *models.ConfigProject, iosConfig *models.ConfigIOS) flows.Flow {
	flow := SetupFlow{}
	flow.globalConfig = globalConfig
	flow.projectConfig = projectConfig
	flow.iosConfig = iosConfig

	return &flow
}

func (flow *SetupFlow) setup() {
	globalError := flow.globalConfig.Validate()
	projectError := flow.projectConfig.Validate()
	iosError := flow.iosConfig.Validate()

	if globalError == nil && projectError == nil && iosError == nil {
		print.PrintlnSuccessMessage("Файл уже сконфигурирован")
		return
	}

	// global config
	flow.globalSection(flow.globalConfig)

	// project config
	flow.commonSection(flow.projectConfig)
	flow.readmeSection(flow.projectConfig)
	flow.customProjectStructSection(flow.projectConfig)
	flow.templatesSection(flow.projectConfig)
	if flow.projectConfig.GetProjectType() == "iOS" {
		flow.iosSection(flow.iosConfig)
	}

	flow.globalConfig.Write()
	flow.projectConfig.Write()
	flow.iosConfig.Write()
}
