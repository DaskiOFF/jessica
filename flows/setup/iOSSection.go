package setup

import (
	"strings"

	"github.com/daskioff/jessica/configs/models"
	"github.com/daskioff/jessica/utils/bundle"
	"github.com/daskioff/jessica/utils/files"
	"github.com/daskioff/jessica/utils/print"
	"github.com/daskioff/jessica/utils/question"
)

func (flow *SetupFlow) iosSection(config *models.ConfigIOS) {
	answer := question.AskQuestionWithBoolAnswer("Использовать Gemfile?")
	config.SetGemfileUse(answer)

	answer = question.AskQuestionWithBoolAnswer("Использовать Podfile?")
	config.SetPodfileUse(answer)

	xcodeprojFilename := question.AskQuestionWithChooseFileAnswer("Выберите .xcodeproj файл:", ".xcodeproj")
	if xcodeprojFilename == "" {
		print.PrintlnAttentionMessage("Пропущена настройка iOS проекта")
		return
	}
	config.SetXcodeprojFilename(xcodeprojFilename)
	config.SetProjectName(strings.Replace(xcodeprojFilename, ".xcodeproj", "", 1))

	codeProjectFolderName := question.AskQuestionWithChooseFolderAnswer("\nВыберите папку с кодом проекта (обычно она называется так же как и xcodeproj файл): ")
	if codeProjectFolderName == "" {
		codeProjectFolderName = "."
	}
	config.SetFolderNameCode(codeProjectFolderName)
	config.SetTargetNameCode(codeProjectFolderName)

	unitTestsFolderName := question.AskQuestionWithChooseFolderAnswer("\nВыберите папку с UNIT тестами: ")
	if unitTestsFolderName != "" {
		config.SetFolderNameUnitTests(unitTestsFolderName)
		config.SetTargetNameUnitTests(unitTestsFolderName)
	}

	uiTestsFolderName := question.AskQuestionWithChooseFolderAnswer("\nВыберите папку с UI тестами: ")
	if uiTestsFolderName != "" {
		config.SetFolderNameUITests(uiTestsFolderName)
		config.SetTargetNameUITests(uiTestsFolderName)
	}

	if config.GetGemfileUse() && files.IsFileExist("Gemfile") {
		bundle.Install()
	}
}
