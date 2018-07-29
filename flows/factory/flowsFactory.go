package factory

import (
	"github.com/daskioff/jessica/flows"
	"github.com/daskioff/jessica/flows/hi"
	projectstruct "github.com/daskioff/jessica/flows/projectStruct"
	"github.com/daskioff/jessica/flows/readme"
	"github.com/daskioff/jessica/flows/setup"
	"github.com/daskioff/jessica/flows/templategenerator"
)

func Hi(version string) flows.Flow {
	return hi.NewFlow(version)
}

func Struct() flows.Flow {
	return projectstruct.NewFlow()
}

func Readme() flows.Flow {
	return readme.NewFlow()
}

func Setup() flows.Flow {
	return setup.NewFlow()
}

func Generator() flows.Flow {
	return templategenerator.NewFlow()
}
