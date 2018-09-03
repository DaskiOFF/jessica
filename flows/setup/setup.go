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
	flow.globalConfig.Write()

	project.Setup(flow.projectConfig, isForce)
	flow.projectConfig.Write()

	switch flow.projectConfig.GetProjectType() {
	case models.ConfigProjectTypeIOS:
		ios.Setup(flow.iosConfig, isForce)
		flow.iosConfig.Write()
	case models.ConfigProjectTypeOther:
		other.Setup(flow.otherConfig, isForce)
		flow.otherConfig.Write()
	}

	print.PrintlnSuccessMessage("Файл сконфигурирован")
}
