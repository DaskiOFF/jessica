package templategenerator

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	textTemplate "text/template"
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

func (flow *TemplateGeneratorFlow) newTemplateFiles(in []interface{}, templateName string, moduleName string, params MapKeys) []templateFile {
	templateFiles := []templateFile{}

	params["projectName"] = flow.iosConfig.GetFolderNameCode()
	params["projectTestsName"] = flow.iosConfig.GetFolderNameUnitTests()
	params["projectUITestsName"] = flow.iosConfig.GetFolderNameUITests()

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
			templateFileResult.name = flow.executeStringTemplate("generator template name", templateFileResult.name, params)
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
			templateFileResult.templatePath = flow.executeStringTemplate("generator template templatePath", templateFileResult.templatePath, params)
		}
		templateFileResult.templatePath = filepath.Join(flow.templatesRootPath(), templateName, templateFileResult.templatePath)

		templateFileResult.outputPathFolder = code["output_path"].(string)
		if strings.Contains(templateFileResult.outputPathFolder, "{{") {
			templateFileResult.outputPathFolder = flow.executeStringTemplate("generator template outputPathFolder", templateFileResult.outputPathFolder, params)
		}
		templateFileResult.outputProjectPath = templateFileResult.outputPathFolder
		templateFileResult.outputPathFolder = filepath.Join(root, templateFileResult.outputPathFolder)

		templateFileResult.outputPathFile = filepath.Join(templateFileResult.outputPathFolder, templateFileResult.name)

		templateFiles = append(templateFiles, templateFileResult)
	}

	return templateFiles
}

func (flow *TemplateGeneratorFlow) executeStringTemplate(name string, stringTemplate string, params interface{}) string {
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
