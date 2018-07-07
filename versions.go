package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func readVersionFile(fileName string) (string, error) {
	fileContent, err := ioutil.ReadFile(fileName)
	content := string(fileContent)
	content = strings.Replace(content, "\n", "", -1)
	return content, err
}

func readXcodeVersion() (string, error) {
	return readVersionFile(".xcode-version")
}

func readSwiftVersion() (string, error) {
	return readVersionFile(".swift-version")
}

func checkVersionFiles() {
	fileName := ".xcode-version"
	var reader *bufio.Reader
	if !isFileExist(fileName) {
		reader = bufio.NewReader(os.Stdin)
		fmt.Print("Enter xcode version: ")
		xcodeVersion, _ := reader.ReadString('\n')

		writeToFile(fileName, xcodeVersion)
		printlnSuccessMessage(fileName + " successfully created")
	}

	fileName = ".swift-version"
	if !isFileExist(fileName) {
		if reader == nil {
			reader = bufio.NewReader(os.Stdin)
		}
		fmt.Print("Enter Swift version: ")
		swiftVersion, _ := reader.ReadString('\n')

		writeToFile(fileName, swiftVersion)
		printlnSuccessMessage(fileName + " successfully created")
	}
}
