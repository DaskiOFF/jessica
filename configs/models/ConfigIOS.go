package models

import (
	"errors"
	"strings"

	"github.com/daskioff/jessica/configs/keys"
	"github.com/spf13/viper"
)

type ConfigIOS struct {
	config *viper.Viper
}

func NewIOS(config *viper.Viper) *ConfigIOS {
	return &ConfigIOS{config}
}

func (c ConfigIOS) Validate() error {
	fields := []string{}

	if !c.HasGemfileUse() {
		fields = append(fields, keys.KeyIOSDependenciesGemfileUse)
	}
	if !c.HasPodfileUse() {
		fields = append(fields, keys.KeyIOSDependenciesPodfileUse)
	}
	if !c.HasXcodeprojFilename() {
		fields = append(fields, keys.KeyIOSXcodeprojFilename)
	}
	if !c.HasTargetNameCode() {
		fields = append(fields, keys.KeyIOSTargetNameCode)
	}
	if !c.HasFolderNameCode() {
		fields = append(fields, keys.KeyIOSFolderNameCode)
	}

	if len(fields) > 0 {
		return errors.New("Отсутствуют значения для некоторых полей в конфиг файле для проекта типа `iOS` (" + strings.Join(fields, ", ") + ")")
	}

	return nil
}

func (c ConfigIOS) Write() error {
	return c.config.WriteConfig()
}

// ------------

// Gemfile
func (c ConfigIOS) SetGemfileUse(value bool) {
	c.config.Set(keys.KeyIOSDependenciesGemfileUse, value)
}

func (c ConfigIOS) HasGemfileUse() bool {
	return c.config.IsSet(keys.KeyIOSDependenciesGemfileUse)
}

func (c ConfigIOS) GetGemfileUse() bool {
	return c.config.GetBool(keys.KeyIOSDependenciesGemfileUse)
}

// Podfile
func (c ConfigIOS) SetPodfileUse(value bool) {
	c.config.Set(keys.KeyIOSDependenciesPodfileUse, value)
}

func (c ConfigIOS) HasPodfileUse() bool {
	return c.config.IsSet(keys.KeyIOSDependenciesPodfileUse)
}

func (c ConfigIOS) GetPodfileUse() bool {
	return c.config.GetBool(keys.KeyIOSDependenciesPodfileUse)
}

// ProjectName
func (c ConfigIOS) SetProjectName(value string) {
	c.config.Set(keys.KeyIOSProjectName, value)
}

func (c ConfigIOS) HasProjectName() bool {
	return c.config.IsSet(keys.KeyIOSProjectName)
}

func (c ConfigIOS) GetProjectName() string {
	return c.config.GetString(keys.KeyIOSProjectName)
}

// XcodeprojFilename
func (c ConfigIOS) SetXcodeprojFilename(value string) {
	c.config.Set(keys.KeyIOSXcodeprojFilename, value)
}

func (c ConfigIOS) HasXcodeprojFilename() bool {
	return c.config.IsSet(keys.KeyIOSXcodeprojFilename)
}

func (c ConfigIOS) GetXcodeprojFilename() string {
	return c.config.GetString(keys.KeyIOSXcodeprojFilename)
}

// Target Folder Code
func (c ConfigIOS) SetTargetNameCode(value string) {
	c.config.Set(keys.KeyIOSTargetNameCode, value)
}

func (c ConfigIOS) HasTargetNameCode() bool {
	return c.config.IsSet(keys.KeyIOSTargetNameCode)
}

func (c ConfigIOS) GetTargetNameCode() string {
	return c.config.GetString(keys.KeyIOSTargetNameCode)
}

func (c ConfigIOS) SetFolderNameCode(value string) {
	c.config.Set(keys.KeyIOSFolderNameCode, value)
}

func (c ConfigIOS) HasFolderNameCode() bool {
	return c.config.IsSet(keys.KeyIOSFolderNameCode)
}

func (c ConfigIOS) GetFolderNameCode() string {
	return c.config.GetString(keys.KeyIOSFolderNameCode)
}

// Target Folder Unit Tests
func (c ConfigIOS) SetTargetNameUnitTests(value string) {
	c.config.Set(keys.KeyIOSTargetNameUnitTests, value)
}

func (c ConfigIOS) HasTargetNameUnitTests() bool {
	return c.config.IsSet(keys.KeyIOSTargetNameUnitTests)
}

func (c ConfigIOS) GetTargetNameUnitTests() string {
	return c.config.GetString(keys.KeyIOSTargetNameUnitTests)
}

func (c ConfigIOS) SetFolderNameUnitTests(value string) {
	c.config.Set(keys.KeyIOSFolderNameUnitTests, value)
}

func (c ConfigIOS) HasFolderNameUnitTests() bool {
	return c.config.IsSet(keys.KeyIOSFolderNameUnitTests)
}

func (c ConfigIOS) GetFolderNameUnitTests() string {
	return c.config.GetString(keys.KeyIOSFolderNameUnitTests)
}

// Target Folder UI Tests
func (c ConfigIOS) SetTargetNameUITests(value string) {
	c.config.Set(keys.KeyIOSTargetNameUITests, value)
}

func (c ConfigIOS) HasTargetNameUITests() bool {
	return c.config.IsSet(keys.KeyIOSTargetNameUITests)
}

func (c ConfigIOS) GetTargetNameUITests() string {
	return c.config.GetString(keys.KeyIOSTargetNameUITests)
}

func (c ConfigIOS) SetFolderNameUITests(value string) {
	c.config.Set(keys.KeyIOSFolderNameUITests, value)
}

func (c ConfigIOS) HasFolderNameUITests() bool {
	return c.config.IsSet(keys.KeyIOSFolderNameUITests)
}

func (c ConfigIOS) GetFolderNameUITests() string {
	return c.config.GetString(keys.KeyIOSFolderNameUITests)
}
