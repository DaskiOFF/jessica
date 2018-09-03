package projectstruct

import (
	"github.com/daskioff/jessica/configs/models"
	"github.com/daskioff/jessica/utils/print"
)

var useCustomStruct bool
var hasCustomStruct bool

type ProjectStructFlow struct {
	projectConfig *models.ConfigProject
	iosConfig     *models.ConfigIOS
	otherConfig   *models.ConfigOther
}

func (flow *ProjectStructFlow) Start(args []string) {
	if len(args) == 0 {
		print.PrintlnAttentionMessage("Необходимо указать какое действие вы хотите выполнить. Чтобы увидеть список действий воспользуйтесь командой help")
		return
	}

	actionName := args[0]
	args = args[1:]

	switch actionName {
	case "gen":
		flow.gen()
	case "example":
		printExample()
	default:
		print.PrintlnAttentionMessage("Действие не найдено. Чтобы увидеть список действий воспользуйтесь командой help")
	}
}

func (flow *ProjectStructFlow) Description() string {
	return `--------------------------------------------------------------------------------
Создание и описание структуры проекта.

  Действия:
    gen     – Генерация структуры файлов описаных в конфиге
    example – Пример описания файловой структуры в файле конфигурации
--------------------------------------------------------------------------------`
}

// ----------------------------------------------------------------------------
func New(projectConfig *models.ConfigProject, iosConfig *models.ConfigIOS, otherConfig *models.ConfigOther) *ProjectStructFlow {
	flow := ProjectStructFlow{}
	flow.projectConfig = projectConfig
	flow.iosConfig = iosConfig
	flow.otherConfig = otherConfig

	return &flow
}
