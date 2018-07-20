package setup

import (
	"github.com/daskioff/jessica/configs"
	"github.com/daskioff/jessica/flows"
	"github.com/daskioff/jessica/utils"
)

type SetupFlow struct {
}

func (flow *SetupFlow) Start(args []string) {
	err := configs.ValidateProjectConfig()
	if err == nil {
		utils.PrintlnSuccessMessage("Файл уже сконфигурирован")
		return
	}
	username := userName()
	companyName := companyName()
	projectName := projectName()

	localConfig := configs.ProjectConfig
	globalConfig := configs.GlobalConfig

	globalConfig.Set(configs.KeyUserName, username)

	localConfig.Set(configs.KeyCompanyName, companyName)
	localConfig.Set(configs.KeyProjectName, projectName)
	localConfig.Set(configs.KeyProjectXcodeProjName, projectName+".xcodeproj")

	err = localConfig.WriteConfig()
	if err != nil {
		utils.PrintlnErrorMessage("Ошибка сохранения локального файла конфигурации: " + err.Error())
	}

	err = globalConfig.WriteConfig()
	if err != nil {
		utils.PrintlnErrorMessage("Ошибка сохранения глобального файла конфигурации: " + err.Error())
	}
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
