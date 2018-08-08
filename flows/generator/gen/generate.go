package gen

import (
	"bufio"
	"io"
	"os"
	"strings"
	textTemplate "text/template"
	"time"

	"github.com/daskioff/jessica/configs/models"
	"github.com/daskioff/jessica/utils/files"
	"github.com/daskioff/jessica/utils/print"
	"github.com/daskioff/jessica/utils/question"
	"github.com/daskioff/jessica/utils/xcodeproj"
	"github.com/spf13/viper"
)

type MapKeys map[string]interface{}

type generateParams struct {
	customKeys    MapKeys
	answers       MapKeys
	globalConfig  *models.ConfigGlobal
	projectConfig *models.ConfigProject
	iosConfig     *models.ConfigIOS
	otherConfig   *models.ConfigOther
}

func generateTemplates(v *viper.Viper, key string, templateName string, moduleName string, generateParams generateParams) []xcodeproj.AddedFile {
	codeTemplates := v.Get(key)
	if codeTemplates == nil {
		return []xcodeproj.AddedFile{}
	}

	listCodeTemplates := codeTemplates.([]interface{})
	return generateTemplatesFromList(listCodeTemplates, templateName, moduleName, generateParams)
}

func generateTemplatesFromList(list []interface{}, templateName string, moduleName string, generateParams generateParams) []xcodeproj.AddedFile {
	addedFiles := []xcodeproj.AddedFile{}

	currentTime := time.Now()
	templateInfoParams := params(moduleName, generateParams.customKeys, generateParams.answers)
	params := templateInfoParams
	params["developer"] = MapKeys{
		"name":        generateParams.globalConfig.GetUsername(),
		"companyName": generateParams.projectConfig.GetCompanyName(),
	}
	params["date"] = currentTime.Format("02.01.2006")
	params["year"] = currentTime.Year()

	switch generateParams.projectConfig.GetProjectType() {
	case "iOS":
		params["projectName"] = generateParams.iosConfig.GetFolderNameCode()

		if generateParams.iosConfig.HasFolderNameUnitTests() {
			params["projectTestsName"] = generateParams.iosConfig.GetFolderNameUnitTests()
		} else {
			params["projectTestsName"] = params["projectName"]
		}

		if generateParams.iosConfig.HasFolderNameUITests() {
			params["projectUITestsName"] = generateParams.iosConfig.GetFolderNameUITests()
		} else {
			params["projectUITestsName"] = params["projectName"]
		}

	case "other":
		params["projectName"] = generateParams.otherConfig.GetProjectFolderName()
	default:
		break
	}

	templateFiles := newTemplateFiles(list,
		templateName,
		moduleName,
		generateParams.projectConfig.GetTemplatesFolderName(),
		params)
	for _, templateFile := range templateFiles {
		addedFiles = append(addedFiles, xcodeproj.AddedFile{
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
