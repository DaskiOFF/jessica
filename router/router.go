package router

import (
	"errors"

	"github.com/daskioff/jessica/configs"
	"github.com/daskioff/jessica/configs/models"
	"github.com/daskioff/jessica/flows"
	"github.com/daskioff/jessica/flows/factory"
	"github.com/daskioff/jessica/utils/print"
)

const version = "1.3.4"

type Router struct {
	mapFlows map[string]flows.Flow

	globalConfig  *models.ConfigGlobal
	projectConfig *models.ConfigProject
	iosConfig     *models.ConfigIOS
}

func NewRouter() *Router {
	router := Router{}
	router.globalConfig = configs.Global()
	router.projectConfig = configs.Project()
	router.iosConfig = configs.IOS()

	mapFlows := make(map[string]flows.Flow)
	mapFlows["hi"] = factory.Hi(version)
	mapFlows["readme"] = factory.Readme(router.projectConfig, router.iosConfig)
	mapFlows["setup"] = factory.Setup(router.globalConfig, router.projectConfig, router.iosConfig)
	mapFlows["struct"] = factory.Struct(router.projectConfig, router.iosConfig)
	mapFlows["generator"] = factory.Generator(router.globalConfig, router.projectConfig, router.iosConfig)

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
		if command != "setup" && command != "hi" {
			globalError := r.globalConfig.Validate()
			projectError := r.projectConfig.Validate()
			iosError := r.iosConfig.Validate()

			if globalError == nil && projectError == nil && iosError == nil {
				return errors.New("\nДля начала необходимо настроить конфигурацию вызвав команду `jessica setup`")
			}
		}

		flow.Start(args[1:])
	}

	return nil
}
