package models

import (
	"errors"

	"github.com/daskioff/jessica/configs/keys"
	"github.com/spf13/viper"
)

type ConfigGlobal struct {
	config *viper.Viper
}

func NewGlobal(config *viper.Viper) *ConfigGlobal {
	return &ConfigGlobal{config}
}

func (c ConfigGlobal) Validate() error {
	if !c.config.IsSet(keys.KeyUserName) {
		return errors.New("Отсутствуют значения для некоторых полей в глобальном конфиг файле")
	}

	return nil
}

func (c ConfigGlobal) Write() error {
	return c.config.WriteConfig()
}

// ------------

func (c ConfigGlobal) SetUsername(value string) {
	c.config.Set(keys.KeyUserName, value)
}

func (c ConfigGlobal) HasUsername() bool {
	return c.config.IsSet(keys.KeyUserName)
}

func (c ConfigGlobal) GetUsername() string {
	return c.config.GetString(keys.KeyUserName)
}
