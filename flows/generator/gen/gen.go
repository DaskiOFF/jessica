package gen

import (
	"errors"

	"github.com/daskioff/jessica/flows/generator/gen/gentemplate"

	"github.com/daskioff/jessica/configs/models"
	"github.com/daskioff/jessica/utils/files"
	"github.com/daskioff/jessica/utils/path"
	"github.com/daskioff/jessica/utils/print"
	"github.com/daskioff/jessica/utils/question"
	"github.com/daskioff/jessica/utils/xcodeproj"
)

// DescriptionFileName содержит имя файла описывающего шаблон
const DescriptionFileName = "templates.yml"

// Execute выполняет генерацию шаблона
func Execute(args []string,
	templatesRootPath string,
	globalConfig *models.ConfigGlobal,
	projectConfig *models.ConfigProject,
	iosConfig *models.ConfigIOS,
	otherConfig *models.ConfigOther) error {

	if len(args) == 0 {
		return errors.New("Не указано имя шаблона")
	}

	templatesFolderName := projectConfig.GetTemplatesFolderName()
	absTemplatesFolderPath, err := path.InProjectRoot(templatesFolderName)
	if err != nil {
		return err
	}

	generatorParams := NewParams(args)

	absTemplateConfigPath, err := path.InProjectRoot(templatesFolderName, generatorParams.TemplateName, DescriptionFileName)
	if err != nil {
		return err
	}

	if !files.IsFileExist(absTemplateConfigPath) {
		return errors.New("Шаблон с именем " + generatorParams.TemplateName + " не найден")
	}

	templateDescription, err := gentemplate.ParseDescription(absTemplateConfigPath)
	if err != nil {
		return err
	}

	answers := map[string]interface{}{}
	if len(templateDescription.Questions) > 0 {
		answers = askQuestions(templateDescription.Questions)
	}

	templatesParams := gentemplate.New(
		absTemplatesFolderPath,
		generatorParams.TemplateName,
		generatorParams.ModuleName,
		generatorParams.CustomKeys,
		answers,
		templateDescription.Variables)
	templatesParams.AppendFrom(globalConfig, projectConfig, iosConfig, otherConfig)

	codeAddedFiles := generateFiles(templateDescription.CodeFiles, templatesParams)

	testCodeAddedFiles := []xcodeproj.AddedFile{}
	if generatorParams.NeedGenerateTests {
		testCodeAddedFiles = generateFiles(templateDescription.TestFiles, templatesParams)
	}

	mockCodeAddedFiles := []xcodeproj.AddedFile{}
	if generatorParams.NeedGenerateMock {
		mockCodeAddedFiles = generateFiles(templateDescription.MockFiles, templatesParams)
	}

	if projectConfig.GetProjectType() == models.ConfigProjectTypeIOS {
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

	print.PrintlnSuccessMessage(generatorParams.TemplateName + " сгенерирован")
	return nil
}

func askQuestions(questions []gentemplate.Question) map[string]interface{} {
	answers := make(map[string]interface{}, 0)

	for _, quest := range questions {
		answer := question.AskQuestion(quest.Text, quest.IsRequired)
		answers[quest.Key] = answer
	}

	return answers
}
