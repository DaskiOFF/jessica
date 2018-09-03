package gen

import (
	"bufio"
	"os"
	"path/filepath"

	"github.com/daskioff/jessica/flows/generator/gen/gentemplate"

	"github.com/daskioff/jessica/utils/files"
	"github.com/daskioff/jessica/utils/path"
	"github.com/daskioff/jessica/utils/print"
	"github.com/daskioff/jessica/utils/question"
	"github.com/daskioff/jessica/utils/template"
	"github.com/daskioff/jessica/utils/xcodeproj"
)

func generateFiles(generatedFiles []gentemplate.GeneratedFile, templatesParams *gentemplate.Params) []xcodeproj.AddedFile {
	result := []xcodeproj.AddedFile{}

	if len(generatedFiles) == 0 {
		return result
	}

	for _, f := range generatedFiles {
		if addedFile := generateFile(&f, templatesParams); addedFile != nil {
			result = append(result, *addedFile)
		}
	}

	return result
}

func generateFile(generatedFile *gentemplate.GeneratedFile, templatesParams *gentemplate.Params) *xcodeproj.AddedFile {

	params := templatesParams.Map()

	projectRoot, err := path.ProjectRoot()
	if err != nil {
		panic(err)
	}

	generatedFilePaths := newGeneratedFilePaths(generatedFile, params, templatesParams.TemplateName, templatesParams.AbsTemplateFolderPath, projectRoot)

	err = os.MkdirAll(generatedFilePaths.AbsOutputPath, os.ModePerm)
	if err != nil {
		panic(err)
	}

	if files.IsFileExist(generatedFilePaths.AbsFilePath) {
		switch generatedFile.RewriteType {
		case gentemplate.GeneratedFileRewriteNo:
			return nil

		case gentemplate.GeneratedFileRewriteRequest:
			print.PrintlnAttentionMessage("Файл уже существует: " + generatedFilePaths.AbsFilePath)
			answer := question.AskQuestionWithBoolAnswer("Перезаписать файл?")
			if !answer {
				return nil
			}
		}
	}

	os.Remove(generatedFilePaths.AbsFilePath)
	file, err := os.OpenFile(generatedFilePaths.AbsFilePath, os.O_CREATE|os.O_WRONLY, os.ModePerm)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	params["fileName"] = generatedFilePaths.FileName

	err = template.ExecuteFile(generatedFilePaths.AbsTemplatePath, writer, params)
	if err != nil {
		panic(err)
	}

	err = writer.Flush()
	if err != nil {
		panic(err)
	}

	return &xcodeproj.AddedFile{
		Path:     generatedFilePaths.OutputProjectPath,
		Filename: generatedFilePaths.FileName,
	}
}

// GeneratedFilePaths содержит абсолютные и относительные пути необходимые для генерации шаблона
type GeneratedFilePaths struct {
	FileName          string
	OutputProjectPath string
	AbsTemplatePath   string
	AbsOutputPath     string
	AbsFilePath       string
}

func newGeneratedFilePaths(generatedFile *gentemplate.GeneratedFile, params map[string]interface{}, templateName string, AbsTemplateFolderPath string, projectRoot string) *GeneratedFilePaths {
	fileName := template.ExecuteString("generatedFile template name", generatedFile.Name, params)

	templatePath := template.ExecuteString("generatedFile template templatePath", generatedFile.TemplatePath, params)
	AbsTemplatePath := filepath.Join(AbsTemplateFolderPath, templateName, templatePath)

	outputProjectPath := template.ExecuteString("generator template outputPathFolder", generatedFile.OutputPath, params)
	AbsOutputPath := filepath.Join(projectRoot, outputProjectPath)
	AbsFilePath := filepath.Join(AbsOutputPath, fileName)

	return &GeneratedFilePaths{
		FileName:          fileName,
		OutputProjectPath: outputProjectPath,
		AbsTemplatePath:   AbsTemplatePath,
		AbsOutputPath:     AbsOutputPath,
		AbsFilePath:       AbsFilePath,
	}
}
