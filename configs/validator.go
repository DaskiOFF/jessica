package configs

import "errors"

func ValidateProjectConfig() error {
	config := ProjectConfig

	if !config.IsSet(KeyProjectName) ||
		!config.IsSet(KeyProjectXcodeProjName) {
		return errors.New("Отсутствуют значения для некоторых полей")
	}

	return nil
}
