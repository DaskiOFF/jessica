package gemfile

import (
	"errors"
	"io/ioutil"
	"regexp"

	"github.com/daskioff/jessica/utils/files"
)

// DefaultFilename Default file name for Gemfile
const DefaultFilename = "Gemfile"

// ErrGemfileExist means that a Gemfile already exists
var ErrGemfileExist = errors.New(DefaultFilename + " already exists")

// CreateDefaultIOS create Gemfile with default content for iOS Project
func CreateDefaultIOS() error {
	content := `source "https://rubygems.org"

gem "fastlane", ">= 2.96.1", "<= 3.0.0"
gem "cocoapods", "~> 1.5"`

	fileName := DefaultFilename
	if files.IsFileExist(fileName) {
		return ErrGemfileExist
	}
	files.WriteToFile(fileName, content)

	return nil
}

// Dependencies returns a list of dependencies described in the Gemfile
func Dependencies() ([]string, error) {
	var re = regexp.MustCompile(`(?m)^gem .*"$`)

	filename := DefaultFilename
	fileContent, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	text := string(fileContent)
	matches := re.FindAllString(text, -1)
	return matches, nil
}
