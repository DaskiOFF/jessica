package templategenerator

import (
	"errors"
)

func (flow *TemplateGeneratorFlow) Validate() error {
	templates := flow.searchTemplates()

	if len(templates) == 0 {
		return errors.New("Шаблоны не найдены")
	}

	return nil
}
