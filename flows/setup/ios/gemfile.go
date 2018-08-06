package ios

import (
	"github.com/daskioff/jessica/flows/internal"
	"github.com/daskioff/jessica/utils/bundle"
	"github.com/daskioff/jessica/utils/files"
	"github.com/daskioff/jessica/utils/print"
)

// Check Проверяет существование Gemfile, если его нет, то его создает и заполняет значением по умолчанию
func checkGemfile() {
	content := `source "https://rubygems.org"

gem "fastlane", ">= 2.96.1", "<= 3.0.0"
gem "cocoapods", "~> 1.5"`

	fileName := internal.GemfileFileName
	if !files.IsFileExist(fileName) {
		files.WriteToFile(fileName, content)
		print.PrintlnSuccessMessage(fileName + " создан")

		bundle.Install()
	} else if files.IsFileExist(fileName + ".lock") {
		bundle.Update()
	}
}
