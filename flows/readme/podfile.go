package readme

import (
	"io/ioutil"
	"regexp"

	"github.com/daskioff/jessica/flows/internal"
)

// Read Читает Podfile и выбирает из него список зависимостей для каждого таргета
func (flow *ReadmeFlow) readPodfile() ([]string, error) {
	var re = regexp.MustCompile(`(?ms)target (.*?)end`)

	filename := internal.PodfileFileName
	fileContent, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	text := string(fileContent)
	matches := re.FindAllString(text, -1)
	return matches, nil
}
