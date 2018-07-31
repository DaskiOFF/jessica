package readme

import (
	"bufio"
	"io"
	"os"
	"strings"
	textTemplate "text/template"

	"github.com/daskioff/jessica/utils/files"
	"github.com/daskioff/jessica/utils/jstrings"
	"github.com/daskioff/jessica/utils/print"
)

// UpdateREADME Проверяет обновляет файл README.md согласно шаблону
func (flow *ReadmeFlow) updateREADME() {
	gemFile, _ := flow.readGemfile()
	gemFileDependencies := strings.Join(gemFile, "\n")

	podFile, _ := flow.readPodfile()
	podFileDependencies := strings.Join(podFile, "\n")

	xcodeVersion, _ := flow.readXcodeVersion()
	swiftVersion, _ := flow.readSwiftVersion()

	fileNameREADME := "README.md"
	os.Remove(fileNameREADME)

	file, err := os.OpenFile(fileNameREADME, os.O_CREATE|os.O_WRONLY, os.ModePerm)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	params := map[string]interface{}{
		"xcodeVersion":        xcodeVersion,
		"swiftVersion":        swiftVersion,
		"gemFileDependencies": gemFileDependencies,
		"podFileDependencies": podFileDependencies,
		"projectName":         flow.iosConfig.GetProjectName(),
	}

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

func (flow *ReadmeFlow) templateFileName() string {
	return flow.projectConfig.GetReadmeTemplateFilename()
}

func (flow *ReadmeFlow) executeTemplate(templateFileName string, writer io.Writer, params map[string]interface{}) error {
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
func (flow *ReadmeFlow) checkReadmeTpl() {
	content := `[![Swift Version {{ .swiftVersion }}](https://img.shields.io/badge/Swift-{{ .swiftVersion }}-blue.svg?style=flat)](https://developer.apple.com/swift)
[![Recommend xcode version {{ .xcodeVersion }}](https://img.shields.io/badge/Xcode-{{ .xcodeVersion }}-blue.svg?style=flat)](https://developer.apple.com/ios)

**Это сгенерированный файл, для изменения контента редактируйте файл .readme.tpl.md**

# Описание проекта {{ .projectName }}

# Краткие данные по проекту

## [Dependencies](https://ios-factor.com/dependencies)
Последний раз проект собирался с версией **Xcode {{ .xcodeVersion }}** указанной в файле %*%.xcode-version%*% ([Подробнее](https://github.com/fastlane/ci/blob/master/docs/xcode-version.md))

Последний раз проект собирался с версией **Swift {{ .swiftVersion }}** указанной в файле %*%.swift-version%*%

### Gemfile
В %*%Gemfile%*% описаны зависимости инструментов. Для установки использовать команду %*%bundle install%*% ([Подробнее](https://bundler.io/))
%***%
{{ .gemFileDependencies }}
%***%

### Podfile
Зависимости проекта подключены через %*%cocoapods%*% и описаны в %*%Podfile%*%. Для установки использовать: %*%[bundle exec] pod install%*% или %*%[bundle exec] pod update%*%
%***%
{{ .podFileDependencies }}
%***%`

	content = jstrings.FixBackQuotes(content)
	fileName := flow.templateFileName()
	if !files.IsFileExist(fileName) {
		files.WriteToFile(fileName, content)
		print.PrintlnSuccessMessage(fileName + " создан")
	}
}
