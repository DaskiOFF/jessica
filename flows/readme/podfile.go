package readme

import (
	"io/ioutil"
	"regexp"
)

const podFileName = "Podfile"

// Read Читает Podfile и выбирает из него список зависимостей для каждого таргета
func (flow *ReadmeFlow) readPodfile() ([]string, error) {
	var re = regexp.MustCompile(`(?ms)target (.*?)end`)

	fileContent, err := ioutil.ReadFile(podFileName)
	if err != nil {
		return nil, err
	}

	text := string(fileContent)
	matches := re.FindAllString(text, -1)
	return matches, nil
}
