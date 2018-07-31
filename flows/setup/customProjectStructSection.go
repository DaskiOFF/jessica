package setup

import (
	"github.com/daskioff/jessica/configs/keys"
	"github.com/daskioff/jessica/configs/models"
	"github.com/daskioff/jessica/utils/print"
	"github.com/daskioff/jessica/utils/question"
)

func (flow *SetupFlow) customProjectStructSection(config *models.ConfigProject) {
	answer := question.AskQuestionWithBoolAnswer("Use custom project struct?")
	config.SetCustomProjectStructUse(answer)

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
