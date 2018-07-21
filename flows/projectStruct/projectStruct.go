package projectstruct

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/daskioff/jessica/configs"

	"github.com/daskioff/jessica/flows"
	"github.com/daskioff/jessica/utils"
)

const FileName = ".project_struct.tpl.md"

type ProjectStructFlow struct {
}

func (flow *ProjectStructFlow) Start(args []string) {
	useCustomStruct := configs.ProjectConfig.GetBool(configs.KeyUseCustomProjectStruct)
	hasCustomStruct := configs.ProjectConfig.IsSet(configs.KeyCustomProjectStruct)

	if len(args) > 0 && args[0] == "gen" {
		if !useCustomStruct || !hasCustomStruct {
			utils.PrintlnAttentionMessage("Необходимо сначала сконфигурировать с помощью команды `struct`")
			return
		}

		generate()

		createTemplateFile()
		utils.PrintlnSuccessMessage("Отредактируйте файл " + FileName + ", чтобы описать вашу структуру")
		return
	}

	if !useCustomStruct {
		useCustomStruct = requestUseCustomStruct()
	}

	if !useCustomStruct {
		return
	}

	if hasCustomStruct {
		showCurrentProjectStruct()
	} else {
		showExample()
	}
}

func (flow *ProjectStructFlow) Description() string {
	return `
	--------------------------------------------------------------------------------
	Генерация структуры файлов описаных в конфиге
	--------------------------------------------------------------------------------
	`
}

// ----------------------------------------------------------------------------
func NewFlow() flows.Flow {
	flow := ProjectStructFlow{}
	return &flow
}

// requestUseCustomStruct Проверяет наличие файла шаблона описывающего структуру проекта, если его нет, то предлагает создать его и структуру из папок
func requestUseCustomStruct() (result bool) {
	answer := utils.AskQuestionWithAnswers("Use custom project struct? (y/n)", []string{"y", "n", "Y", "N"})
	if answer == "y" || answer == "Y" {
		configs.ProjectConfig.Set(configs.KeyUseCustomProjectStruct, true)
		result = true
	} else {
		configs.ProjectConfig.Set(configs.KeyUseCustomProjectStruct, false)
		utils.PrintlnAttentionMessage("Skipped the creation of the project structure")
		result = false
	}

	configs.WriteProject()
	return
}

func showExample() {
	const exampleStruct = configs.KeyCustomProjectStruct + `:
		- config
		- di:
			- factories
		- extensions
		- models
		- services:
			- api
		- usecases
		- presentation:
			- resources:
				- r
				- localization
				- fonts
			- flows
			- components:
				- views
				- tableCells
				- collectionCells
			- controllers
		- support`
	utils.PrintlnInfoMessage(`
	Для создания генерируемой структуры вам необходимо описать ее в локальном файле конфигурации .jessica.yml
	Описываемая файловая структура будет создаваться внутри папки проекта
	
	Например
	` + exampleStruct)
}

func showCurrentProjectStruct() {
	projectStructure := configs.ProjectConfig.Get(configs.KeyCustomProjectStruct)
	info := getPathsString(projectStructure, "  ", "  ")
	utils.PrintlnInfoMessage("Структура из файла конфигурации\n\n" + info)
}

func projectPaths() []string {
	projectStructure := configs.ProjectConfig.Get(configs.KeyCustomProjectStruct)
	return getPaths(projectStructure)
}

func getPaths(in interface{}) []string {
	switch v := in.(type) {

	case string:
		return []string{v}

	case map[interface{}]interface{}:
		response := make([]string, 0)
		for s, b := range v {
			prefix, _ := s.(string)
			for _, path := range getPaths(b) {
				response = append(response, filepath.Join(prefix, path))
			}
		}
		return response

	case []string:
		return v

	case []interface{}:
		response := make([]string, 0)
		for _, b := range v {
			for _, path := range getPaths(b) {
				response = append(response, path)
			}
		}
		return response

	default:
		return make([]string, 0)
	}
}

func getPathsString(in interface{}, space string, currentSpace string) string {
	switch v := in.(type) {

	case string:
		return currentSpace + "- " + v

	case map[interface{}]interface{}:
		response := ""
		for s, b := range v {
			prefix, _ := s.(string)
			response = response + currentSpace + "- " + prefix + "\n"
			response = response + getPathsString(b, currentSpace+space, space)
		}
		return strings.TrimSuffix(response, "\n")

	case []string:
		response := ""
		for _, v := range v {
			response = response + currentSpace + "- " + v + "\n"
		}
		return strings.TrimSuffix(response, "\n")

	case []interface{}:
		response := ""
		for _, b := range v {
			response = response + getPathsString(b, currentSpace, space) + "\n"
		}
		return response

	default:
		return ""
	}
}

func generate() {
	projectName := configs.ProjectConfig.GetString(configs.KeyProjectName)
	if len(projectName) == 0 {
		utils.PrintlnAttentionMessage("Skipped the creation of the project structure. Project name is empty")
		return
	}
	paths := projectPaths()
	for _, path := range paths {
		resultPath := filepath.Join(projectName, path)

		os.MkdirAll(resultPath, os.ModePerm)
		utils.PrintlnInfoMessage(resultPath)
	}

	utils.PrintlnSuccessMessage("Project structure created")
}

func createTemplateFile() {
	projectStructure := configs.ProjectConfig.Get(configs.KeyCustomProjectStruct)
	projectStructureString := getPathsString(projectStructure, "  ", "  ")

	content := `### Структура проекта
- %*%{{ .projectName }}%*% – папка проекта
` + projectStructureString

	content = utils.FixBackQuotes(content)
	fileName := FileName
	if !utils.IsFileExist(fileName) {
		utils.WriteToFile(fileName, content)
		utils.PrintlnSuccessMessage(fileName + " successfully created")
	}
}
