package projectstruct

import (
	"github.com/daskioff/jessica/configs/keys"
	"github.com/daskioff/jessica/configs/models"
	"github.com/daskioff/jessica/utils/files"
	"github.com/daskioff/jessica/utils/jstrings"
	"github.com/daskioff/jessica/utils/print"

	"github.com/daskioff/jessica/flows"
)

var useCustomStruct bool
var hasCustomStruct bool

type ProjectStructFlow struct {
	projectConfig *models.ConfigProject
	iosConfig     *models.ConfigIOS
}

func (flow *ProjectStructFlow) Start(args []string) {
	if len(args) == 0 {
		print.PrintlnAttentionMessage("Необходимо указать какое действие вы хотите выполнить. Чтобы увидеть список действий воспользуйтесь командой help")
		return
	}

	flow.updateFlags()

	if args[0] == "gen" {
		if !useCustomStruct {
			print.PrintlnAttentionMessage("Вы не можете генерировать файловую структуру, т.к. эта функция отключена в конфигурационном файле .jessica.yml по ключу '" + keys.KeyCustomProjectStructUse + "'. Для конфигурации можно воспользоваться командой setup")
			return
		}

		if !hasCustomStruct {
			print.PrintlnAttentionMessage("Вы не можете генерировать файловую структуру, т.к. не описали ее в конфигурационном файле .jessica.yml по ключу '" + keys.KeyCustomProjectStructDescription + "'. Для конфигурации можно воспользоваться командой setup")
			return
		}

		flow.generateProjectStruct()
		flow.createTemplateProjectStructDescriptionFile()

		print.PrintlnSuccessMessage("Отредактируйте файл " + flow.templateFileName() + ", чтобы описать вашу структуру. Этот шаблон будет использоваться для генерации раздела структуры проекта в файле README.md")
	} else {
		print.PrintlnAttentionMessage("Действие не найдено. Чтобы увидеть список действий воспользуйтесь командой help")
	}
}

func (flow *ProjectStructFlow) Description() string {
	return `
--------------------------------------------------------------------------------
	Генерация структуры файлов описаных в конфиге

	Действия:
		gen  - Генерация структуры и необходимых файлов
--------------------------------------------------------------------------------`
}

// ----------------------------------------------------------------------------
func NewFlow(projectConfig *models.ConfigProject, iosConfig *models.ConfigIOS) flows.Flow {
	flow := ProjectStructFlow{}
	flow.projectConfig = projectConfig
	flow.iosConfig = iosConfig

	return &flow
}

func (flow *ProjectStructFlow) updateFlags() {
	useCustomStruct = flow.projectConfig.GetCustomProjectStructUse()
	hasCustomStruct = flow.projectConfig.GetCustomProjectStructDescription() != ""
}

func (flow *ProjectStructFlow) templateFileName() string {
	return flow.projectConfig.GetCustomProjectStructDescriptionTemplateFilename()
}

func (flow *ProjectStructFlow) createTemplateProjectStructDescriptionFile() {
	projectStructure := flow.projectConfig.GetCustomProjectStructDescription()
	projectStructureString := flow.projectStructToString(projectStructure, "  ", "  ")

	content := `# Структура проекта
- %*%{{ .projectName }}%*% – папка проекта
` + projectStructureString

	content = jstrings.FixBackQuotes(content)
	fileName := flow.templateFileName()
	if !files.IsFileExist(fileName) {
		files.WriteToFile(fileName, content)
		print.PrintlnSuccessMessage(fileName + " создан")
	}
}
