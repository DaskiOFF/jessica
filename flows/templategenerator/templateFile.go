package templategenerator

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	textTemplate "text/template"

	"github.com/daskioff/jessica/configs"
)

const (
	rewriteRequest = iota
	rewriteNo
	rewriteYes
)

type templateFile struct {
	name              string
	outputProjectPath string
	templatePath      string
	outputPathFolder  string
	outputPathFile    string
	rewriteResult     int
}

func newTemplateFiles(in []interface{}, templateName string, moduleName string, params MapKeys) []templateFile {
	templateFiles := []templateFile{}

	params["projectName"] = configs.ProjectConfig.GetString(configs.KeyIOSFolderNameCode)
	params["projectTestsName"] = configs.ProjectConfig.GetString(configs.KeyIOSFolderNameUnitTests)
	params["projectUITestsName"] = configs.ProjectConfig.GetString(configs.KeyIOSFolderNameUITests)

	root, err := os.Getwd()
	if err != nil {
		fmt.Println(err.Error())
		return templateFiles
	}

	for _, codeInterface := range in {
		code := codeInterface.(map[interface{}]interface{})
		templateFileResult := templateFile{}

		templateFileResult.name = code["name"].(string)
		if strings.Contains(templateFileResult.name, "{{") {
			templateFileResult.name = executeStringTemplate("generator template name", templateFileResult.name, params)
		} else {
			templateFileResult.name = moduleName + templateFileResult.name
		}

		templateFileResult.rewriteResult = rewriteRequest
		rewrite := code["rewrite"]
		if rewrite != nil {
			if rewrite.(bool) {
				templateFileResult.rewriteResult = rewriteYes
			} else {
				templateFileResult.rewriteResult = rewriteNo
			}
		}

		templateFileResult.templatePath = code["template_path"].(string)
		if strings.Contains(templateFileResult.templatePath, "{{") {
			templateFileResult.templatePath = executeStringTemplate("generator template templatePath", templateFileResult.templatePath, params)
		}
		templateFileResult.templatePath = filepath.Join(templatesRootPath(), templateName, templateFileResult.templatePath)

		templateFileResult.outputPathFolder = code["output_path"].(string)
		if strings.Contains(templateFileResult.outputPathFolder, "{{") {
			templateFileResult.outputPathFolder = executeStringTemplate("generator template outputPathFolder", templateFileResult.outputPathFolder, params)
		}
		templateFileResult.outputProjectPath = templateFileResult.outputPathFolder
		templateFileResult.outputPathFolder = filepath.Join(root, templateFileResult.outputPathFolder)

		templateFileResult.outputPathFile = filepath.Join(templateFileResult.outputPathFolder, templateFileResult.name)

		templateFiles = append(templateFiles, templateFileResult)
	}

	return templateFiles
}

func executeStringTemplate(name string, stringTemplate string, params interface{}) string {
	t, err := textTemplate.New(name).Parse(stringTemplate)
	if err != nil {
		panic(err)
	}

	var tpl bytes.Buffer
	writer := bufio.NewWriter(&tpl)
	err = t.Execute(writer, params)
	if err != nil {
		panic(err)
	}

	writer.Flush()
	return tpl.String()
}
