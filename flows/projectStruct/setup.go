package projectstruct

import (
	"github.com/daskioff/jessica/configs"
	"github.com/daskioff/jessica/utils"
)

func setup() {
	if !useCustomStruct {
		useCustomStruct = requestUseCustomStruct()

		configs.ProjectConfig.Set(configs.KeyUseCustomProjectStruct, useCustomStruct)
		configs.WriteProject()
	}

	if !useCustomStruct {
		utils.PrintlnAttentionMessage("Skipped the creation of the project structure")
		return
	}

	if hasCustomStruct {
		showCurrentProjectStruct()
	} else {
		showExample()
	}

	if !useTemplateStruct {
		useTemplateStruct = requestUseCustomTemplatesStruct()

		configs.ProjectConfig.Set(configs.KeyCustomTemplatesStruct, useTemplateStruct)
		configs.WriteProject()
	}
}

func requestUseCustomTemplatesStruct() bool {
	answer := utils.AskQuestionWithAnswers("Use templates? (y/n)", []string{"y", "n", "Y", "N"})

	if answer == "y" || answer == "Y" {
		return true
	}

	return false
}

// requestUseCustomStruct Проверяет наличие файла шаблона описывающего структуру проекта, если его нет, то предлагает создать его и структуру из папок
func requestUseCustomStruct() bool {
	answer := utils.AskQuestionWithAnswers("Use custom project struct? (y/n)", []string{"y", "n", "Y", "N"})

	if answer == "y" || answer == "Y" {
		return true
	}

	return false
}

func showExample() {
	const exampleStruct = configs.KeyCustomProjectStruct + `:
  - config
  - di:
    - factories
  - extensions
  - models
  - services:
    - api
  - usecases
  - presentation:
    - resources:
      - r
      - localization
      - fonts
    - flows
    - components:
      - views
      - tableCells
      - collectionCells
    - controllers
  - support`

	utils.PrintlnInfoMessage(`
Для создания генерируемой структуры вам необходимо описать ее в локальном файле конфигурации .jessica.yml
Описываемая файловая структура будет создаваться внутри папки проекта
	
Например
` + exampleStruct)
}

func showCurrentProjectStruct() {
	projectStructure := configs.ProjectConfig.Get(configs.KeyCustomProjectStruct)

	info := projectStructToString(projectStructure, "  ", "  ")
	utils.PrintlnInfoMessage("Структура из файла конфигурации\n\n" + info)
}
