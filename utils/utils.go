package utils

import (
	"io/ioutil"
	"strings"
)

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
