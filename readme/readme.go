package readme

import (
	"bufio"
	"io"
	"os"
	"strings"
	"text/template"

	"github.com/daskioff/update_readme_ios/gemfile"
	"github.com/daskioff/update_readme_ios/podfile"
	"github.com/daskioff/update_readme_ios/projectStruct"
	"github.com/daskioff/update_readme_ios/utils"
	"github.com/daskioff/update_readme_ios/versions"
)

const templateFileName = ".readme.tpl.md"

// UpdateREADME Проверяет обновляет файл README.md согласно шаблону
func UpdateREADME() {
	gemFile, _ := gemfile.Read()
	gemFileDependencies := strings.Join(gemFile, "\n")

	podFile, _ := podfile.Read()
	podFileDependencies := strings.Join(podFile, "\n")

	xcodeVersion, _ := versions.ReadXcodeVersion()
	swiftVersion, _ := versions.ReadSwiftVersion()

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
		"projectName":         utils.ProjectName(),
	}

	err = executeTemplate(templateFileName, writer, params)
	if err != nil {
		panic(err)
	}
	if utils.IsFileExist(projectStruct.FileName) {
		writer.WriteString("\n\n")
		executeTemplate(projectStruct.FileName, writer, params)
	}

	err = writer.Flush()
	if err != nil {
		panic(err)
	}

	utils.PrintlnSuccessMessage(fileNameREADME + " successfully updated")
}

func executeTemplate(templateFileName string, writer io.Writer, params map[string]interface{}) error {
	structTemplate, err := template.ParseFiles(templateFileName)
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
func CheckReadmeTpl() {
	content := `[![Swift Version {{ .swiftVersion }}](https://img.shields.io/badge/Swift-{{ .swiftVersion }}-blue.svg?style=flat)](https://developer.apple.com/swift)
[![Recommend xcode version {{ .xcodeVersion }}](https://img.shields.io/badge/Xcode-{{ .xcodeVersion }}-blue.svg?style=flat)](https://developer.apple.com/ios)

**Это сгенерированный файл, для изменения контента редактируйте файл .readme.tpl**

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
	fileName := templateFileName
	if !utils.IsFileExist(fileName) {
		utils.WriteToFile(fileName, content)
		utils.PrintlnSuccessMessage(fileName + " successfully created")
	}
}
