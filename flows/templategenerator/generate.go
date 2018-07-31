package templategenerator

import (
	"bufio"
	"io"
	"os"
	"strings"
	textTemplate "text/template"
	"time"

	"github.com/daskioff/jessica/utils/files"
	"github.com/daskioff/jessica/utils/print"
	"github.com/daskioff/jessica/utils/question"
	"github.com/spf13/viper"
)

func (flow *TemplateGeneratorFlow) generateTemplates(v *viper.Viper, key string, templateName string, moduleName string, customKeys MapKeys, answers MapKeys) []AddedFile {
	codeTemplates := v.Get(key)
	if codeTemplates == nil {
		return []AddedFile{}
	}

	listCodeTemplates := codeTemplates.([]interface{})
	return flow.generateTemplatesFromList(listCodeTemplates, templateName, moduleName, customKeys, answers)
}

func (flow *TemplateGeneratorFlow) generateTemplatesFromList(list []interface{}, templateName string, moduleName string, customKeys MapKeys, answers MapKeys) []AddedFile {
	addedFiles := []AddedFile{}

	currentTime := time.Now()
	templateInfoParams := flow.params(moduleName, customKeys, answers)
	params := templateInfoParams
	params["developer"] = MapKeys{
		"name":        flow.globalConfig.GetUsername(),
		"companyName": flow.projectConfig.GetCompanyName(),
	}
	params["projectName"] = flow.iosConfig.GetProjectName()
	params["date"] = currentTime.Format("02.01.2006")
	params["year"] = currentTime.Year()

	templateFiles := flow.newTemplateFiles(list, templateName, moduleName, params)
	for _, templateFile := range templateFiles {
		addedFiles = append(addedFiles, AddedFile{
			Path:     templateFile.outputProjectPath,
			Filename: templateFile.name,
		})

		err := os.MkdirAll(templateFile.outputPathFolder, os.ModePerm)
		if err != nil {
			panic(err)
		}

		if templateFile.rewriteResult == rewriteRequest && files.IsFileExist(templateFile.outputPathFile) {
			print.PrintlnAttentionMessage("Файл уже существует: " + templateFile.outputPathFile)
			answer := question.AskQuestionWithBoolAnswer("Перезаписать файл?")
			if !answer {
				continue
			}
		} else if templateFile.rewriteResult == rewriteNo && files.IsFileExist(templateFile.outputPathFile) {
			continue
		}

		file, err := os.OpenFile(templateFile.outputPathFile, os.O_CREATE|os.O_WRONLY, os.ModePerm)
		if err != nil {
			panic(err)
		}
		defer file.Close()

		writer := bufio.NewWriter(file)
		params["fileName"] = templateFile.name

		err = flow.executeTemplate(templateFile.templatePath, writer, params)
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

func (flow *TemplateGeneratorFlow) params(moduleName string, customKeys MapKeys, answers MapKeys) MapKeys {
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

func (flow *TemplateGeneratorFlow) executeTemplate(templateFileName string, writer io.Writer, params MapKeys) error {
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
