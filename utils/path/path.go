package path

import (
	"os"
	"path/filepath"
)

// ProjectRoot возвращает абсолютный путь до места запуска программы
func ProjectRoot() (string, error) {
	getwd, err := os.Getwd()
	if err != nil {
		return "", err
	}

	return getwd, nil
}

// InProjectRoot вовращает абсолютный путь составленный из компонентов добавленных к ProjectRoot
func InProjectRoot(elem ...string) (string, error) {
	root, err := ProjectRoot()
	if err != nil {
		return root, err
	}

	components := []string{root}
	components = append(components, elem...)
	path := filepath.Join(components...)

	return path, nil
}
