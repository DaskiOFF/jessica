package templategenerator

import (
	"fmt"

	"github.com/daskioff/jessica/flows"
	"github.com/daskioff/jessica/utils"
)

const TemplateExt = ".tpl"

type TemplateGeneratorFlow struct {
}

func (flow *TemplateGeneratorFlow) Start(args []string) {
	err := Validate()
	if err != nil {
		utils.PrintlnErrorMessage(err.Error())
		return
	}

	templates := searchTemplates()
	if len(templates) == 0 {
		utils.PrintlnAttentionMessage("Шаблоны не найдены")
	} else {
		fmt.Println(templates)
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
