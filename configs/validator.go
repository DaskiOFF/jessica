package configs

import "errors"

func ValidateProjectConfig() error {
	config := ProjectConfig

	if !config.IsSet(KeyProjectName) ||
		!config.IsSet(KeyProjectXcodeProjName) ||
		!config.IsSet(KeyCompanyName) ||
		!config.IsSet(KeyProjectTestsFolderName) {
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
