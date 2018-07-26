package templategenerator

import (
	"bufio"
	"io"
	"os"
	"strings"
	textTemplate "text/template"
	"time"

	"github.com/daskioff/jessica/configs"
	"github.com/daskioff/jessica/utils"
	"github.com/spf13/viper"
)

func generateTemplates(v *viper.Viper, key string, templateName string, moduleName string, customKeys MapKeys, answers MapKeys) []AddedFile {
	codeTemplates := v.Get(key)
	listCodeTemplates := codeTemplates.([]interface{})
	return generateTemplatesFromList(listCodeTemplates, templateName, moduleName, customKeys, answers)
}

func generateTemplatesFromList(list []interface{}, templateName string, moduleName string, customKeys MapKeys, answers MapKeys) []AddedFile {
	addedFiles := []AddedFile{}

	currentTime := time.Now()
	templateInfoParams := params(moduleName, customKeys, answers)
	params := templateInfoParams
	params["developer"] = MapKeys{
		"name":        configs.GlobalConfig.GetString(configs.KeyUserName),
		"companyName": configs.ProjectConfig.GetString(configs.KeyCompanyName),
	}
	params["projectName"] = configs.ProjectConfig.Get(configs.KeyIOSProjectName)
	params["date"] = currentTime.Format("02.01.2006")
	params["year"] = currentTime.Year()

	templateFiles := newTemplateFiles(list, templateName, moduleName, params)
	for _, templateFile := range templateFiles {
		addedFiles = append(addedFiles, AddedFile{
			Path:     templateFile.outputProjectPath,
			Filename: templateFile.name,
		})

		err := os.MkdirAll(templateFile.outputPathFolder, os.ModePerm)
		if err != nil {
			panic(err)
		}

		if templateFile.rewriteResult == rewriteRequest && utils.IsFileExist(templateFile.outputPathFile) {
			utils.PrintlnAttentionMessage("Файл уже существует: " + templateFile.outputPathFile)
			answer := utils.AskQuestionWithBoolAnswer("Перезаписать файл?")
			if !answer {
				continue
			}
		} else if templateFile.rewriteResult == rewriteNo && utils.IsFileExist(templateFile.outputPathFile) {
			continue
		}

		file, err := os.OpenFile(templateFile.outputPathFile, os.O_CREATE|os.O_WRONLY, os.ModePerm)
		if err != nil {
			panic(err)
		}
		defer file.Close()

		writer := bufio.NewWriter(file)
		params["fileName"] = templateFile.name

		err = executeTemplate(templateFile.templatePath, writer, params)
		if err != nil {
			panic(err)
		}

		err = writer.Flush()
		if err != nil {
			panic(err)
		}
	}

	return addedFiles
}

func params(moduleName string, customKeys MapKeys, answers MapKeys) MapKeys {
	return MapKeys{
		"custom":  customKeys,
		"answers": answers,
		"moduleInfo": MapKeys{
			"name":           moduleName,
			"nameUppercase":  strings.ToUpper(moduleName),
			"nameLowercase":  strings.ToLower(moduleName),
			"nameCapitalize": strings.Title(moduleName),
			"nameFirstLower": strings.ToLower(moduleName[:1]) + moduleName[1:],
		},
	}
}

func executeTemplate(templateFileName string, writer io.Writer, params MapKeys) error {
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
