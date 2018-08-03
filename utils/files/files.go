package files

import (
	"fmt"
	"io/ioutil"
	"os"
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

func Folders(root string) []string {
	files, err := ioutil.ReadDir(root)
	if err != nil {
		fmt.Println(err)
		return []string{}
	}

	folders := make([]string, 0)
	for _, file := range files {
		if file.IsDir() {
			folders = append(folders, file.Name())
		}
	}

	return folders
}
