package templategenerator

import (
	"errors"

	"github.com/daskioff/jessica/utils/files"
)

func (flow *TemplateGeneratorFlow) Validate() error {
	templates := flow.searchTemplates()

	if flow.projectConfig.GetProjectType() == "iOS" {
		if !flow.iosConfig.GetGemfileUse() || !files.IsFileExist("Gemfile.lock") {
			return errors.New("Вы не используете Gemfile. Для использования генератора с iOS проектом необходим Gemfile с зависимостью `gem 'xcodeproj'`, для добавления сгенерированных файлов в xcodeproj")
		}
		// TODO Проверять наличие Gemfile.lock
		// TODO Проверять в файле Gemfile.lock наличие зависимости xcodeproj
	}

	if len(templates) == 0 {
		return errors.New("Шаблоны не найдены")
	}

	return nil
}
