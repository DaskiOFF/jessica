package projectstruct

import (
	"github.com/daskioff/jessica/configs"
	"github.com/daskioff/jessica/utils/files"
	"github.com/daskioff/jessica/utils/jstrings"
	"github.com/daskioff/jessica/utils/print"

	"github.com/daskioff/jessica/flows"
)

var useCustomStruct bool
var hasCustomStruct bool

type ProjectStructFlow struct {
}

func (flow *ProjectStructFlow) Start(args []string) {
	if len(args) == 0 {
		print.PrintlnAttentionMessage("Необходимо указать какое действие вы хотите выполнить. Чтобы увидеть список действий воспользуйтесь командой help")
		return
	}

	updateFlags()

	if args[0] == "gen" {
		if !useCustomStruct {
			print.PrintlnAttentionMessage("Вы не можете генерировать файловую структуру, т.к. эта функция отключена в конфигурационном файле .jessica.yml по ключу '" + configs.KeyCustomProjectStructUse + "'. Для конфигурации можно воспользоваться командой setup")
			return
		}

		if !hasCustomStruct {
			print.PrintlnAttentionMessage("Вы не можете генерировать файловую структуру, т.к. не описали ее в конфигурационном файле .jessica.yml по ключу '" + configs.KeyCustomProjectStructDescription + "'. Для конфигурации можно воспользоваться командой setup")
			return
		}

		generateProjectStruct()
		createTemplateProjectStructDescriptionFile()

		print.PrintlnSuccessMessage("Отредактируйте файл " + templateFileName() + ", чтобы описать вашу структуру. Этот шаблон будет использоваться для генерации раздела структуры проекта в файле README.md")
	} else {
		print.PrintlnAttentionMessage("Действие не найдено. Чтобы увидеть список действий воспользуйтесь командой help")
	}
}

func (flow *ProjectStructFlow) Setup() {
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
func NewFlow() flows.Flow {
	flow := ProjectStructFlow{}
	return &flow
}

func updateFlags() {
	useCustomStruct = configs.ProjectConfig.GetBool(configs.KeyCustomProjectStructUse)
	hasCustomStruct = configs.ProjectConfig.IsSet(configs.KeyCustomProjectStructDescription)
}

func templateFileName() string {
	return configs.ProjectConfig.GetString(configs.KeyCustomProjectStructDescriptionTemplateFilename)
}

func createTemplateProjectStructDescriptionFile() {
	projectStructure := configs.ProjectConfig.Get(configs.KeyCustomProjectStructDescription)
	projectStructureString := projectStructToString(projectStructure, "  ", "  ")

	content := `# Структура проекта
- %*%{{ .projectName }}%*% – папка проекта
` + projectStructureString

	content = jstrings.FixBackQuotes(content)
	fileName := templateFileName()
	if !files.IsFileExist(fileName) {
		files.WriteToFile(fileName, content)
		print.PrintlnSuccessMessage(fileName + " создан")
	}
}
