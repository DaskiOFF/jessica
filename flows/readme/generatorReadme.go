package readme

import (
	"bufio"
	"io"
	"os"
	"strings"
	textTemplate "text/template"

	"github.com/daskioff/jessica/configs"
	"github.com/daskioff/jessica/utils/files"
	"github.com/daskioff/jessica/utils/print"

	"github.com/daskioff/jessica/utils"
)

// UpdateREADME Проверяет обновляет файл README.md согласно шаблону
func updateREADME() {
	gemFile, _ := readGemfile()
	gemFileDependencies := strings.Join(gemFile, "\n")

	podFile, _ := readPodfile()
	podFileDependencies := strings.Join(podFile, "\n")

	xcodeVersion, _ := readXcodeVersion()
	swiftVersion, _ := readSwiftVersion()

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
		"projectName":         configs.ProjectConfig.Get(configs.KeyIOSProjectName),
	}

	err = executeTemplate(templateFileName(), writer, params)
	if err != nil {
		panic(err)
	}

	projectStructTemplateFilename := configs.ProjectConfig.GetString(configs.KeyCustomProjectStructDescriptionTemplateFilename)
	if files.IsFileExist(projectStructTemplateFilename) {
		writer.WriteString("\n\n")
		executeTemplate(projectStructTemplateFilename, writer, params)
	}

	err = writer.Flush()
	if err != nil {
		panic(err)
	}

	print.PrintlnSuccessMessage(fileNameREADME + " обновлен")
}

func templateFileName() string {
	return configs.ProjectConfig.GetString(configs.KeyReadmeTemplateFilename)
}

func executeTemplate(templateFileName string, writer io.Writer, params map[string]interface{}) error {
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
func checkReadmeTpl() {
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

	content = utils.FixBackQuotes(content)
	fileName := templateFileName()
	if !files.IsFileExist(fileName) {
		files.WriteToFile(fileName, content)
		print.PrintlnSuccessMessage(fileName + " создан")
	}
}
