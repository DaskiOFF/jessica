package projectStruct

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/daskioff/update_readme_ios/utils"
)

const FileName = ".project_struct.tpl.md"

func Check() {
	var reader *bufio.Reader
	if !utils.IsFileExist(FileName) {
		reader = bufio.NewReader(os.Stdin)
		fmt.Print("Use custom project struct (y/n): ")
		answer, _ := reader.ReadString('\n')

		if len(answer) > 0 && answer[0] == 'y' {
			createProjectStruct()
			return
		}

		utils.PrintlnAttentionMessage("Skipped the creation of the project structure")
	}
}

func createProjectStruct() {
	createTemplateFile()

	projectName := requestProjectName()
	if len(projectName) == 0 {
		utils.PrintlnAttentionMessage("Skipped the creation of the project structure. Project name is empty")
		return
	}

	os.MkdirAll(projectName+"/config", os.ModePerm)

	os.MkdirAll(projectName+"/extensions", os.ModePerm)

	os.MkdirAll(projectName+"/models", os.ModePerm)

	os.MkdirAll(projectName+"/services", os.ModePerm)
	os.MkdirAll(projectName+"/services/api", os.ModePerm)

	os.MkdirAll(projectName+"/usecases", os.ModePerm)

	os.MkdirAll(projectName+"/presentation/resources/r", os.ModePerm)
	os.MkdirAll(projectName+"/presentation/resources/localization", os.ModePerm)
	os.MkdirAll(projectName+"/presentation/flows", os.ModePerm)
	os.MkdirAll(projectName+"/presentation/components/views", os.ModePerm)
	os.MkdirAll(projectName+"/presentation/components/tableCells", os.ModePerm)
	os.MkdirAll(projectName+"/presentation/components/collectionCells", os.ModePerm)
	os.MkdirAll(projectName+"/presentation/controllers/Main", os.ModePerm)

	os.MkdirAll(projectName+"/support", os.ModePerm)

	utils.PrintlnSuccessMessage("Project structure created")
}

func requestProjectName() string {
	projectName := utils.ProjectName()

	reader := bufio.NewReader(os.Stdin)
	if len(projectName) > 0 {
		fmt.Print("Your project has name " + projectName + "? (y/n): ")
		answer, _ := reader.ReadString('\n')

		if len(answer) > 0 && answer[0] == 'y' {
			return projectName
		}
	}

	fmt.Print("Enter project name: ")
	answer, _ := reader.ReadString('\n')
	return strings.Replace(answer, "\n", "", -1)
}

func createTemplateFile() {
	content := `### Структура проекта
- %*%{{ .projectName }}%*% – папка проекта
	- %*%config%*%
		- AppConfig (protocol)
		- AppConfigDev
		- AppConfigRelease
	- %*%extensions%*%
		- Файл называем именем класса, для которого описываем extension
		- Если расширение скорей всего понадобится в других проектах, то его надо вынести в RKFoundationExtensions или RKUIExtensions
	- %*%models%*%
	- %*%services%*%
		- %*%api%*%
		- Каждый сервис закрывается протоколом и имеет реализацию по умолчанию, например, %*%ServiceFilters%*% (protocol), %*%ServiceFilterUserDefault%*%
		- Работа с какими-то сервисами и фреймворками iOS, например, %*%corelocation%*%
	- %*%usecases%*%
		- Разбивка по сущностям
	- %*%presentation%*%
		- %*%resources%*%
			- %*%r%*%
			- %*%localization%*%
			- Assets, LaunchScreen
		- %*%flows%*% – Flow приложения (координаторы)
		- %*%components%*%
			- %*%views%*%
			- %*%tableCells%*%
			- %*%collectionCells%*%
		- %*%controllers%*%
			- Каждый контроллер создаем в папке с его именем без суффикса %*%VC%*%, компоненты, которые разрабатываются только для этого экрана создаются внутри этой папки
			- Название контроллера содержит суффикс %*%VC%*%, например %*%MainVC%*%
			- %*%Main%*% (Пример)
				- %*%components%*%
					- HeaderView
				- MainVC
	- %*%support%*%
		- AppDelegate
		- Info.plist and other .plist
- %*%Project name Tests%*% – папка с тестами проекта
- %*%fastlane%*% – поддержка Fastlane
- %*%Pods%*% – зависимости проекта`

	content = utils.FixBackQuotes(content)
	fileName := FileName
	if !utils.IsFileExist(fileName) {
		utils.WriteToFile(fileName, content)
		utils.PrintlnSuccessMessage(fileName + " successfully created")
	}
}
