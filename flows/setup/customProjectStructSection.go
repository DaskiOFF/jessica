package setup

import (
	"github.com/daskioff/jessica/configs"
	"github.com/daskioff/jessica/utils"
	"github.com/daskioff/jessica/utils/print"
	"github.com/spf13/viper"
)

func customProjectStructSection(config *viper.Viper) {
	answer := utils.AskQuestionWithBoolAnswer("Use custom project struct?")
	config.Set(configs.KeyCustomProjectStructUse, answer)

	config.Set(configs.KeyCustomProjectStructDescriptionTemplateFilename, ".project_struct.tpl.md")

	const exampleStruct = configs.KeyCustomProjectStructDescription + `:
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
