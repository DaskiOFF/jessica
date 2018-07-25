package configs

import "errors"

func ValidateProjectConfig() error {
	config := ProjectConfig

	if !config.IsSet(KeyCompanyName) ||
		!config.IsSet(KeyProjectType) ||
		!config.IsSet(KeyReadmeTemplateFilename) ||
		!config.IsSet(KeyCustomProjectStructUse) ||
		!config.IsSet(KeyCustomProjectStructDescription) ||
		!config.IsSet(KeyCustomProjectStructDescriptionTemplateFilename) ||
		!config.IsSet(KeyTemplatesUse) ||
		!config.IsSet(KeyTemplatesFolderName) {

		if config.GetString(KeyProjectType) == "iOS" {
			if !config.IsSet(KeyIOSDependenciesGemfileUse) ||
				!config.IsSet(KeyIOSDependenciesPodfileUse) ||
				!config.IsSet(KeyIOSXcodeprojFilename) ||
				!config.IsSet(KeyIOSTargetnameCode) ||
				!config.IsSet(KeyIOSTargetnameUnitTests) ||
				!config.IsSet(KeyIOSTargetnameUITests) ||
				!config.IsSet(KeyIOSFolderNameCode) ||
				!config.IsSet(KeyIOSFolderNameUnitTests) ||
				!config.IsSet(KeyIOSFolderNameUITests) {

				return errors.New("Отсутствуют значения для некоторых полей в конфиг файле для iOS проекта")
			}
		}
		return errors.New("Отсутствуют значения для некоторых полей в конфиг файле проекта")
	}

	err := validateGlobalConfig()

	return err
}

func validateGlobalConfig() error {
	config := GlobalConfig

	if !config.IsSet(KeyUserName) {
		return errors.New("Отсутствуют значения для некоторых полей в глобальном конфиг файле")
	}

	return nil
}
