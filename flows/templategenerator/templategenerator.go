package templategenerator

import (
	"fmt"
	"path/filepath"

	"github.com/spf13/viper"

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
			} else {
				v := viper.New()
				v.SetConfigFile(templatePath)

				err = v.ReadInConfig()
				if err != nil {
					utils.PrintlnErrorMessage(err.Error())
					return
				}

				fmt.Println(v.AllSettings())
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
