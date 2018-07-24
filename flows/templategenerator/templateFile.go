package templategenerator

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/daskioff/jessica/configs"
)

const (
	rewriteRequest = iota
	rewriteNo
	rewriteYes
)

type templateFile struct {
	name             string
	templatePath     string
	outputPathFolder string
	outputPathFile   string
	rewriteResult    int
}

func replaceTemplateVariableInPaths(inPath string, moduleName string) string {
	resultPath := inPath
	resultPath = strings.Replace(resultPath, "{{.projectName}}", configs.ProjectConfig.GetString(configs.KeyProjectName), -1)
	resultPath = strings.Replace(resultPath, "{{.projectTestsName}}", configs.ProjectConfig.GetString(configs.KeyProjectTestsFolderName), -1)
	resultPath = strings.Replace(resultPath, "{{.moduleName}}", moduleName, -1)

	return resultPath
}

func newTemplateFiles(in []interface{}, templateName string, moduleName string) []templateFile {
	templateFiles := []templateFile{}

	root, err := os.Getwd()
	if err != nil {
		fmt.Println(err.Error())
		return templateFiles
	}

	for _, codeInterface := range in {
		code := codeInterface.(map[interface{}]interface{})
		template := templateFile{}

		template.name = code["name"].(string)
		if strings.Contains(template.name, "{{.moduleName}}") {
			template.name = strings.Replace(template.name, "{{.moduleName}}", moduleName, -1)
		} else {
			template.name = moduleName + template.name
		}

		template.rewriteResult = rewriteRequest
		rewrite := code["rewrite"]
		if rewrite != nil {
			if rewrite.(bool) {
				template.rewriteResult = rewriteYes
			} else {
				template.rewriteResult = rewriteNo
			}
		}

		template.templatePath = code["template_path"].(string)
		template.templatePath = replaceTemplateVariableInPaths(template.templatePath, moduleName)
		template.templatePath = filepath.Join(templatesRootPath(), templateName, template.templatePath)

		template.outputPathFolder = code["output_path"].(string)
		template.outputPathFolder = replaceTemplateVariableInPaths(template.outputPathFolder, moduleName)
		template.outputPathFolder = filepath.Join(root, template.outputPathFolder)

		template.outputPathFile = filepath.Join(template.outputPathFolder, template.name)

		templateFiles = append(templateFiles, template)
	}

	return templateFiles
}
