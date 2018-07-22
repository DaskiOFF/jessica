package templategenerator

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/daskioff/jessica/flows/projectstruct"
)

func searchTemplates() map[string][]string {
	templatesRoot, err := os.Getwd()
	if err != nil {
		fmt.Println(err.Error())
		return make(map[string][]string)
	}
	templatesRoot = filepath.Join(templatesRoot, projectstruct.TemplatesFolderName)

	files := make(map[string][]string)
	filepath.Walk(templatesRoot, func(path string, f os.FileInfo, _ error) error {
		if !f.IsDir() && filepath.Ext(path) == TemplateExt {
			path = strings.Replace(path, templatesRoot, "", -1)

			key := strings.TrimSuffix(f.Name(), TemplateExt)
			if files[key] == nil {
				files[key] = []string{path}
			} else {
				files[key] = append(files[key], path)
			}
		}
		return nil
	})

	for k, v := range files {
		if len(v) > 1 {
			files[k], _ = fixPaths(v)
		}
	}

	return files
}

func slitPath(path string) []string {
	return strings.Split(path, string(os.PathSeparator))
}

func fixPaths(paths []string) ([]string, bool) {
	if len(paths) < 2 {
		return paths, false
	}

	clearAll := true
	partsOfFirstPath := slitPath(paths[0])
	for i, path := range paths {
		if i == 0 {
			continue
		}

		partsOfCurrentPath := slitPath(path)
		if partsOfCurrentPath[0] != partsOfFirstPath[0] {
			clearAll = false
			break
		}
	}

	if clearAll {
		for i, path := range paths {
			partsOfCurrentPath := slitPath(path)
			paths[i] = filepath.Join(partsOfCurrentPath[1:]...)
		}

		paths, clearAll = fixPaths(paths)
	}

	return paths, clearAll
}
