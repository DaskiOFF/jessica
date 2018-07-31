package readme

import (
	"io/ioutil"
	"regexp"

	"github.com/daskioff/jessica/utils/bundle"
	"github.com/daskioff/jessica/utils/files"
	"github.com/daskioff/jessica/utils/print"
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

// Check Проверяет существование Gemfile, если его нет, то его создает и заполняет значением по умолчанию
func (flow *ReadmeFlow) checkGemfile() {
	content := `source "https://rubygems.org"

gem "xcodeproj"
gem "fastlane", ">= 2.96.1", "<= 3.0.0"
gem "cocoapods", "~> 1.5"`

	fileName := gemFileName
	if !files.IsFileExist(fileName) {
		files.WriteToFile(fileName, content)
		print.PrintlnSuccessMessage(fileName + " создан")
	}

	if flow.iosConfig.GetGemfileUse() && files.IsFileExist("Gemfile") {
		bundle.Install()
	}
}
