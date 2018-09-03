package gentemplate

import (
	"testing"
)

func Test_ParsegeneratedFile(t *testing.T) {
	anyString := "anyString"

	data := []struct {
		data          map[interface{}]interface{}
		generatedFile *GeneratedFile
		err           error
	}{
		{
			data: map[interface{}]interface{}{
				"name":          anyString,
				"template_path": anyString,
				"output_path":   anyString,
				"rewrite":       false,
			},
			generatedFile: &GeneratedFile{
				Name:         anyString,
				TemplatePath: anyString,
				OutputPath:   anyString,
				RewriteType:  GeneratedFileRewriteNo,
			},
			err: nil,
		},
		{
			data: map[interface{}]interface{}{
				"name":          anyString,
				"template_path": anyString,
				"output_path":   anyString,
				"rewrite":       true,
			},
			generatedFile: &GeneratedFile{
				Name:         anyString,
				TemplatePath: anyString,
				OutputPath:   anyString,
				RewriteType:  GeneratedFileRewriteYes,
			},
			err: nil,
		},
		{
			data: map[interface{}]interface{}{
				"name":          anyString,
				"template_path": anyString,
				"output_path":   anyString,
			},
			generatedFile: &GeneratedFile{
				Name:         anyString,
				TemplatePath: anyString,
				OutputPath:   anyString,
				RewriteType:  GeneratedFileRewriteRequest,
			},
			err: nil,
		},
		{
			data: map[interface{}]interface{}{
				"template_path": anyString,
				"output_path":   anyString,
				"rewrite":       false,
			},
			generatedFile: nil,
			err:           ErrGeneratedFileNeedName,
		},
		{
			data: map[interface{}]interface{}{
				"name":        anyString,
				"output_path": anyString,
				"rewrite":     false,
			},
			generatedFile: nil,
			err:           ErrGeneratedFileNeedTemplatePath,
		},
		{
			data: map[interface{}]interface{}{
				"name":          anyString,
				"template_path": anyString,
				"rewrite":       false,
			},
			generatedFile: nil,
			err:           ErrGeneratedFileNeedOutputPath,
		},
	}

	for _, d := range data {
		result, err := parseGeneratedFile(d.data)

		if d.err == nil && err != nil || d.err != nil && err == nil {
			t.Error("Expected err == d.err, got ", err.Error())
		}

		if d.generatedFile == nil {
			if result != nil {
				t.Errorf("Expected %v == nil", result)
			}
			return
		}
		if result.Name != d.generatedFile.Name || result.TemplatePath != d.generatedFile.TemplatePath || result.OutputPath != d.generatedFile.OutputPath || result.RewriteType != d.generatedFile.RewriteType {
			t.Errorf("Expected data %v == %v, got %v", d.data, d.generatedFile, result)
		}
	}
}

func Test_ParsegeneratedFiles(t *testing.T) {
	anyString := "anyString"

	data := []struct {
		data          map[interface{}]interface{}
		generatedFile *GeneratedFile
		err           error
	}{
		{
			data: map[interface{}]interface{}{
				"name":          anyString,
				"template_path": anyString,
				"output_path":   anyString,
				"rewrite":       false,
			},
			generatedFile: &GeneratedFile{
				Name:         anyString,
				TemplatePath: anyString,
				OutputPath:   anyString,
				RewriteType:  GeneratedFileRewriteNo,
			},
			err: nil,
		},
		{
			data: map[interface{}]interface{}{
				"name":          anyString,
				"template_path": anyString,
				"output_path":   anyString,
				"rewrite":       true,
			},
			generatedFile: &GeneratedFile{
				Name:         anyString,
				TemplatePath: anyString,
				OutputPath:   anyString,
				RewriteType:  GeneratedFileRewriteYes,
			},
			err: nil,
		},
		{
			data: map[interface{}]interface{}{
				"name":          anyString,
				"template_path": anyString,
				"output_path":   anyString,
			},
			generatedFile: &GeneratedFile{
				Name:         anyString,
				TemplatePath: anyString,
				OutputPath:   anyString,
				RewriteType:  GeneratedFileRewriteRequest,
			},
			err: nil,
		},
		{
			data: map[interface{}]interface{}{
				"template_path": anyString,
				"output_path":   anyString,
				"rewrite":       false,
			},
			generatedFile: nil,
			err:           ErrGeneratedFileNeedName,
		},
		{
			data: map[interface{}]interface{}{
				"name":        anyString,
				"output_path": anyString,
				"rewrite":     false,
			},
			generatedFile: nil,
			err:           ErrGeneratedFileNeedTemplatePath,
		},
		{
			data: map[interface{}]interface{}{
				"name":          anyString,
				"template_path": anyString,
				"rewrite":       false,
			},
			generatedFile: nil,
			err:           ErrGeneratedFileNeedOutputPath,
		},
	}

	result, err := parseGeneratedFiles([]interface{}{data[0].data, data[1].data})
	if len(result) != 2 || err != nil {
		t.Error("Expected length result == 2, error == nil, got, ", len(result), err)
	}

	result, err = parseGeneratedFiles([]interface{}{data[0].data, data[3].data})
	if len(result) != 0 || err == nil || (err != nil && err.Error() != data[3].err.Error()) {
		t.Error("Expected result == nil, error == error key, got, ", result, err)
	}

	result, err = parseGeneratedFiles([]interface{}{data[0].data, data[4].data})
	if len(result) != 0 || err == nil || (err != nil && err.Error() != data[4].err.Error()) {
		t.Error("Expected result == nil, error == error text, got, ", result, err)
	}

	result, err = parseGeneratedFiles([]interface{}{data[0].data, data[5].data})
	if len(result) != 0 || err == nil || (err != nil && err.Error() != data[5].err.Error()) {
		t.Error("Expected result == nil, error == error text, got, ", result, err)
	}
}
