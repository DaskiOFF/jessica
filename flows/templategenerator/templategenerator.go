package templategenerator

import (
	"path/filepath"
	"strings"

	"github.com/daskioff/jessica/configs/models"
	"github.com/daskioff/jessica/utils/files"
	"github.com/daskioff/jessica/utils/print"
	"github.com/daskioff/jessica/utils/xcodeproj"

	"github.com/spf13/viper"

	"github.com/daskioff/jessica/flows"
	"github.com/daskioff/jessica/flows/templategenerator/list"
	"github.com/daskioff/jessica/flows/templategenerator/pull"
	q "github.com/daskioff/jessica/flows/templategenerator/questions"
)

type MapKeys map[string]interface{}

const TemplateDescriptionFileName = "templates.yml"

type TemplateGeneratorFlow struct {
	globalConfig  *models.ConfigGlobal
	projectConfig *models.ConfigProject
	iosConfig     *models.ConfigIOS
}

func (flow *TemplateGeneratorFlow) Start(args []string) {
	if len(args) == 0 {
		print.PrintlnAttentionMessage("Необходимо указать какое действие вы хотите выполнить. Чтобы увидеть список действий воспользуйтесь командой help")
		return
	}

	if args[0] == "pull" {
		pull.Execute(args[1:], flow.projectConfig.GetTemplatesFolderName())
		return
	}

	err := flow.Validate()
	if err != nil {
		print.PrintlnErrorMessage(err.Error())
		return
	}

	templates := flow.searchTemplates()
	switch args[0] {
	case "list":
		list.Show(templates)
	case "gen":
		if len(args) == 1 {
			print.PrintlnAttentionMessage("Не указано имя шаблона для генерации")
		} else {
			templatePath := filepath.Join(flow.templatesRootPath(), args[1], TemplateDescriptionFileName)
			if !files.IsFileExist(templatePath) {
				print.PrintlnErrorMessage("Шаблон с данным именем не найден")
			} else if len(args) < 3 {
				print.PrintlnErrorMessage("Не указано имя генерируемого модуля")
			} else {
				v := viper.New()
				v.SetConfigFile(templatePath)

				err = v.ReadInConfig()
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
				questions := []q.Quest{}
				answers := map[string]interface{}{}
				if questionsInterface != nil {
					questions = q.NewQuestions(questionsInterface.([]interface{}))
					answers = q.AskQuestions(questions)
				}

				templateName := args[1]
				codeAddedFiles := flow.generateTemplates(v, "code_files", templateName, args[2], customKeys, answers)

				testCodeAddedFiles := []xcodeproj.AddedFile{}
				if needGenerateTests {
					testCodeAddedFiles = flow.generateTemplates(v, "test_files", templateName, args[2], customKeys, answers)
				}

				mockCodeAddedFiles := []xcodeproj.AddedFile{}
				if needGenerateMocks {
					mockCodeAddedFiles = flow.generateTemplates(v, "mock_files", templateName, args[2], customKeys, answers)
				}

				if flow.projectConfig.GetProjectType() == "iOS" {
					xcodeproj.AddFilesToTarget([]xcodeproj.XcodeProjAdded{
						xcodeproj.XcodeProjAdded{
							flow.iosConfig.GetXcodeprojFilename(),
							[]xcodeproj.XcodeProjTargetAddedFiles{
								xcodeproj.XcodeProjTargetAddedFiles{
									flow.iosConfig.GetTargetNameCode(),
									codeAddedFiles,
								},
								xcodeproj.XcodeProjTargetAddedFiles{
									flow.iosConfig.GetTargetNameUnitTests(),
									testCodeAddedFiles,
								},
								xcodeproj.XcodeProjTargetAddedFiles{
									flow.iosConfig.GetTargetNameUnitTests(),
									mockCodeAddedFiles,
								}}}})
				}

				print.PrintlnSuccessMessage(templateName + " сгенерирован")
			}
		}
	}
}

func (flow *TemplateGeneratorFlow) Description() string {
	return `
--------------------------------------------------------------------------------
	Генерация шаблонов
		- pull – Скачать шаблоны с репозитория
		- list – Вывести список шаблонов
		- gen – Генерация шаблона
--------------------------------------------------------------------------------`
}

// ----------------------------------------------------------------------------
func NewFlow(globalConfig *models.ConfigGlobal, projectConfig *models.ConfigProject, iosConfig *models.ConfigIOS) flows.Flow {
	flow := TemplateGeneratorFlow{}
	flow.globalConfig = globalConfig
	flow.projectConfig = projectConfig
	flow.iosConfig = iosConfig

	return &flow
}
