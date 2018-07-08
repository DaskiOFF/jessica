package versions

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/daskioff/update_readme_ios/utils"
)

func readVersionFile(fileName string) (string, error) {
	fileContent, err := ioutil.ReadFile(fileName)
	content := string(fileContent)
	content = strings.Replace(content, "\n", "", -1)
	return content, err
}

func ReadXcodeVersion() (string, error) {
	return readVersionFile(".xcode-version")
}

func ReadSwiftVersion() (string, error) {
	return readVersionFile(".swift-version")
}

func CheckVersionFiles() {
	fileName := ".xcode-version"
	var reader *bufio.Reader
	if !utils.IsFileExist(fileName) {
		reader = bufio.NewReader(os.Stdin)
		fmt.Print("Enter xcode version: ")
		xcodeVersion, _ := reader.ReadString('\n')

		utils.WriteToFile(fileName, xcodeVersion)
		utils.PrintlnSuccessMessage(fileName + " successfully created")
	}

	fileName = ".swift-version"
	if !utils.IsFileExist(fileName) {
		if reader == nil {
			reader = bufio.NewReader(os.Stdin)
		}
		fmt.Print("Enter Swift version: ")
		swiftVersion, _ := reader.ReadString('\n')

		utils.WriteToFile(fileName, swiftVersion)
		utils.PrintlnSuccessMessage(fileName + " successfully created")
	}
}
