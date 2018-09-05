package models

import (
	"errors"
	"strings"

	"github.com/daskioff/jessica/configs/keys"
	"github.com/spf13/viper"
)

const (
	// ConfigProjectTypeUnknown is `unknown` project type
	ConfigProjectTypeUnknown int = iota
	// ConfigProjectTypeIOS is `iOS` project type
	ConfigProjectTypeIOS
	// ConfigProjectTypeOther is `Other` project type
	ConfigProjectTypeOther
)

type ConfigProject struct {
	config *viper.Viper
}

func NewProject(config *viper.Viper) *ConfigProject {
	return &ConfigProject{config}
}

func (c ConfigProject) Validate() error {
	fields := []string{}

	if !c.HasCompanyName() {
		fields = append(fields, keys.KeyCompanyName)
	}
	if !c.HasProjectType() {
		fields = append(fields, keys.KeyProjectType)
	}
	if !c.HasReadmeTemplateFilename() {
		fields = append(fields, keys.KeyReadmeTemplateFilename)
	}
	if !c.HasCustomProjectStructUse() {
		fields = append(fields, keys.KeyCustomProjectStructUse)
	}
	if !c.HasCustomProjectStructDescriptionTemplateFilename() {
		fields = append(fields, keys.KeyCustomProjectStructDescriptionTemplateFilename)
	}
	if !c.HasTemplatesUse() {
		fields = append(fields, keys.KeyTemplatesUse)
	}
	if !c.HasTemplatesFolderName() {
		fields = append(fields, keys.KeyTemplatesFolderName)
	}

	if len(fields) > 0 {
		return errors.New("Отсутствуют значения для некоторых полей в конфиг файле проекта (" + strings.Join(fields, ", ") + ")")
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
// SetProjectType setup project type from `value` (ConfigProjectType) to config
func (c ConfigProject) SetProjectType(value int) {
	switch value {
	case ConfigProjectTypeIOS:
		c.config.Set(keys.KeyProjectType, "iOS")
	case ConfigProjectTypeOther:
		c.config.Set(keys.KeyProjectType, "other")
	case ConfigProjectTypeUnknown:
		c.config.Set(keys.KeyProjectType, nil)
	}
}

func (c ConfigProject) HasProjectType() bool {
	return c.config.IsSet(keys.KeyProjectType)
}

// GetProjectType return ConfigProjectType
func (c ConfigProject) GetProjectType() int {
	typeString := strings.ToLower(c.config.GetString(keys.KeyProjectType))

	switch typeString {
	case "ios":
		return ConfigProjectTypeIOS
	case "other":
		return ConfigProjectTypeOther
	}

	return ConfigProjectTypeUnknown
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
