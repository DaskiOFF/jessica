package projectstruct

import (
	"path/filepath"

	"github.com/daskioff/jessica/configs"
)

func projectPaths() []string {
	projectStructure := configs.ProjectConfig.Get(configs.KeyCustomProjectStruct)
	return projectStructToPaths(projectStructure)
}

func projectStructToPaths(in interface{}) []string {
	switch v := in.(type) {

	case string:
		return []string{v}

	case map[interface{}]interface{}:
		response := make([]string, 0)
		for s, b := range v {
			prefix, _ := s.(string)
			for _, path := range projectStructToPaths(b) {
				response = append(response, filepath.Join(prefix, path))
			}
		}
		return response

	case []string:
		return v

	case []interface{}:
		response := make([]string, 0)
		for _, b := range v {
			for _, path := range projectStructToPaths(b) {
				response = append(response, path)
			}
		}
		return response

	default:
		return make([]string, 0)
	}
}
