package readme

import (
	"io/ioutil"
	"regexp"

	"github.com/daskioff/jessica/flows/internal"
)

// Read Читает Gemfile и выбирает из него список зависимостей
func (flow *ReadmeFlow) readGemfile() ([]string, error) {
	var re = regexp.MustCompile(`(?m)^gem .*"$`)

	filename := internal.GemfileFileName
	fileContent, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	text := string(fileContent)
	matches := re.FindAllString(text, -1)
	return matches, nil
}
