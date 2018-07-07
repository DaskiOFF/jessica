package main

import (
	"io/ioutil"
	"regexp"
)

func readGemfile() ([]string, error) {
	var re = regexp.MustCompile(`(?m)^gem .*"$`)

	fileContent, err := ioutil.ReadFile("Gemfile")
	if err != nil {
		return nil, err
	}

	text := string(fileContent)
	matches := re.FindAllString(text, -1)
	return matches, nil
}

func checkGemfile() {
	content := `source "https://rubygems.org"

gem "fastlane", ">= 2.96.1", "<= 3.0.0"
gem "cocoapods", "~> 1.5"`

	fileName := "Gemfile"
	if !isFileExist(fileName) {
		writeToFile(fileName, content)
		printlnSuccessMessage(fileName + " successfully created")
	}
}
