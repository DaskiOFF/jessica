package gen

import (
	"path/filepath"

	"github.com/daskioff/jessica/configs/models"
	"github.com/daskioff/jessica/utils/files"
	"github.com/daskioff/jessica/utils/print"
	"github.com/daskioff/jessica/utils/xcodeproj"
	"github.com/spf13/viper"

	q "github.com/daskioff/jessica/flows/generator/gen/questions"
)

const TemplateDescriptionFileName = "templates.yml"

func Execute(args []string,
	templatesRootPath string,
	globalConfig *models.ConfigGlobal,
	projectConfig *models.ConfigProject,
	iosConfig *models.ConfigIOS,
	otherConfig *models.ConfigOther) {

	if len(args) < 1 {
		print.PrintlnAttentionMessage("Не указано имя шаблона")
		return
	}

	p := NewGenParams(args)

	templateConfigPath := filepath.Join(templatesRootPath, p.TemplateName, TemplateDescriptionFileName)
	if !files.IsFileExist(templateConfigPath) {
		print.PrintlnErrorMessage("Шаблон с именем " + p.TemplateName + " не найден")
		return
	}

	v := viper.New()
	v.SetConfigFile(templateConfigPath)

	err := v.ReadInConfig()
	if err != nil {
		print.PrintlnErrorMessage(err.Error())
		return
	}

	questionsInterface := v.Get("questions")
	answers := map[string]interface{}{}
	if questionsInterface != nil {
		questions := q.NewQuestions(questionsInterface.([]interface{}))
		answers = q.AskQuestions(questions)
	}

	params := generateParams{
		customKeys:    p.CustomKeys,
		answers:       answers,
		globalConfig:  globalConfig,
		projectConfig: projectConfig,
		iosConfig:     iosConfig,
		otherConfig:   otherConfig,
	}
	codeAddedFiles := generateTemplates(v, "code_files", p.TemplateName, p.ModuleName, params)

	testCodeAddedFiles := []xcodeproj.AddedFile{}
	if p.NeedGenerateTests {
		testCodeAddedFiles = generateTemplates(v, "test_files", p.TemplateName, p.ModuleName, params)
	}

	mockCodeAddedFiles := []xcodeproj.AddedFile{}
	if p.NeedGenerateMock {
		mockCodeAddedFiles = generateTemplates(v, "mock_files", p.TemplateName, p.ModuleName, params)
	}

	if projectConfig.GetProjectType() == "iOS" {
		unitTestsTargetName := iosConfig.GetTargetNameCode()
		if iosConfig.HasTargetNameUnitTests() {
			unitTestsTargetName = iosConfig.GetTargetNameUnitTests()
		}

		xcodeproj.AddFilesToTarget([]xcodeproj.XcodeProjAdded{
			xcodeproj.XcodeProjAdded{
				XcodeprojFilePath: iosConfig.GetXcodeprojFilename(),
				TargetFiles: []xcodeproj.XcodeProjTargetAddedFiles{
					xcodeproj.XcodeProjTargetAddedFiles{
						TargetName: iosConfig.GetTargetNameCode(),
						AddedFiles: codeAddedFiles,
					},
					xcodeproj.XcodeProjTargetAddedFiles{
						TargetName: unitTestsTargetName,
						AddedFiles: testCodeAddedFiles,
					},
					xcodeproj.XcodeProjTargetAddedFiles{
						TargetName: unitTestsTargetName,
						AddedFiles: mockCodeAddedFiles,
					}}}})
	}

	print.PrintlnSuccessMessage(p.TemplateName + " сгенерирован")
}
