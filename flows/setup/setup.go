package setup

import (
	"github.com/daskioff/jessica/configs/models"
	"github.com/daskioff/jessica/flows/setup/global"
	"github.com/daskioff/jessica/flows/setup/ios"
	"github.com/daskioff/jessica/flows/setup/other"
	"github.com/daskioff/jessica/flows/setup/project"
	"github.com/daskioff/jessica/utils/print"
	"github.com/daskioff/jessica/utils/slices"
)

type SetupFlow struct {
	globalConfig  *models.ConfigGlobal
	projectConfig *models.ConfigProject
	iosConfig     *models.ConfigIOS
	otherConfig   *models.ConfigOther
}

func (flow *SetupFlow) Start(args []string) {
	isForce := slices.Contains(args, "--force") || slices.Contains(args, "--f")

	flow.setup(isForce)
}

func (flow *SetupFlow) Description() string {
	return `--------------------------------------------------------------------------------
Первичная настройка файла конфигурации

  Параметры:
    --force, --f – Обновить всю конфигурацию
--------------------------------------------------------------------------------
	`
}

// ----------------------------------------------------------------------------
func New(globalConfig *models.ConfigGlobal, projectConfig *models.ConfigProject, iosConfig *models.ConfigIOS, otherConfig *models.ConfigOther) *SetupFlow {
	flow := SetupFlow{}
	flow.globalConfig = globalConfig
	flow.projectConfig = projectConfig
	flow.iosConfig = iosConfig
	flow.otherConfig = otherConfig

	return &flow
}

func (flow *SetupFlow) setup(isForce bool) {
	global.Setup(flow.globalConfig, isForce)
	project.Setup(flow.projectConfig, isForce)

	switch flow.projectConfig.GetProjectType() {
	case "iOS":
		ios.Setup(flow.iosConfig, isForce)
	case "Other":
		other.Setup(flow.otherConfig, isForce)
	}

	flow.globalConfig.Write()
	flow.projectConfig.Write()
	flow.iosConfig.Write()
	flow.otherConfig.Write()

	print.PrintlnSuccessMessage("Файл сконфигурирован")
}
