package ios

import (
	"strings"

	"github.com/daskioff/jessica/configs/models"
	"github.com/daskioff/jessica/utils/print"
	"github.com/daskioff/jessica/utils/question"
)

func Setup(config *models.ConfigIOS, isForce bool) {
	err := config.Validate()
	if err == nil && !isForce {
		return
	}

	if !config.HasGemfileUse() || isForce {
		answer := question.AskQuestionWithBoolAnswer("Использовать Gemfile?")
		config.SetGemfileUse(answer)

		if answer {
			checkGemfile()
		}
	}

	if !config.HasPodfileUse() || isForce {
		answer := question.AskQuestionWithBoolAnswer("Использовать Podfile?")
		config.SetPodfileUse(answer)

		if answer {
			checkPodfile()
		}
	}

	if !config.HasProjectName() || !config.HasXcodeprojFilename() || isForce {
		xcodeprojFilename := question.AskQuestionWithChooseFileAnswer("Выберите .xcodeproj файл:", ".xcodeproj")
		if xcodeprojFilename == "" {
			print.PrintlnAttentionMessage("Пропущена настройка iOS проекта")
			return
		}
		config.SetXcodeprojFilename(xcodeprojFilename)
		config.SetProjectName(strings.Replace(xcodeprojFilename, ".xcodeproj", "", 1))
	}

	if !config.HasTargetNameCode() || !config.HasFolderNameCode() || isForce {
		codeProjectFolderName := question.AskQuestionWithChooseFolderAnswer("\nВыберите папку с кодом проекта (обычно она называется так же как и xcodeproj файл): ")
		if codeProjectFolderName == "" {
			codeProjectFolderName = "."
		}
		config.SetFolderNameCode(codeProjectFolderName)
		config.SetTargetNameCode(codeProjectFolderName)
	}

	if !config.HasTargetNameUnitTests() || !config.HasFolderNameUnitTests() || isForce {
		unitTestsFolderName := question.AskQuestionWithChooseFolderAnswer("\nВыберите папку с UNIT тестами: ")
		if unitTestsFolderName != "" {
			config.SetFolderNameUnitTests(unitTestsFolderName)
			config.SetTargetNameUnitTests(unitTestsFolderName)
		}
	}

	if !config.HasTargetNameUITests() || !config.HasFolderNameUITests() || isForce {
		uiTestsFolderName := question.AskQuestionWithChooseFolderAnswer("\nВыберите папку с UI тестами: ")
		if uiTestsFolderName != "" {
			config.SetFolderNameUITests(uiTestsFolderName)
			config.SetTargetNameUITests(uiTestsFolderName)
		}
	}
}
