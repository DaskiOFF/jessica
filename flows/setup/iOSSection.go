package setup

import (
	"strings"

	"github.com/daskioff/jessica/configs"
	"github.com/daskioff/jessica/utils/bundle"
	"github.com/daskioff/jessica/utils/files"
	"github.com/daskioff/jessica/utils/print"
	"github.com/daskioff/jessica/utils/question"
	"github.com/spf13/viper"
)

func iosSection(config *viper.Viper) {
	answer := question.AskQuestionWithBoolAnswer("Использовать Gemfile?")
	config.Set(configs.KeyIOSDependenciesGemfileUse, answer)

	answer = question.AskQuestionWithBoolAnswer("Использовать Podfile?")
	config.Set(configs.KeyIOSDependenciesPodfileUse, answer)

	xcodeprojFilename := question.AskQuestionWithChooseFileAnswer("Выберите .xcodeproj файл:", ".xcodeproj")
	if xcodeprojFilename == "" {
		print.PrintlnAttentionMessage("Пропущена настройка iOS проекта")
		return
	}
	config.Set(configs.KeyIOSXcodeprojFilename, xcodeprojFilename)
	config.Set(configs.KeyIOSProjectName, strings.Replace(xcodeprojFilename, ".xcodeproj", "", 1))

	codeProjectFolderName := question.AskQuestionWithChooseFolderAnswer("\nВыберите папку с кодом проекта (обычно она называется так же как и xcodeproj файл): ")
	if codeProjectFolderName == "" {
		codeProjectFolderName = "."
	}
	config.Set(configs.KeyIOSFolderNameCode, codeProjectFolderName)
	config.Set(configs.KeyIOSTargetnameCode, codeProjectFolderName)

	unitTestsFolderName := question.AskQuestionWithChooseFolderAnswer("\nВыберите папку с UNIT тестами: ")
	if unitTestsFolderName != "" {
		config.Set(configs.KeyIOSFolderNameUnitTests, unitTestsFolderName)
		config.Set(configs.KeyIOSTargetnameUnitTests, unitTestsFolderName)
	}

	uiTestsFolderName := question.AskQuestionWithChooseFolderAnswer("\nВыберите папку с UI тестами: ")
	if uiTestsFolderName != "" {
		config.Set(configs.KeyIOSFolderNameUITests, uiTestsFolderName)
		config.Set(configs.KeyIOSTargetnameUITests, uiTestsFolderName)
	}

	if config.GetBool(configs.KeyIOSDependenciesGemfileUse) && files.IsFileExist("Gemfile") {
		bundle.Install()
	}
}
