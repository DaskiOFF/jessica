package utils

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

// IsFileExist Проверяет существует ли файл с указанным именем
func IsFileExist(filename string) bool {
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		return false
	}

	return true
}

// WriteToFile Записывает переданный текст в указанный файл
func WriteToFile(fileName string, text string) {
	d1 := []byte(text)
	err := ioutil.WriteFile(fileName, d1, os.ModePerm)
	if err != nil {
		fmt.Println("Error create file with name: fileName")
	}
}

// PrintlnSuccessMessage Выводит сообщение успеха в лог
func PrintlnSuccessMessage(message string) {
	fmt.Println(message + "  🎉")
}

// PrintlnErrorMessage Выводит сообщение ошибки в лог
func PrintlnErrorMessage(message string) {
	fmt.Println("❌  " + message + "  ❌")
}

// PrintlnAttentionMessage Выводит сообщение заслуживающее внимания в лог
func PrintlnAttentionMessage(message string) {
	fmt.Println("🔶  " + message + "  🔶")
}

// FixBackQuotes Replace \%\*\*\*\% to \`\`\` and \%\*\% to \`
func FixBackQuotes(content string) string {
	content = strings.Replace(content, "%***%", "```", -1)
	content = strings.Replace(content, "%*%", "`", -1)

	return content
}

// ProjectName Возвращает имя проекта (определяется по наличию в папке файла с расширением .xcodeproj)
func ProjectName() string {
	files, err := ioutil.ReadDir(".")
	if err != nil {
		return ""
	}

	for _, file := range files {
		const suffix = ".xcodeproj"
		if strings.HasSuffix(file.Name(), suffix) {
			return strings.Replace(file.Name(), suffix, "", 1)
		}
	}

	return ""
}
