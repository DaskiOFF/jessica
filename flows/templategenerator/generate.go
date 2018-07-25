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

func generateTemplates(v *viper.Viper, key string, templateName string, moduleName string, customKeys map[string]interface{}) []AddedFile {
	codeTemplates := v.Get(key)
	listCodeTemplates := codeTemplates.([]interface{})
	return generateTemplatesFromList(listCodeTemplates, templateName, moduleName, customKeys)
}

func generateTemplatesFromList(list []interface{}, templateName string, moduleName string, customKeys map[string]interface{}) []AddedFile {
	addedFiles := []AddedFile{}

	templateFiles := newTemplateFiles(list, templateName, moduleName)
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
			answer := utils.AskQuestionWithAnswers("Перезаписать файл? (y/n): ", []string{"y", "n", "Y", "N"})
			if strings.ToLower(answer) == "n" {
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

		currentTime := time.Now()

		writer := bufio.NewWriter(file)
		params := map[string]interface{}{
			"custom": customKeys,
			"moduleInfo": map[string]interface{}{
				"name":           moduleName,
				"nameUppercase":  strings.ToUpper(moduleName),
				"nameLowercase":  strings.ToLower(moduleName),
				"nameCapitalize": strings.Title(moduleName),
				"nameFirstLower": strings.ToLower(moduleName[:1]) + moduleName[1:],
			},
			"developer": map[string]interface{}{
				"name":         configs.GlobalConfig.GetString(configs.KeyUserName),
				"company_name": configs.ProjectConfig.GetString(configs.KeyCompanyName),
			},
			"fileName":    templateFile.name,
			"projectName": configs.ProjectConfig.Get(configs.KeyIOSProjectName),
			"date":        currentTime.Format("02.01.2006"),
			"year":        currentTime.Year(),
		}

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
