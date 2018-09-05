package gentemplate

import (
	"github.com/spf13/viper"
)

type Description struct {
	Variables map[string]interface{}
	Questions []Question
	CodeFiles []GeneratedFile
	TestFiles []GeneratedFile
	MockFiles []GeneratedFile
}

func ParseDescription(filePath string) (*Description, error) {
	v := viper.New()
	v.SetConfigFile(filePath)

	err := v.ReadInConfig()
	if err != nil {
		return nil, err
	}

	td := Description{
		Questions: []Question{},
	}

	variables := v.Get("variables")
	if variables != nil {
		td.Variables, err = parseVariables(variables)
	}

	questions := v.Get("questions")
	if questions != nil {
		td.Questions, err = parseQuestions(questions.([]interface{}))
	}

	generatedFiles := v.Get("code_files")
	if generatedFiles != nil {
		td.CodeFiles, err = parseGeneratedFiles(generatedFiles.([]interface{}))
	}

	generatedFiles = v.Get("test_files")
	if generatedFiles != nil {
		td.TestFiles, err = parseGeneratedFiles(generatedFiles.([]interface{}))
	}

	generatedFiles = v.Get("mock_files")
	if generatedFiles != nil {
		td.MockFiles, err = parseGeneratedFiles(generatedFiles.([]interface{}))
	}

	return &td, nil
}
