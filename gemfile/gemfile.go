package gemfile

import (
	"io/ioutil"
	"regexp"

	"github.com/daskioff/update_readme_ios/utils"
)

// Read Читает Gemfile и выбирает из него список зависимостей
func Read() ([]string, error) {
	var re = regexp.MustCompile(`(?m)^gem .*"$`)

	fileContent, err := ioutil.ReadFile("Gemfile")
	if err != nil {
		return nil, err
	}

	text := string(fileContent)
	matches := re.FindAllString(text, -1)
	return matches, nil
}

// Check Проверяет существование Gemfile, если его нет, то его создает и заполняет значением по умолчанию
func Check() {
	content := `source "https://rubygems.org"

gem "fastlane", ">= 2.96.1", "<= 3.0.0"
gem "cocoapods", "~> 1.5"`

	fileName := "Gemfile"
	if !utils.IsFileExist(fileName) {
		utils.WriteToFile(fileName, content)
		utils.PrintlnSuccessMessage(fileName + " successfully created")
	}
}
