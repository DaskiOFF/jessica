package main

import (
	"os"
	"strings"
	"text/template"
)

const templateFileName = ".readme.tpl.md"

func updateREADME() {
	gemFile, _ := readGemfile()
	gemFileDependencies := strings.Join(gemFile, "\n")

	podFile, _ := readPodfile()
	podFileDependencies := strings.Join(podFile, "\n")

	xcodeVersion, _ := readXcodeVersion()
	swiftVersion, _ := readSwiftVersion()

	template, err := template.ParseFiles(templateFileName)
	if err != nil {
		panic(err)
	}

	fileNameREADME := "README.md"
	os.Remove(fileNameREADME)

	file, err := os.OpenFile(fileNameREADME, os.O_CREATE|os.O_WRONLY, os.ModePerm)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	err = template.Execute(file, map[string]interface{}{
		"xcodeVersion":        xcodeVersion,
		"swiftVersion":        swiftVersion,
		"gemFileDependencies": gemFileDependencies,
		"podFileDependencies": podFileDependencies,
	})

	if err != nil {
		panic(err)
	}
	printlnSuccessMessage(fileNameREADME + " successfully updated")
}

func checkReadmeTpl() {
	content := `[![Swift Version {{ .swiftVersion }}](https://img.shields.io/badge/Swift-{{ .swiftVersion }}-blue.svg?style=flat)](https://developer.apple.com/swift)
[![Recommend xcode version {{ .xcodeVersion }}](https://img.shields.io/badge/Xcode-{{ .xcodeVersion }}-blue.svg?style=flat)](https://developer.apple.com/ios)

**Это сгенерированный файл, для изменения контента редактируйте файл .readme.tpl**

# Описание проекта

# Краткие данные по проекту

%%**%%Согласно методолгии%%**%% https://ios-factor.com/dependencies

## [Dependencies](https://ios-factor.com/dependencies)
Последний раз проект собирался с версией **Xcode {{ .xcodeVersion }}** указанной в файле %%**%%.xcode-version%%**%% ([Подробнее](https://github.com/fastlane/ci/blob/master/docs/xcode-version.md))

Последний раз проект собирался с версией **Swift {{ .swiftVersion }}** указанной в файле %%**%%.swift-version%%**%%

### Gemfile
В %%**%%Gemfile%%**%% описаны зависимости инструментов. Для установки использовать команду %%**%%bundle install%%**%% ([Подробнее](https://bundler.io/))
%%*%%
{{ .gemFileDependencies }}
%%*%%

### Podfile
Зависимости проекта подключены через %%**%%cocoapods%%**%% и описаны в %%**%%Podfile%%**%%. Для установки использовать: %%**%%[bundle exec] pod install%%**%% или %%**%%[bundle exec] pod update%%**%%
%%*%%
{{ .podFileDependencies }}
%%*%%`

	content = strings.Replace(content, "%%*%%", "```", -1)
	content = strings.Replace(content, "%%**%%", "`", -1)
	fileName := templateFileName
	if !isFileExist(fileName) {
		writeToFile(fileName, content)
		printlnSuccessMessage(fileName + " successfully created")
	}
}
