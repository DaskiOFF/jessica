package projectstruct

import (
	"os"

	"github.com/daskioff/jessica/configs"

	"github.com/daskioff/jessica/flows"
	"github.com/daskioff/jessica/utils"
)

const FileName = ".project_struct.tpl.md"
const TemplatesFolderName = "TemplatesJessica"

var useCustomStruct bool
var hasCustomStruct bool
var useTemplateStruct bool

type ProjectStructFlow struct {
}

func (flow *ProjectStructFlow) Start(args []string) {
	if len(args) == 0 {
		utils.PrintlnAttentionMessage("Необходимо указать какое действие вы хотите выполнить. Чтобы увидеть список действий воспользуйтесь командой help")
		return
	}

	updateFlags()

	if args[0] == "gen" {
		if !useCustomStruct {
			utils.PrintlnAttentionMessage("Вы не можете генерировать файловую структуру, т.к. эта функция отключена в конфигурационном файле .jessica.yml по ключу '" + configs.KeyCustomProjectStructUse + "'. Для конфигурации можно воспользоваться действием setup")
			return
		}

		if !hasCustomStruct {
			utils.PrintlnAttentionMessage("Вы не можете генерировать файловую структуру, т.к. не описали ее в конфигурационном файле .jessica.yml по ключу '" + configs.KeyCustomProjectStructDescription + "'. Для конфигурации можно воспользоваться действием setup")
			return
		}

		generateProjectStruct()
		createTemplateProjectStructDescriptionFile()

		utils.PrintlnSuccessMessage("Отредактируйте файл " + FileName + ", чтобы описать вашу структуру. Этот шаблон будет использоваться для генерации раздела структуры проекта в файле README.md")

		if useTemplateStruct {
			os.Mkdir(TemplatesFolderName, os.ModeDir)
		}

	} else if args[0] == "setup" {
		setup()

	} else {
		utils.PrintlnAttentionMessage("Действие не найдено. Чтобы увидеть список действий воспользуйтесь командой help")
	}
}

func (flow *ProjectStructFlow) Setup() {
	setup()
}

func (flow *ProjectStructFlow) Description() string {
	return `
--------------------------------------------------------------------------------
	Генерация структуры файлов описаных в конфиге

	Действия:
		setup – Инициализация необходимых данных
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
	useTemplateStruct = configs.ProjectConfig.GetBool(configs.KeyTemplatesUse)
}

func createTemplateProjectStructDescriptionFile() {
	projectStructure := configs.ProjectConfig.Get(configs.KeyCustomProjectStructDescription)
	projectStructureString := projectStructToString(projectStructure, "  ", "  ")

	content := `# Структура проекта
- %*%{{ .projectName }}%*% – папка проекта
` + projectStructureString

	content = utils.FixBackQuotes(content)
	fileName := FileName
	if !utils.IsFileExist(fileName) {
		utils.WriteToFile(fileName, content)
		utils.PrintlnSuccessMessage(fileName + " successfully created")
	}
}
