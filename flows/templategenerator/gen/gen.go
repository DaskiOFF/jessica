package gen

import (
	"path/filepath"
	"strings"

	"github.com/daskioff/jessica/configs/models"
	"github.com/daskioff/jessica/utils/files"
	"github.com/daskioff/jessica/utils/print"
	"github.com/daskioff/jessica/utils/xcodeproj"
	"github.com/spf13/viper"

	q "github.com/daskioff/jessica/flows/templategenerator/gen/questions"
)

const TemplateDescriptionFileName = "templates.yml"

func Execute(args []string,
	templatesRootPath string,
	globalConfig *models.ConfigGlobal,
	projectConfig *models.ConfigProject,
	iosConfig *models.ConfigIOS,
	otherConfig *models.ConfigOther) {

	if len(args) < 2 {
		print.PrintlnAttentionMessage("Не указано имя шаблона или название модуля для генерации")
		return
	}

	templateName := args[0]
	moduleName := args[1]
	templateConfigPath := filepath.Join(templatesRootPath, templateName, TemplateDescriptionFileName)
	if !files.IsFileExist(templateConfigPath) {
		print.PrintlnErrorMessage("Шаблон с именем " + args[0] + " не найден")
		return
	}

	v := viper.New()
	v.SetConfigFile(templateConfigPath)

	err := v.ReadInConfig()
	if err != nil {
		print.PrintlnErrorMessage(err.Error())
		return
	}

	needGenerateTests := true
	needGenerateMocks := true
	customKeys := map[string]interface{}{}
	if len(args) > 3 {
		for _, arg := range args[3:] {
			if arg == "--notest" {
				needGenerateTests = false
			}
			if arg == "--nomock" {
				needGenerateMocks = false
			}

			// custom keys
			splitResult := strings.Split(arg, ":")
			if len(splitResult) == 2 {
				customKeys[splitResult[0]] = splitResult[1]
			}
		}
	}

	questionsInterface := v.Get("questions")
	answers := map[string]interface{}{}
	if questionsInterface != nil {
		questions := q.NewQuestions(questionsInterface.([]interface{}))
		answers = q.AskQuestions(questions)
	}

	params := generateParams{
		customKeys:    customKeys,
		answers:       answers,
		globalConfig:  globalConfig,
		projectConfig: projectConfig,
		iosConfig:     iosConfig,
		otherConfig:   otherConfig,
	}
	codeAddedFiles := generateTemplates(v, "code_files", templateName, moduleName, params)

	testCodeAddedFiles := []xcodeproj.AddedFile{}
	if needGenerateTests {
		testCodeAddedFiles = generateTemplates(v, "test_files", templateName, moduleName, params)
	}

	mockCodeAddedFiles := []xcodeproj.AddedFile{}
	if needGenerateMocks {
		mockCodeAddedFiles = generateTemplates(v, "mock_files", templateName, moduleName, params)
	}

	if projectConfig.GetProjectType() == "iOS" {
		xcodeproj.AddFilesToTarget([]xcodeproj.XcodeProjAdded{
			xcodeproj.XcodeProjAdded{
				XcodeprojFilePath: iosConfig.GetXcodeprojFilename(),
				TargetFiles: []xcodeproj.XcodeProjTargetAddedFiles{
					xcodeproj.XcodeProjTargetAddedFiles{
						TargetName: iosConfig.GetTargetNameCode(),
						AddedFiles: codeAddedFiles,
					},
					xcodeproj.XcodeProjTargetAddedFiles{
						TargetName: iosConfig.GetTargetNameUnitTests(),
						AddedFiles: testCodeAddedFiles,
					},
					xcodeproj.XcodeProjTargetAddedFiles{
						TargetName: iosConfig.GetTargetNameUnitTests(),
						AddedFiles: mockCodeAddedFiles,
					}}}})
	}

	print.PrintlnSuccessMessage(templateName + " сгенерирован")
}
