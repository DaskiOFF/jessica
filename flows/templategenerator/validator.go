package templategenerator

import (
	"errors"
	"strings"
)

func Validate() error {
	templates := searchTemplates()

	for k, v := range templates {
		if len(v) > 1 {
			return errors.New("TemplateGenerator. Duplicate Templates with name: '" + k + "'. In paths [" + strings.Join(v, "], [") + "].")
		}
	}

	return nil
}
