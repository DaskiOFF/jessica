package gentemplate

import "strconv"

type Variable struct {
	Key   string
	Value string
}

func parseVariables(data interface{}) ([]Variable, error) {
	switch data.(type) {
	case []interface{}:
		return parseVariablesArray(data.([]interface{}))
	case map[interface{}]interface{}:
		return parseVariablesMap(data.(map[interface{}]interface{}))
	}

	return []Variable{}, nil
}

func parseVariablesArray(data []interface{}) ([]Variable, error) {
	if data == nil || len(data) == 0 {
		return []Variable{}, nil
	}

	variables := []Variable{}
	for i, value := range data {
		if value != nil {
			v := value.(string)
			variables = append(variables, Variable{
				Key:   "_" + strconv.Itoa(i),
				Value: v,
			})
		}
	}

	return variables, nil
}

func parseVariablesMap(data map[interface{}]interface{}) ([]Variable, error) {
	variables := []Variable{}
	for k, v := range data {
		if k != nil && v != nil {
			variables = append(variables, Variable{
				Key:   k.(string),
				Value: v.(string),
			})
		}
	}

	return variables, nil
}
