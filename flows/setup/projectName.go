package setup

import (
	"strings"

	"github.com/daskioff/jessica/utils"
)

func projectName() string {
	projectName := utils.ProjectName()
	answer := "n"
	if projectName != "" {
		answer = utils.AskQuestionWithAnswers("Your project has name '"+projectName+"'? (y/n): ", []string{"y", "n", "Y", "N"})
	}

	if strings.ToLower(answer) == "n" {
		for {
			answer := utils.AskQuestion("Enter project name: ", true)
			projectName = answer
			if !strings.HasSuffix(answer, ".xcodeproj") {
				projectName = answer + ".xcodeproj"
			}
			if utils.IsFileExist(projectName) {
				break
			}
			utils.PrintlnInfoMessage("Файл '" + projectName + "' не найден")
		}
	}

	return projectName
}

func projectTestsFolderName(projectName string) string {
	testFolderName := projectName + "Tests"
	answer := "n"
	if projectName != "" && utils.IsFileExist(testFolderName) {
		answer = utils.AskQuestionWithAnswers("Your project has tests folder name '"+testFolderName+"'? (y/n): ", []string{"y", "n", "Y", "N"})
	}

	if strings.ToLower(answer) == "n" {
		for {
			answer := utils.AskQuestion("Enter project tests folder name: ", true)
			testFolderName = answer

			if testFolderName == projectName {
				utils.PrintlnAttentionMessage("Папка с тестами не может быть папкой проекта")
				continue
			}

			if utils.IsFileExist(testFolderName) {
				break
			}
			utils.PrintlnInfoMessage("Папка '" + testFolderName + "' не найдена")
		}
	}

	return testFolderName
}
