package readme

import (
	"io/ioutil"
	"regexp"

	"github.com/daskioff/jessica/configs"
	"github.com/daskioff/jessica/utils/bundle"
	"github.com/daskioff/jessica/utils/files"
	"github.com/daskioff/jessica/utils/print"
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
	if !files.IsFileExist(fileName) {
		files.WriteToFile(fileName, content)
		print.PrintlnSuccessMessage(fileName + " создан")
	}

	if configs.ProjectConfig.GetBool(configs.KeyIOSDependenciesGemfileUse) && files.IsFileExist("Gemfile") {
		bundle.Install()
	}
}
