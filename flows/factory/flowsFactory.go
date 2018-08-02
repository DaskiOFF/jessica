package factory

import (
	"github.com/daskioff/jessica/configs/models"
	"github.com/daskioff/jessica/flows"
	"github.com/daskioff/jessica/flows/hi"
	"github.com/daskioff/jessica/flows/projectstruct"
	"github.com/daskioff/jessica/flows/readme"
	"github.com/daskioff/jessica/flows/setup"
	"github.com/daskioff/jessica/flows/templategenerator"
)

func Hi(version string) flows.Flow {
	return hi.NewFlow(version)
}

func Struct(projectConfig *models.ConfigProject, iosConfig *models.ConfigIOS) flows.Flow {
	return projectstruct.NewFlow(projectConfig, iosConfig)
}

func Readme(projectConfig *models.ConfigProject, iosConfig *models.ConfigIOS) flows.Flow {
	return readme.NewFlow(projectConfig, iosConfig)
}

func Setup(globalConfig *models.ConfigGlobal, projectConfig *models.ConfigProject, iosConfig *models.ConfigIOS, otherConfig *models.ConfigOther) flows.Flow {
	return setup.NewFlow(globalConfig, projectConfig, iosConfig, otherConfig)
}

func Generator(globalConfig *models.ConfigGlobal, projectConfig *models.ConfigProject, iosConfig *models.ConfigIOS) flows.Flow {
	return templategenerator.NewFlow(globalConfig, projectConfig, iosConfig)
}
