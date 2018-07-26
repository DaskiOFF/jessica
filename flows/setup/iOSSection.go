package setup

import (
	"strings"

	"github.com/daskioff/jessica/configs"
	"github.com/daskioff/jessica/utils"
	"github.com/spf13/viper"
)

func iosSection(config *viper.Viper) {
	answer := utils.AskQuestionWithBoolAnswer("Использовать Gemfile?")
	config.Set(configs.KeyIOSDependenciesGemfileUse, answer)

	answer = utils.AskQuestionWithBoolAnswer("Использовать Podfile?")
	config.Set(configs.KeyIOSDependenciesPodfileUse, answer)

	xcodeprojFilename := utils.AskQuestionWithChooseFileAnswer("Выберите .xcodeproj файл:", ".xcodeproj")
	if xcodeprojFilename == "" {
		utils.PrintlnAttentionMessage("Пропущена настройка iOS проекта")
		return
	}
	config.Set(configs.KeyIOSXcodeprojFilename, xcodeprojFilename)
	config.Set(configs.KeyIOSProjectName, strings.Replace(xcodeprojFilename, ".xcodeproj", "", 1))

	codeProjectFolderName := utils.AskQuestionWithChooseFolderAnswer("\nВыберите папку с кодом проекта (обычно она называется так же как и xcodeproj файл): ")
	if codeProjectFolderName == "" {
		codeProjectFolderName = "."
	}
	config.Set(configs.KeyIOSFolderNameCode, codeProjectFolderName)
	config.Set(configs.KeyIOSTargetnameCode, codeProjectFolderName)

	unitTestsFolderName := utils.AskQuestionWithChooseFolderAnswer("\nВыберите папку с UNIT тестами: ")
	if unitTestsFolderName != "" {
		config.Set(configs.KeyIOSFolderNameUnitTests, unitTestsFolderName)
		config.Set(configs.KeyIOSTargetnameUnitTests, unitTestsFolderName)
	}

	uiTestsFolderName := utils.AskQuestionWithChooseFolderAnswer("\nВыберите папку с UI тестами: ")
	if uiTestsFolderName != "" {
		config.Set(configs.KeyIOSFolderNameUITests, uiTestsFolderName)
		config.Set(configs.KeyIOSTargetnameUITests, uiTestsFolderName)
	}

	if config.GetBool(configs.KeyIOSDependenciesGemfileUse) && utils.IsFileExist("Gemfile") {
		utils.InstallGemDependencies()
	}
}
