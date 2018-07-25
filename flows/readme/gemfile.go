package readme

import (
	"io/ioutil"
	"regexp"

	"github.com/daskioff/jessica/utils"
)

const gemFileName = "Gemfile"

// Read Читает Gemfile и выбирает из него список зависимостей
func readGemfile() ([]string, error) {
	var re = regexp.MustCompile(`(?m)^gem .*"$`)

	fileContent, err := ioutil.ReadFile(gemFileName)
	if err != nil {
		return nil, err
	}

	text := string(fileContent)
	matches := re.FindAllString(text, -1)
	return matches, nil
}

// Check Проверяет существование Gemfile, если его нет, то его создает и заполняет значением по умолчанию
func checkGemfile() {
	content := `source "https://rubygems.org"

gem "xcodeproj"
gem "fastlane", ">= 2.96.1", "<= 3.0.0"
gem "cocoapods", "~> 1.5"`

	fileName := gemFileName
	if !utils.IsFileExist(fileName) {
		utils.WriteToFile(fileName, content)
		utils.PrintlnSuccessMessage(fileName + " successfully created")
	}
}
