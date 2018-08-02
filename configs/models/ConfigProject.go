package models

import (
	"errors"

	"github.com/daskioff/jessica/configs/keys"
	"github.com/spf13/viper"
)

type ConfigProject struct {
	config *viper.Viper
}

func NewProject(config *viper.Viper) *ConfigProject {
	return &ConfigProject{config}
}

func (c ConfigProject) Validate() error {
	config := c.config

	if !config.IsSet(keys.KeyCompanyName) ||
		!config.IsSet(keys.KeyProjectType) ||
		!config.IsSet(keys.KeyReadmeTemplateFilename) ||
		!config.IsSet(keys.KeyCustomProjectStructUse) ||
		!config.IsSet(keys.KeyCustomProjectStructDescriptionTemplateFilename) ||
		!config.IsSet(keys.KeyTemplatesUse) ||
		!config.IsSet(keys.KeyTemplatesFolderName) {

		return errors.New("Отсутствуют значения для некоторых полей в конфиг файле проекта")
	}

	return nil
}

func (c ConfigProject) Write() error {
	return c.config.WriteConfig()
}

// ------------

// Company name
func (c ConfigProject) SetCompanyName(value string) {
	c.config.Set(keys.KeyCompanyName, value)
}

func (c ConfigProject) HasCompanyName() bool {
	return c.config.IsSet(keys.KeyCompanyName)
}

func (c ConfigProject) GetCompanyName() string {
	return c.config.GetString(keys.KeyCompanyName)
}

// Project type
func (c ConfigProject) SetProjectType(value string) {
	c.config.Set(keys.KeyProjectType, value)
}

func (c ConfigProject) HasProjectType() bool {
	return c.config.IsSet(keys.KeyProjectType)
}

func (c ConfigProject) GetProjectType() string {
	return c.config.GetString(keys.KeyProjectType)
}

// Readme
func (c ConfigProject) SetReadmeTemplateFilename(value string) {
	c.config.Set(keys.KeyReadmeTemplateFilename, value)
}

func (c ConfigProject) HasReadmeTemplateFilename() bool {
	return c.config.IsSet(keys.KeyReadmeTemplateFilename)
}

func (c ConfigProject) GetReadmeTemplateFilename() string {
	return c.config.GetString(keys.KeyReadmeTemplateFilename)
}

// Custom project struct
func (c ConfigProject) SetCustomProjectStructUse(value bool) {
	c.config.Set(keys.KeyCustomProjectStructUse, value)
}

func (c ConfigProject) HasCustomProjectStructUse() bool {
	return c.config.IsSet(keys.KeyCustomProjectStructUse)
}

func (c ConfigProject) GetCustomProjectStructUse() bool {
	return c.config.GetBool(keys.KeyCustomProjectStructUse)
}

func (c ConfigProject) HasCustomProjectStructDescription() bool {
	return c.config.IsSet(keys.KeyCustomProjectStructDescription)
}

func (c ConfigProject) GetCustomProjectStructDescription() interface{} {
	return c.config.Get(keys.KeyCustomProjectStructDescription)
}

func (c ConfigProject) SetCustomProjectStructDescriptionTemplateFilename(value string) {
	c.config.Set(keys.KeyCustomProjectStructDescriptionTemplateFilename, value)
}

func (c ConfigProject) HasCustomProjectStructDescriptionTemplateFilename() bool {
	return c.config.IsSet(keys.KeyCustomProjectStructDescriptionTemplateFilename)
}

func (c ConfigProject) GetCustomProjectStructDescriptionTemplateFilename() string {
	return c.config.GetString(keys.KeyCustomProjectStructDescriptionTemplateFilename)
}

// Templates
func (c ConfigProject) SetTemplatesUse(value bool) {
	c.config.Set(keys.KeyTemplatesUse, value)
}

func (c ConfigProject) HasTemplatesUse() bool {
	return c.config.IsSet(keys.KeyTemplatesUse)
}

func (c ConfigProject) GetTemplatesUse() bool {
	return c.config.GetBool(keys.KeyTemplatesUse)
}

func (c ConfigProject) SetTemplatesFolderName(value string) {
	c.config.Set(keys.KeyTemplatesFolderName, value)
}

func (c ConfigProject) HasTemplatesFolderName() bool {
	return c.config.IsSet(keys.KeyTemplatesFolderName)
}

func (c ConfigProject) GetTemplatesFolderName() string {
	return c.config.GetString(keys.KeyTemplatesFolderName)
}
