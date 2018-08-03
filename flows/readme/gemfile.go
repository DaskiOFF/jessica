package readme

import (
	"io/ioutil"
	"regexp"
)

const gemFileName = "Gemfile"

// Read Читает Gemfile и выбирает из него список зависимостей
func (flow *ReadmeFlow) readGemfile() ([]string, error) {
	var re = regexp.MustCompile(`(?m)^gem .*"$`)

	fileContent, err := ioutil.ReadFile(gemFileName)
	if err != nil {
		return nil, err
	}

	text := string(fileContent)
	matches := re.FindAllString(text, -1)
	return matches, nil
}
