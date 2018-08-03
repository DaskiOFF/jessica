package project

import (
	"os"

	"github.com/daskioff/jessica/configs/keys"
	"github.com/daskioff/jessica/configs/models"
	"github.com/daskioff/jessica/utils/files"
	"github.com/daskioff/jessica/utils/print"
	"github.com/daskioff/jessica/utils/question"
)

func Setup(config *models.ConfigProject, isForce bool) {
	err := config.Validate()
	if err == nil && !isForce {
		return
	}

	commonSection(config, isForce)
	readmeSection(config, isForce)
	customProjectStructSection(config, isForce)
	templatesSection(config, isForce)
}

func commonSection(config *models.ConfigProject, isForce bool) {
	if !config.HasCompanyName() || isForce {
		companyName := question.AskQuestion("Название комании (для проекта): ", false)
		config.SetCompanyName(companyName)
	}

	if !config.HasProjectType() || isForce {
		projectType := question.AskQuestionWithAnswers("Введите тип проекта [iOS, other]: ", []string{"iOS", "other"})
		config.SetProjectType(projectType)
	}
}

func readmeSection(config *models.ConfigProject, isForce bool) {
	if !config.HasReadmeTemplateFilename() || isForce {
		readmeFilename := ".readme.tpl.md"
		config.SetReadmeTemplateFilename(readmeFilename)
	}
}

func customProjectStructSection(config *models.ConfigProject, isForce bool) {
	if !config.HasCustomProjectStructUse() || isForce {
		answer := question.AskQuestionWithBoolAnswer("Use custom project struct?")
		config.SetCustomProjectStructUse(answer)
	}

	if !config.GetCustomProjectStructUse() {
		return
	}

	descriptionFilename := ".project_struct.tpl.md"
	config.SetCustomProjectStructDescriptionTemplateFilename(descriptionFilename)

	const exampleStruct = keys.KeyCustomProjectStructDescription + `:
  - AppLayer:
		- Configs
  - ServiceLayer
  - DataLayer:
		- Entities
  - DomainLayer:
		- Entities
  - PresentationLayer:
    - Resources
		- Components
		- Flows
  - Support`

	print.PrintlnInfoMessage(`
Для создания генерируемой структуры вам необходимо описать ее в локальном файле конфигурации .jessica.yml
Описываемая файловая структура будет создаваться внутри папки проекта
	
Например
` + exampleStruct)
}

func templatesSection(config *models.ConfigProject, isForce bool) {
	if !config.HasTemplatesUse() || isForce {
		answer := question.AskQuestionWithBoolAnswer("Использовать шаблоны для генерации?")
		config.SetTemplatesUse(answer)
	}

	if !config.GetTemplatesUse() {
		return
	}

	templatesFolderName := "TemplatesJessica"
	config.SetTemplatesFolderName(templatesFolderName)

	if !files.IsFileExist(templatesFolderName) {
		os.Mkdir(templatesFolderName, os.ModePerm)
	}
}
