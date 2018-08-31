package other

import (
	"bufio"
	"io"
	"os"
	textTemplate "text/template"

	"github.com/daskioff/jessica/utils/files"
	"github.com/daskioff/jessica/utils/jstrings"
	"github.com/daskioff/jessica/utils/print"
)

// UpdateREADME Проверяет обновляет файл README.md согласно шаблону
func (flow *ReadmeOtherFlow) updateREADME() {
	params := map[string]interface{}{}
	params["projectName"] = flow.otherConfig.GetProjectName()

	fileNameREADME := "README.md"
	os.Remove(fileNameREADME)

	file, err := os.OpenFile(fileNameREADME, os.O_CREATE|os.O_WRONLY, os.ModePerm)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	writer := bufio.NewWriter(file)

	err = flow.executeTemplate(flow.templateFileName(), writer, params)
	if err != nil {
		panic(err)
	}

	projectStructTemplateFilename := flow.projectConfig.GetCustomProjectStructDescriptionTemplateFilename()
	if files.IsFileExist(projectStructTemplateFilename) {
		writer.WriteString("\n\n")
		flow.executeTemplate(projectStructTemplateFilename, writer, params)
	}

	err = writer.Flush()
	if err != nil {
		panic(err)
	}

	print.PrintlnSuccessMessage(fileNameREADME + " обновлен")
}

func (flow *ReadmeOtherFlow) templateFileName() string {
	return flow.projectConfig.GetReadmeTemplateFilename()
}

func (flow *ReadmeOtherFlow) executeTemplate(templateFileName string, writer io.Writer, params map[string]interface{}) error {
	structTemplate, err := textTemplate.ParseFiles(templateFileName)
	if err != nil {
		return err
	}

	err = structTemplate.Execute(writer, params)
	if err != nil {
		return err
	}

	return nil
}

// CheckReadmeTpl Проверяет существование файла описывающего шаблон README, если его нет, то его создает и заполняет значением по умолчанию
func (flow *ReadmeOtherFlow) checkReadmeTplOther() {
	content := `**Это сгенерированный файл, для изменения контента редактируйте файл .readme.tpl.md**

# Описание проекта {{ .projectName }}

# Краткие данные по проекту`

	content = jstrings.FixBackQuotes(content)
	fileName := flow.templateFileName()
	if !files.IsFileExist(fileName) {
		files.WriteToFile(fileName, content)
		print.PrintlnSuccessMessage(fileName + " создан")
	}
}
