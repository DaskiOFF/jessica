package gen

import (
	"path/filepath"
	"testing"

	"github.com/daskioff/jessica/flows/generator/gen/gentemplate"
)

func Test_GeneratedFilePaths(t *testing.T) {
	anyString := "anyString"
	anyPath1 := "any/path1"
	anyPath2 := "any/path2"

	absProjectRootPath := "root"
	absTemplateFolderPath := "any/path3"
	templateName := anyString

	generatedFile := gentemplate.GeneratedFile{
		Name:         anyString,
		OutputPath:   anyPath1,
		TemplatePath: anyPath2,
		RewriteType:  gentemplate.GeneratedFileRewriteRequest,
	}

	generatedFilePaths := newGeneratedFilePaths(&generatedFile, map[string]interface{}{}, templateName, absTemplateFolderPath, absProjectRootPath)

	if generatedFilePaths.FileName != generatedFile.Name {
		t.Errorf("Expected FileName == %v, got %v", generatedFile.Name, generatedFilePaths.FileName)
	}

	if generatedFilePaths.OutputProjectPath != generatedFile.OutputPath {
		t.Errorf("Expected OutputProjectPath == %v, got %v", generatedFile.OutputPath, generatedFilePaths.OutputProjectPath)
	}

	fp := filepath.Join(absTemplateFolderPath, templateName, generatedFile.TemplatePath)
	if generatedFilePaths.AbsTemplatePath != fp {
		t.Errorf("Expected AbsTemplatePath == %v, got %v", fp, generatedFilePaths.AbsTemplatePath)
	}

	fp = filepath.Join(absProjectRootPath, generatedFile.OutputPath)
	if generatedFilePaths.AbsOutputPath != fp {
		t.Errorf("Expected AbsOutputPath == %v, got %v", fp, generatedFilePaths.AbsOutputPath)
	}

	fp = filepath.Join(absProjectRootPath, generatedFile.OutputPath, generatedFile.Name)
	if generatedFilePaths.AbsFilePath != fp {
		t.Errorf("Expected AbsFilePath == %v, got %v", fp, generatedFilePaths.AbsFilePath)
	}
}
