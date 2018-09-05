package gentemplate

import (
	"strconv"
)

type Variable struct {
	Key   string
	Value string
}

func parseVariables(data interface{}) (map[string]interface{}, error) {
	switch data.(type) {
	case []interface{}:
		return parseVariablesArray(data.([]interface{}))
	case map[string]interface{}:
		return parseVariablesMap(data.(map[string]interface{}))
	}

	return map[string]interface{}{}, nil
}

func parseVariablesArray(data []interface{}) (map[string]interface{}, error) {
	if data == nil || len(data) == 0 {
		return map[string]interface{}{}, nil
	}

	variables := map[string]interface{}{}
	for i, value := range data {
		if value != nil {
			v := value.(string)
			variables["_"+strconv.Itoa(i)] = v
		}
	}

	return variables, nil
}

func parseVariablesMap(data map[string]interface{}) (map[string]interface{}, error) {
	variables := map[string]interface{}{}
	for k, v := range data {
		if v != nil {
			variables[k] = v.(string)
		}
	}

	return variables, nil
}
