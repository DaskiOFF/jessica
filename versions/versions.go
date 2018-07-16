package versions

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/daskioff/jessica/utils"
)

func readVersionFile(fileName string) (string, error) {
	fileContent, err := ioutil.ReadFile(fileName)
	content := string(fileContent)
	content = strings.Replace(content, "\n", "", -1)
	return content, err
}

// ReadXcodeVersion Чтение версии xcode из файла .xcode-version
func ReadXcodeVersion() (string, error) {
	return readVersionFile(".xcode-version")
}

// ReadSwiftVersion Чтение версии swift из файла .swift-version
func ReadSwiftVersion() (string, error) {
	return readVersionFile(".swift-version")
}

// CheckVersionFiles Проверка существования файлов для версии xcode и swift, если их нет,
// то пользователя спросят какие версии он использует, и эти версии будут сохранены в соответствующие файлы
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
