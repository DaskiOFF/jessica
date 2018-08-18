package factory

import (
	"github.com/daskioff/jessica/configs/models"
	"github.com/daskioff/jessica/flows"
	"github.com/daskioff/jessica/flows/generator"
	"github.com/daskioff/jessica/flows/hi"
	"github.com/daskioff/jessica/flows/projectstruct"
	"github.com/daskioff/jessica/flows/readme"
	"github.com/daskioff/jessica/flows/setup"
)

func Hi(version string) flows.Flow {
	return hi.New(version)
}

func Struct(projectConfig *models.ConfigProject, iosConfig *models.ConfigIOS, otherConfig *models.ConfigOther) flows.Flow {
	return projectstruct.New(projectConfig, iosConfig, otherConfig)
}

func Readme(projectConfig *models.ConfigProject, iosConfig *models.ConfigIOS, otherConfig *models.ConfigOther) flows.Flow {
	return readme.New(projectConfig, iosConfig, otherConfig)
}

func Setup(globalConfig *models.ConfigGlobal, projectConfig *models.ConfigProject, iosConfig *models.ConfigIOS, otherConfig *models.ConfigOther) flows.Flow {
	return setup.New(globalConfig, projectConfig, iosConfig, otherConfig)
}

func Generator(globalConfig *models.ConfigGlobal, projectConfig *models.ConfigProject, iosConfig *models.ConfigIOS, otherConfig *models.ConfigOther) flows.Flow {
	return generator.New(globalConfig, projectConfig, iosConfig, otherConfig)
}
