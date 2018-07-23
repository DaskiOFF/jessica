package templategenerator

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
	textTemplate "text/template"

	"github.com/spf13/viper"

	"github.com/daskioff/jessica/configs"
	"github.com/daskioff/jessica/flows"
	"github.com/daskioff/jessica/utils"
)

const TemplateExt = ".tpl"
const TemplateDescriptionFileName = "templates.yml"

type TemplateGeneratorFlow struct {
}

func (flow *TemplateGeneratorFlow) Start(args []string) {
	if len(args) == 0 {
		utils.PrintlnAttentionMessage("Необходимо указать какое действие вы хотите выполнить. Чтобы увидеть список действий воспользуйтесь командой help")
		return
	}

	err := Validate()
	if err != nil {
		utils.PrintlnErrorMessage(err.Error())
		return
	}

	templates := searchTemplates()
	if args[0] == "list" {
		if len(templates) == 0 {
			utils.PrintlnAttentionMessage("Шаблоны не найдены")
		} else {
			list := ""
			for _, template := range templates {
				if len(list) == 0 {
					list = template
				} else {
					list = list + "\n" + template
				}
			}
			utils.PrintlnInfoMessage(list)
		}
	} else if args[0] == "gen" {
		if len(args) == 1 {
			utils.PrintlnAttentionMessage("Не указано имя шаблона для генерации")
		} else {
			templatePath := filepath.Join(templatesRootPath(), args[1], TemplateDescriptionFileName)
			if !utils.IsFileExist(templatePath) {
				utils.PrintlnErrorMessage("Шаблон с данным именем не найден")
			} else if len(args) < 3 {
				utils.PrintlnErrorMessage("Не указано имя генерируемого модуля")
			} else {
				v := viper.New()
				v.SetConfigFile(templatePath)

				err = v.ReadInConfig()
				if err != nil {
					utils.PrintlnErrorMessage(err.Error())
					return
				}

				fmt.Println(v.AllSettings())
				generateTemplates(v, args[1], args[2])
			}
		}
	}
}

func (flow *TemplateGeneratorFlow) Description() string {
	return `
--------------------------------------------------------------------------------
	Генерация шаблонов
--------------------------------------------------------------------------------`
}

// ----------------------------------------------------------------------------
func NewFlow() flows.Flow {
	flow := TemplateGeneratorFlow{}
	return &flow
}

func generateTemplates(v *viper.Viper, templateName string, moduleName string) {
	codeTemplates := v.Get("code_files")
	listCodeTemplates := codeTemplates.([]interface{})
	for _, codeInterface := range listCodeTemplates {
		code := codeInterface.(map[interface{}]interface{})
		root, err := os.Getwd()
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		resultFileName := moduleName + code["name"].(string)

		templateFileName := code["template_path"].(string)
		templateFileName = strings.Replace(templateFileName, "{{.projectName}}", configs.ProjectConfig.GetString(configs.KeyProjectName), -1)
		templateFileName = strings.Replace(templateFileName, "{{.moduleName}}", moduleName, -1)
		templateFileName = filepath.Join(templatesRootPath(), templateName, templateFileName)

		filePath := code["output_path"].(string)
		filePath = strings.Replace(filePath, "{{.projectName}}", configs.ProjectConfig.GetString(configs.KeyProjectName), -1)
		filePath = strings.Replace(filePath, "{{.moduleName}}", moduleName, -1)
		filePath = filepath.Join(root, filePath)

		err = os.MkdirAll(filePath, os.ModePerm)
		if err != nil {
			panic(err)
		}

		fullFilePath := filepath.Join(filePath, resultFileName)

		// TODO спрашивать перезаписывать ли файл при его существовании
		file, err := os.OpenFile(fullFilePath, os.O_CREATE|os.O_WRONLY, os.ModePerm)
		if err != nil {
			panic(err)
		}
		defer file.Close()

		writer := bufio.NewWriter(file)
		params := map[string]interface{}{
			"xcodeVersion":        7,
			"swiftVersion":        3,
			"gemFileDependencies": 4,
			"podFileDependencies": 5,
			"projectName":         configs.ProjectConfig.Get(configs.KeyProjectName),
		}

		err = executeTemplate(templateFileName, writer, params)
		if err != nil {
			panic(err)
		}

		err = writer.Flush()
		if err != nil {
			panic(err)
		}

		utils.PrintlnSuccessMessage(filePath + " successfully created")
	}
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
