package gentemplate

import (
	"testing"
)

func Test_ParseVariablesLikeMap(t *testing.T) {
	anyKey := "anyKey"
	anyString := "anyString"

	data := []struct {
		data     map[interface{}]interface{}
		variable *Variable
		err      error
	}{
		{
			data: map[interface{}]interface{}{
				anyKey: anyString,
			},
			variable: &Variable{
				Key:   anyKey,
				Value: anyString,
			},
			err: nil,
		},
		{
			data: map[interface{}]interface{}{
				anyKey: anyString + anyString,
			},
			variable: &Variable{
				Key:   anyKey,
				Value: anyString + anyString,
			},
			err: nil,
		},
		{
			data:     map[interface{}]interface{}{},
			variable: nil,
			err:      nil,
		},
	}

	for _, d := range data {
		result, err := parseVariables(d.data)

		if d.err == nil && err != nil || d.err != nil && err == nil {
			t.Error("Expected err == d.err, got ", err.Error())
		}

		if d.variable == nil && len(result) > 0 {
			t.Error("Expected result is empty, got ", len(result))
			return
		}

		if len(result) == 0 {
			if d.variable != nil {
				t.Error("Expected result is not empty, got ", len(result))
			}
			return
		}

		if result[0].Key != d.variable.Key || result[0].Value != d.variable.Value {
			t.Errorf("Expected data %v == %v, got %v", d.data, d.variable, result)
		}
	}
}

func Test_ParseVariablesLikeSlice(t *testing.T) {
	anyString := "anyString"

	data := []struct {
		data     []interface{}
		variable *Variable
		err      error
	}{
		{
			data: []interface{}{
				anyString,
			},
			variable: &Variable{
				Key:   "_0",
				Value: anyString,
			},
			err: nil,
		},
		{
			data: []interface{}{
				anyString + anyString,
			},
			variable: &Variable{
				Key:   "_0",
				Value: anyString + anyString,
			},
			err: nil,
		},
		{
			data:     []interface{}{},
			variable: nil,
			err:      nil,
		},
	}

	for _, d := range data {
		result, err := parseVariables(d.data)

		if d.err == nil && err != nil || d.err != nil && err == nil {
			t.Error("Expected err == d.err, got ", err.Error())
		}

		if d.variable == nil && len(result) > 0 {
			t.Error("Expected result is empty, got ", len(result))
			return
		}

		if len(result) == 0 {
			if d.variable != nil {
				t.Error("Expected result is not empty, got ", len(result))
			}
			return
		}

		if result[0].Key != d.variable.Key || result[0].Value != d.variable.Value {
			t.Errorf("Expected data %v == %v, got %v", d.data, d.variable, result)
		}
	}
}

func Test_ParseVariables(t *testing.T) {
	result, err := parseVariables("")
	if len(result) != 0 && err != nil {
		t.Error("Expected err == d.err, got ", err.Error())
	}

	result, err = parseVariables(1)
	if len(result) != 0 && err != nil {
		t.Error("Expected err == d.err, got ", err.Error())
	}
}
