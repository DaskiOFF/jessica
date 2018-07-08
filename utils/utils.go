package utils

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func IsFileExist(filename string) bool {
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		return false
	}

	return true
}

func WriteToFile(fileName string, text string) {
	d1 := []byte(text)
	err := ioutil.WriteFile(fileName, d1, os.ModePerm)
	if err != nil {
		fmt.Println("Error create file with name: fileName")
	}
}

func PrintlnSuccessMessage(message string) {
	fmt.Println(message + "  🎉")
}

func PrintlnErrorMessage(message string) {
	fmt.Println("❌  " + message + "  ❌")
}

func PrintlnAttentionMessage(message string) {
	fmt.Println("🔶  " + message + "  🔶")
}

// FixBackQuotes Replace \%\*\*\*\% to \`\`\` and \%\*\% to \`
func FixBackQuotes(content string) string {
	content = strings.Replace(content, "%***%", "```", -1)
	content = strings.Replace(content, "%*%", "`", -1)

	return content
}

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
