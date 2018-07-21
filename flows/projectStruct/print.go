package projectstruct

import "strings"

func projectStructToString(in interface{}, currentSpace string, spaceStep string) string {
	switch v := in.(type) {

	case string:
		return currentSpace + "- " + v

	case map[interface{}]interface{}:
		response := ""
		for s, b := range v {
			prefix, _ := s.(string)
			response = response + currentSpace + "- " + prefix + "\n"
			response = response + projectStructToString(b, currentSpace+spaceStep, spaceStep)
		}
		return strings.TrimSuffix(response, "\n")

	case []string:
		response := ""
		for _, v := range v {
			response = response + currentSpace + "- " + v + "\n"
		}
		return strings.TrimSuffix(response, "\n")

	case []interface{}:
		response := ""
		for _, b := range v {
			response = response + projectStructToString(b, currentSpace, spaceStep) + "\n"
		}
		return response

	default:
		return ""
	}
}
