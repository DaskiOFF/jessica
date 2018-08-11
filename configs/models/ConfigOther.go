package models

import (
	"errors"
	"strings"

	"github.com/daskioff/jessica/configs/keys"
	"github.com/spf13/viper"
)

type ConfigOther struct {
	config *viper.Viper
}

func NewOther(config *viper.Viper) *ConfigOther {
	return &ConfigOther{config}
}

func (c ConfigOther) Validate() error {
	fields := []string{}

	if !c.HasProjectFolderName() {
		fields = append(fields, keys.KeyOtherProjectFolderName)
	}

	if !c.HasProjectName() {
		fields = append(fields, keys.KeyOtherProjectName)
	}

	if len(fields) > 0 {
		return errors.New("Отсутствуют значения для некоторых полей в конфиг файле для проекта типа `Other` (" + strings.Join(fields, ", ") + ")")
	}

	return nil
}

func (c ConfigOther) Write() error {
	return c.config.WriteConfig()
}

// ------------

// Project Name
func (c ConfigOther) SetProjectName(value string) {
	c.config.Set(keys.KeyOtherProjectName, value)
}

func (c ConfigOther) HasProjectName() bool {
	return c.config.IsSet(keys.KeyOtherProjectName)
}

func (c ConfigOther) GetProjectName() string {
	return c.config.GetString(keys.KeyOtherProjectName)
}

// Project Folder
func (c ConfigOther) SetProjectFolderName(value string) {
	c.config.Set(keys.KeyOtherProjectFolderName, value)
}

func (c ConfigOther) HasProjectFolderName() bool {
	return c.config.IsSet(keys.KeyOtherProjectFolderName)
}

func (c ConfigOther) GetProjectFolderName() string {
	return c.config.GetString(keys.KeyOtherProjectFolderName)
}
