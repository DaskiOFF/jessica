package templategenerator

import (
	"path/filepath"
	"strings"

	"github.com/daskioff/jessica/configs"
	"github.com/daskioff/jessica/utils/files"
	"github.com/daskioff/jessica/utils/git"
	"github.com/daskioff/jessica/utils/print"

	"github.com/spf13/viper"

	"github.com/daskioff/jessica/flows"
)

type MapKeys map[string]interface{}

const TemplateDescriptionFileName = "templates.yml"

type TemplateGeneratorFlow struct {
}

func (flow *TemplateGeneratorFlow) Start(args []string) {
	if len(args) == 0 {
		print.PrintlnAttentionMessage("Необходимо указать какое действие вы хотите выполнить. Чтобы увидеть список действий воспользуйтесь командой help")
		return
	}

	if args[0] == "pull" {
		args := args[1:]

		if len(args) == 0 {
			print.PrintlnErrorMessage("Вы не указали URL git репозитория")
			return
		} else {
			url := args[0]
			args := args[1:]

			if !strings.HasPrefix(url, "http") {
				url = "https://" + url
			}

			if !strings.HasSuffix(url, ".git") {
				url = url + ".git"
			}

			branch := ""
			if len(args) > 0 {
				branch = args[0]
			}

			path := configs.ProjectConfig.GetString(configs.KeyTemplatesFolderName)
			err := git.Clone(url, branch, path)
			if err != nil {
				panic(err)
			}
		}
	}

	err := Validate()
	if err != nil {
		print.PrintlnErrorMessage(err.Error())
		return
	}

	templates := searchTemplates()
	if args[0] == "list" {
		if len(templates) == 0 {
			print.PrintlnAttentionMessage("Шаблоны не найдены")
		} else {
			list := ""
			for _, template := range templates {
				if len(list) == 0 {
					list = template
				} else {
					list = list + "\n" + template
				}
			}
			print.PrintlnInfoMessage(list)
		}
	} else if args[0] == "gen" {
		if len(args) == 1 {
			print.PrintlnAttentionMessage("Не указано имя шаблона для генерации")
		} else {
			templatePath := filepath.Join(templatesRootPath(), args[1], TemplateDescriptionFileName)
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
				questions := []quest{}
				answers := make(map[string]interface{}, 0)
				if questionsInterface != nil {
					questions = newQuestions(questionsInterface.([]interface{}))
					answers = askQuestions(questions)
				}

				templateName := args[1]
				codeAddedFiles := generateTemplates(v, "code_files", templateName, args[2], customKeys, answers)

				testCodeAddedFiles := []AddedFile{}
				if needGenerateTests {
					testCodeAddedFiles = generateTemplates(v, "test_files", templateName, args[2], customKeys, answers)
				}

				mockCodeAddedFiles := []AddedFile{}
				if needGenerateMocks {
					mockCodeAddedFiles = generateTemplates(v, "mock_files", templateName, args[2], customKeys, answers)
				}

				if configs.ProjectConfig.GetString(configs.KeyProjectType) == "iOS" {
					xcodeproj([]XcodeProjAdded{
						XcodeProjAdded{
							configs.ProjectConfig.GetString(configs.KeyIOSXcodeprojFilename),
							[]XcodeProjTargetAddedFiles{
								XcodeProjTargetAddedFiles{
									configs.ProjectConfig.GetString(configs.KeyIOSTargetnameCode),
									codeAddedFiles,
								},
								XcodeProjTargetAddedFiles{
									configs.ProjectConfig.GetString(configs.KeyIOSTargetnameUnitTests),
									testCodeAddedFiles,
								},
								XcodeProjTargetAddedFiles{
									configs.ProjectConfig.GetString(configs.KeyIOSTargetnameUnitTests),
									mockCodeAddedFiles,
								}}}})
				}

				print.PrintlnSuccessMessage(templateName + " сгенерирован")
			}
		}
	}
}

func (flow *TemplateGeneratorFlow) Setup() {}

func (flow *TemplateGeneratorFlow) Description() string {
	return `
--------------------------------------------------------------------------------
	Генерация шаблонов
		- list – Список шаблонов
		- gen – Генерация шаблона
		- pull – Скачать шаблоны с репозитория
--------------------------------------------------------------------------------`
}

// ----------------------------------------------------------------------------
func NewFlow() flows.Flow {
	flow := TemplateGeneratorFlow{}
	return &flow
}
