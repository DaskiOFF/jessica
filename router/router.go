package router

import (
	"github.com/daskioff/jessica/configs"
	"github.com/daskioff/jessica/configs/models"
	"github.com/daskioff/jessica/flows"
	"github.com/daskioff/jessica/flows/factory"
	"github.com/daskioff/jessica/utils/print"
)

const version = "1.4.1"

type Router struct {
	mapFlows map[string]flows.Flow

	globalConfig  *models.ConfigGlobal
	projectConfig *models.ConfigProject
	iosConfig     *models.ConfigIOS
	otherConfig   *models.ConfigOther
}

func New() *Router {
	router := Router{}
	router.globalConfig = configs.Global()
	router.projectConfig = configs.Project()
	router.iosConfig = configs.IOS()
	router.otherConfig = configs.Other()

	mapFlows := make(map[string]flows.Flow)
	mapFlows["readme"] = factory.Readme(router.projectConfig, router.iosConfig, router.otherConfig)
	mapFlows["setup"] = factory.Setup(router.globalConfig, router.projectConfig, router.iosConfig, router.otherConfig)
	mapFlows["struct"] = factory.Struct(router.projectConfig, router.iosConfig, router.otherConfig)
	mapFlows["generator"] = factory.Generator(router.globalConfig, router.projectConfig, router.iosConfig, router.otherConfig)

	router.mapFlows = mapFlows

	return &router
}

// Handle Обрабатывает аргументы переданные программе и направляет на нужный flow
func (r *Router) Handle(args []string) error {
	if len(args) == 0 {
		return errNoArguments
	}

	command := args[0]
	if command == "version" {
		print.PrintlnInfoMessage(version)
		return nil
	}

	isHelp := false
	if command == "help" {
		if len(args) < 2 {
			return errHelpNoArguments
		}
		isHelp = true
		command = args[1]
	}

	flow, ok := r.mapFlows[command]
	if !ok {
		return errRouteNotFound
	}

	if isHelp {
		print.PrintlnInfoMessage(flow.Description())
	} else {
		if command != "setup" {
			if err := r.validateConfigs(); err != nil {
				return err
			}
		}

		flow.Start(args[1:])
	}

	return nil
}
