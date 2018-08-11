package generator

import (
	"github.com/daskioff/jessica/configs/models"
	"github.com/daskioff/jessica/utils/print"

	"github.com/daskioff/jessica/flows"
	"github.com/daskioff/jessica/flows/generator/gen"
	"github.com/daskioff/jessica/flows/generator/list"
	"github.com/daskioff/jessica/flows/generator/pull"
	"github.com/daskioff/jessica/flows/generator/utils"
)

type MapKeys map[string]interface{}

type TemplateGeneratorFlow struct {
	globalConfig  *models.ConfigGlobal
	projectConfig *models.ConfigProject
	iosConfig     *models.ConfigIOS
	otherConfig   *models.ConfigOther
}

func (flow *TemplateGeneratorFlow) Start(args []string) {
	if len(args) == 0 {
		print.PrintlnAttentionMessage("Необходимо указать какое действие вы хотите выполнить. Чтобы увидеть список действий воспользуйтесь командой help")
		return
	}

	actionName := args[0]
	args = args[1:]

	if actionName == "pull" {
		pull.Execute(args, flow.projectConfig.GetTemplatesFolderName())
		return
	}

	templatesFolderName := flow.projectConfig.GetTemplatesFolderName()
	templates := searchTemplates(templatesFolderName, gen.TemplateDescriptionFileName)
	if len(templates) == 0 {
		print.PrintlnErrorMessage("В папке " + templatesFolderName + " шаблоны не найдены")
		return
	}

	switch actionName {
	case "list":
		list.Show(templates)
	case "gen":
		gen.Execute(args, utils.TemplatesRootPath(templatesFolderName), flow.globalConfig, flow.projectConfig, flow.iosConfig, flow.otherConfig)
	}
}

func (flow *TemplateGeneratorFlow) Description() string {
	return `--------------------------------------------------------------------------------
Генератор
  - pull  – Скачать шаблоны с репозитория
  - list  – Вывести список шаблонов
  - gen   – Генерация шаблона
--------------------------------------------------------------------------------`
}

// ----------------------------------------------------------------------------
func NewFlow(globalConfig *models.ConfigGlobal, projectConfig *models.ConfigProject, iosConfig *models.ConfigIOS, otherConfig *models.ConfigOther) flows.Flow {
	flow := TemplateGeneratorFlow{}
	flow.globalConfig = globalConfig
	flow.projectConfig = projectConfig
	flow.iosConfig = iosConfig
	flow.otherConfig = otherConfig

	return &flow
}
