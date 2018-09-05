package template

import (
	"bufio"
	"bytes"
	"io"
	textTemplate "text/template"
)

// ExecuteFile Шаблон из файла по пути path
func ExecuteFile(path string, writer io.Writer, params map[string]interface{}) error {
	structTemplate, err := textTemplate.ParseFiles(path)
	if err != nil {
		return err
	}

	err = structTemplate.Execute(writer, params)
	if err != nil {
		return err
	}

	return nil
}

// ExecuteString шаблон из строки
func ExecuteString(name string, stringTemplate string, params interface{}) string {
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
