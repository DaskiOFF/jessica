package setup

import (
	"github.com/daskioff/jessica/configs"
	"github.com/daskioff/jessica/utils"
	"github.com/spf13/viper"
)

func iosSection(config *viper.Viper) {
	answer := utils.AskQuestionWithBoolAnswer("Use Gemfile?")
	config.Set(configs.KeyIOSDependenciesGemfileUse, answer)

	answer = utils.AskQuestionWithBoolAnswer("Use Podfile?")
	config.Set(configs.KeyIOSDependenciesPodfileUse, answer)

	xcodeprojFilename := utils.AskQuestionWithChooseFileAnswer("Choose .xcodeproj file:", ".xcodeproj")
	if xcodeprojFilename == "" {
		utils.PrintlnAttentionMessage("Skip iOS project setup")
		return
	}
	config.Set(configs.KeyIOSXcodeprojFilename, xcodeprojFilename)

	codeProjectFolderName := utils.AskQuestionWithChooseFolderAnswer("\nChoose project code folder: ")
	if codeProjectFolderName == "" {
		codeProjectFolderName = "."
	}
	config.Set(configs.KeyIOSFolderNameCode, codeProjectFolderName)
	config.Set(configs.KeyIOSTargetnameCode, codeProjectFolderName)

	unitTestsFolderName := utils.AskQuestionWithChooseFolderAnswer("\nChoose project UNIT tests folder: ")
	if unitTestsFolderName != "" {
		config.Set(configs.KeyIOSFolderNameUnitTests, unitTestsFolderName)
		config.Set(configs.KeyIOSTargetnameUnitTests, unitTestsFolderName)
	}

	uiTestsFolderName := utils.AskQuestionWithChooseFolderAnswer("\nChoose project UI tests folder: ")
	if uiTestsFolderName != "" {
		config.Set(configs.KeyIOSFolderNameUITests, uiTestsFolderName)
		config.Set(configs.KeyIOSTargetnameUITests, uiTestsFolderName)
	}
}
