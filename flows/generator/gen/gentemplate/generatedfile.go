package gentemplate

import (
	"errors"
)

const (
	GeneratedFileRewriteRequest int = iota
	GeneratedFileRewriteYes
	GeneratedFileRewriteNo
)

var (
	ErrGeneratedFileNeedName         = errors.New("Поле name у шаблона не может быть пустым")
	ErrGeneratedFileNeedTemplatePath = errors.New("Поле template_path у шаблона не может быть пустым")
	ErrGeneratedFileNeedOutputPath   = errors.New("Поле output_path у шаблона не может быть пустым")
)

type GeneratedFile struct {
	Name         string
	TemplatePath string
	OutputPath   string
	RewriteType  int
}

func parseGeneratedFiles(data []interface{}) ([]GeneratedFile, error) {
	if data == nil || len(data) == 0 {
		return []GeneratedFile{}, nil
	}

	files := []GeneratedFile{}
	for _, fileData := range data {
		g, err := parseGeneratedFile(fileData.(map[interface{}]interface{}))

		if err != nil {
			return nil, err
		}
		files = append(files, *g)
	}

	return files, nil
}

func parseGeneratedFile(data map[interface{}]interface{}) (*GeneratedFile, error) {
	g := GeneratedFile{
		Name:         "",
		TemplatePath: "",
		OutputPath:   "",
		RewriteType:  GeneratedFileRewriteRequest,
	}

	name := data["name"]
	if name == nil {
		return nil, ErrGeneratedFileNeedName
	}
	g.Name = name.(string)

	templatePath := data["template_path"]
	if templatePath == nil {
		return nil, ErrGeneratedFileNeedTemplatePath
	}
	g.TemplatePath = templatePath.(string)

	outputPath := data["output_path"]
	if outputPath == nil {
		return nil, ErrGeneratedFileNeedOutputPath
	}
	g.OutputPath = outputPath.(string)

	rewrite := data["rewrite"]
	if rewrite != nil {
		if rewrite.(bool) {
			g.RewriteType = GeneratedFileRewriteYes
		} else {
			g.RewriteType = GeneratedFileRewriteNo
		}
	}

	return &g, nil
}
