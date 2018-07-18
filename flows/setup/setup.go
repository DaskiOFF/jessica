package setup

import (
	"strings"

	"github.com/daskioff/jessica/flows"
	"github.com/daskioff/jessica/utils"
)

type SetupFlow struct {
}

func (flow *SetupFlow) Start(args []string) {
	projectName()
}

func (flow *SetupFlow) Description() string {
	return `
	--------------------------------------------------------------------------------
	Первичная настройка файла конфигурации
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
