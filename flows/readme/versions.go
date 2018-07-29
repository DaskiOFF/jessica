package readme

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/daskioff/jessica/utils/files"
	"github.com/daskioff/jessica/utils/print"
)

const xcodeVersionFileName = ".xcode-version"
const swiftVersionFileName = ".swift-version"

func readVersionFile(fileName string) (string, error) {
	fileContent, err := ioutil.ReadFile(fileName)
	if err != nil {
		return "", err
	}

	content := string(fileContent)
	content = strings.TrimSpace(content)
	return content, err
}

// ReadXcodeVersion Чтение версии xcode из файла
func readXcodeVersion() (string, error) {
	return readVersionFile(xcodeVersionFileName)
}

// ReadSwiftVersion Чтение версии swift из файла
func readSwiftVersion() (string, error) {
	return readVersionFile(swiftVersionFileName)
}

// CheckXcodeVersionFile Проверка существования файлa для версии xcode, если его нет,
// то пользователя у пользователя запросится версия, которую он использует, и эта версия будет сохранена в соответствующий файл
func checkXcodeVersionFile() {
	fileName := xcodeVersionFileName
	var reader *bufio.Reader

	if !files.IsFileExist(fileName) {
		reader = bufio.NewReader(os.Stdin)
		fmt.Print("Введите используемую версию Xcode: ")
		xcodeVersion, _ := reader.ReadString('\n')

		files.WriteToFile(fileName, xcodeVersion)
		print.PrintlnSuccessMessage(fileName + " создан")
	}
}

// CheckSwiftVersionFile Проверка существования файлa для версии swift, если его нет,
// то пользователя у пользователя запросится версия, которую он использует, и эта версия будет сохранена в соответствующий файл
func checkSwiftVersionFile() {
	fileName := swiftVersionFileName
	var reader *bufio.Reader

	if !files.IsFileExist(fileName) {
		if reader == nil {
			reader = bufio.NewReader(os.Stdin)
		}
		fmt.Print("Введите используемую версию Swift: ")
		swiftVersion, _ := reader.ReadString('\n')

		files.WriteToFile(fileName, swiftVersion)
		print.PrintlnSuccessMessage(fileName + " создан")
	}
}
