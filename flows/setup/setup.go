package setup

import (
	"strings"

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
	projectName := projectName()

	config := configs.ProjectConfig
	config.Set(configs.KeyProjectName, projectName)
	config.Set(configs.KeyProjectXcodeProjName, projectName+".xcodeproj")
	config.WriteConfig()
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

func projectName() string {
	projectName := utils.ProjectName()
	var answer string = "n"
	if projectName != "" {
		answer = utils.AskQuestionWithAnswers("Your project has name '"+projectName+"'? (y/n): ", []string{"y", "n", "Y", "N"})
	}

	if strings.ToLower(answer) == "n" {
		for {
			answer := utils.AskQuestion("Enter project name: ", true)
			projectName = answer
			if !strings.HasSuffix(answer, ".xcodeproj") {
				projectName = answer + ".xcodeproj"
			}
			if utils.IsFileExist(projectName) {
				break
			}
			utils.PrintlnInfoMessage("Файл '" + projectName + "' не найден")
		}
	}

	return projectName
}
