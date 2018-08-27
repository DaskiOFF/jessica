package factory

import (
	"github.com/daskioff/jessica/configs/models"
	"github.com/daskioff/jessica/flows"
	"github.com/daskioff/jessica/flows/generator"
	"github.com/daskioff/jessica/flows/hi"
	"github.com/daskioff/jessica/flows/projectstruct"
	readmeIOS "github.com/daskioff/jessica/flows/readme/ios"
	readmeOther "github.com/daskioff/jessica/flows/readme/other"
	"github.com/daskioff/jessica/flows/setup"
)

// Hi flow
func Hi(version string) flows.Flow {
	return hi.New(version)
}

// Setup flow
func Setup(globalConfig *models.ConfigGlobal, projectConfig *models.ConfigProject, iosConfig *models.ConfigIOS, otherConfig *models.ConfigOther) flows.Flow {
	return setup.New(globalConfig, projectConfig, iosConfig, otherConfig)
}

// Struct flow depending on the project type
func Struct(projectConfig *models.ConfigProject, iosConfig *models.ConfigIOS, otherConfig *models.ConfigOther) flows.Flow {
	return projectstruct.New(projectConfig, iosConfig, otherConfig)
}

// Readme flow depending on the project type
func Readme(projectConfig *models.ConfigProject, iosConfig *models.ConfigIOS, otherConfig *models.ConfigOther) flows.Flow {
	if isIOSProject(projectConfig) {
		return readmeIOS.New(projectConfig, iosConfig)
	}

	return readmeOther.New(projectConfig, otherConfig)
}

// Generator flow depending on the project type
func Generator(globalConfig *models.ConfigGlobal, projectConfig *models.ConfigProject, iosConfig *models.ConfigIOS, otherConfig *models.ConfigOther) flows.Flow {
	return generator.New(globalConfig, projectConfig, iosConfig, otherConfig)
}

// Private
func isIOSProject(projectConfig *models.ConfigProject) bool {
	return projectConfig.GetProjectType() == "iOS"
}
