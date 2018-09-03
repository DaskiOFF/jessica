package gentemplate

import (
	"strings"
	"time"

	"github.com/daskioff/jessica/configs/models"
)

type Params struct {
	AbsTemplateFolderPath string
	TemplateName          string
	ModuleName            string

	CustomKeys    map[string]interface{}
	Answers       map[string]interface{}
	DeveloperInfo map[string]interface{}
	CommonInfo    map[string]interface{}
	ModuleInfo    map[string]interface{}

	ProjectInfo map[string]interface{}
}

func New(
	absTemplateFolderPath string,
	templateName string,
	moduleName string,
	customKeys map[string]interface{},
	answers map[string]interface{}) *Params {

	tp := Params{
		AbsTemplateFolderPath: absTemplateFolderPath,
		TemplateName:          templateName,
		ModuleName:            moduleName,

		CustomKeys:    customKeys,
		Answers:       answers,
		DeveloperInfo: map[string]interface{}{},

		ProjectInfo: map[string]interface{}{},
	}

	tp.CommonInfo = commonInfo()
	tp.ModuleInfo = moduleInfo(moduleName)

	return &tp
}

// AppendFrom добавляет параметры из конфигов
func (tp *Params) AppendFrom(gc *models.ConfigGlobal, pc *models.ConfigProject, iosConfig *models.ConfigIOS, otherConfig *models.ConfigOther) {

	tp.DeveloperInfo = developerParams(gc, pc)

	switch pc.GetProjectType() {
	case models.ConfigProjectTypeIOS:
		tp.ProjectInfo = iosParams(iosConfig)
	case models.ConfigProjectTypeOther:
		tp.ProjectInfo = otherParams(otherConfig)
	}
}

func commonInfo() map[string]interface{} {
	currentTime := time.Now()

	return map[string]interface{}{
		"date": currentTime.Format("02.01.2006"),
		"year": currentTime.Year(),
	}
}

func moduleInfo(moduleName string) map[string]interface{} {
	nameFirstLower := ""
	if len(moduleName) > 1 {
		nameFirstLower = strings.ToLower(moduleName[:1]) + moduleName[1:]
	}

	return map[string]interface{}{
		"name":           moduleName,
		"nameCapitalize": strings.Title(moduleName),
		"nameFirstLower": nameFirstLower,
		"nameUppercase":  strings.ToUpper(moduleName),
		"nameLowercase":  strings.ToLower(moduleName),
	}
}

func developerParams(gc *models.ConfigGlobal, pc *models.ConfigProject) map[string]interface{} {
	return map[string]interface{}{
		"name":        gc.GetUsername(),
		"companyName": pc.GetCompanyName(),
	}
}

func iosParams(config *models.ConfigIOS) map[string]interface{} {
	params := map[string]interface{}{}

	params["projectName"] = config.GetProjectName()
	params["projectCodeFolderPath"] = config.GetFolderNameCode()

	if config.HasFolderNameUnitTests() {
		params["projectTestsName"] = config.GetFolderNameUnitTests()
	} else {
		params["projectTestsName"] = params["projectCodeFolderPath"]
	}

	if config.HasFolderNameUITests() {
		params["projectUITestsName"] = config.GetFolderNameUITests()
	} else {
		params["projectUITestsName"] = params["projectCodeFolderPath"]
	}

	return params
}

func otherParams(config *models.ConfigOther) map[string]interface{} {
	params := map[string]interface{}{}

	params["projectName"] = config.GetProjectName()
	params["projectCodeFolderPath"] = config.GetProjectFolderName()

	return params
}
