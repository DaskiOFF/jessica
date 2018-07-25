package templategenerator

import (
	"path/filepath"

	"github.com/daskioff/jessica/configs"

	"github.com/spf13/viper"

	"github.com/daskioff/jessica/flows"
	"github.com/daskioff/jessica/utils"
)

const TemplateExt = ".tpl"
const TemplateDescriptionFileName = "templates.yml"

type TemplateGeneratorFlow struct {
}

func (flow *TemplateGeneratorFlow) Start(args []string) {
	if len(args) == 0 {
		utils.PrintlnAttentionMessage("Необходимо указать какое действие вы хотите выполнить. Чтобы увидеть список действий воспользуйтесь командой help")
		return
	}

	err := Validate()
	if err != nil {
		utils.PrintlnErrorMessage(err.Error())
		return
	}

	templates := searchTemplates()
	if args[0] == "list" {
		if len(templates) == 0 {
			utils.PrintlnAttentionMessage("Шаблоны не найдены")
		} else {
			list := ""
			for _, template := range templates {
				if len(list) == 0 {
					list = template
				} else {
					list = list + "\n" + template
				}
			}
			utils.PrintlnInfoMessage(list)
		}
	} else if args[0] == "gen" {
		if len(args) == 1 {
			utils.PrintlnAttentionMessage("Не указано имя шаблона для генерации")
		} else {
			templatePath := filepath.Join(templatesRootPath(), args[1], TemplateDescriptionFileName)
			if !utils.IsFileExist(templatePath) {
				utils.PrintlnErrorMessage("Шаблон с данным именем не найден")
			} else if len(args) < 3 {
				utils.PrintlnErrorMessage("Не указано имя генерируемого модуля")
			} else {
				v := viper.New()
				v.SetConfigFile(templatePath)

				err = v.ReadInConfig()
				if err != nil {
					utils.PrintlnErrorMessage(err.Error())
					return
				}

				templateName := args[1]
				codeAddedFiles := generateTemplates(v, "code_files", templateName, args[2])

				needGenerateTests := true
				needGenerateMocks := true
				if len(args) > 3 {
					for _, arg := range args[3:] {
						if arg == "--notest" {
							needGenerateTests = false
						}
						if arg == "--nomock" {
							needGenerateMocks = false
						}
					}
				}

				testCodeAddedFiles := []AddedFile{}
				if needGenerateTests {
					testCodeAddedFiles = generateTemplates(v, "test_files", templateName, args[2])
				}

				mockCodeAddedFiles := []AddedFile{}
				if needGenerateMocks {
					mockCodeAddedFiles = generateTemplates(v, "mock_files", templateName, args[2])
				}

				xcodeproj([]XcodeProjAdded{
					XcodeProjAdded{
						configs.ProjectConfig.GetString(configs.KeyProjectXcodeProjName),
						[]XcodeProjTargetAddedFiles{
							XcodeProjTargetAddedFiles{
								// TODO: Заменить на имя таргета для кода
								configs.ProjectConfig.GetString(configs.KeyProjectName),
								codeAddedFiles,
							},
							XcodeProjTargetAddedFiles{
								// TODO: Заменить на имя таргета для тестов
								configs.ProjectConfig.GetString(configs.KeyProjectTestsFolderName),
								testCodeAddedFiles,
							},
							XcodeProjTargetAddedFiles{
								// TODO: Заменить на имя таргета для моков
								configs.ProjectConfig.GetString(configs.KeyProjectTestsFolderName),
								mockCodeAddedFiles,
							}}}})
				utils.PrintlnSuccessMessage(templateName + " successfully generated")
			}
		}
	}
}

func (flow *TemplateGeneratorFlow) Description() string {
	return `
--------------------------------------------------------------------------------
	Генерация шаблонов
--------------------------------------------------------------------------------`
}

// ----------------------------------------------------------------------------
func NewFlow() flows.Flow {
	flow := TemplateGeneratorFlow{}
	return &flow
}
