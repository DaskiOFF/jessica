package projectstruct

import (
	"github.com/daskioff/jessica/configs/keys"
	"github.com/daskioff/jessica/utils/files"
	"github.com/daskioff/jessica/utils/jstrings"
	"github.com/daskioff/jessica/utils/print"
)

// gen генерация структуры проекта и файла шаблона описывающего эту структуру
func (flow *ProjectStructFlow) gen() {
	flow.updateFlags()

	if !useCustomStruct {
		print.PrintlnErrorMessage("Вы не можете генерировать файловую структуру, т.к. эта функция отключена в конфигурационном файле .jessica.yml по ключу '" + keys.KeyCustomProjectStructUse + "'. Для конфигурации можно воспользоваться командой setup")
		return
	}

	if !hasCustomStruct {
		print.PrintlnErrorMessage("Вы не можете генерировать файловую структуру, т.к. не описали ее в конфигурационном файле .jessica.yml по ключу '" + keys.KeyCustomProjectStructDescription + "'. Чтобы увидеть пример описания структуры воспользуйтесь действием `example`")
		return
	}

	flow.generateProjectStruct()
	flow.createTemplateProjectStructDescriptionFile()

	print.PrintlnSuccessMessage("Отредактируйте файл " + flow.templateFileName() + ", чтобы описать вашу структуру. Этот шаблон будет использоваться для генерации раздела структуры проекта в файле README.md.")
	print.PrintlnAttentionMessage("Внимание, при изменении структуры файл " + flow.templateFileName() + " не перезаписывается")
}

// updateFlags обновляет флаги из файла конфигурации проекта
func (flow *ProjectStructFlow) updateFlags() {
	useCustomStruct = flow.projectConfig.GetCustomProjectStructUse()
	hasCustomStruct = flow.projectConfig.GetCustomProjectStructDescription() != nil
}

// templateFileName возвращает имя файла, в котором описывается структура проекта
func (flow *ProjectStructFlow) templateFileName() string {
	return flow.projectConfig.GetCustomProjectStructDescriptionTemplateFilename()
}

// createTemplateProjectStructDescriptionFile создание файла шаблона описывающего структуру проекта
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
