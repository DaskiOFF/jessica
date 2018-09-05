package gentemplate

import (
	"testing"
)

func Test_ParseVariablesLikeMap(t *testing.T) {
	anyKey := "anyKey"
	anyString := "anyString"

	data := []struct {
		data     map[string]interface{}
		variable map[string]interface{}
		err      error
	}{
		{
			data: map[string]interface{}{
				anyKey: anyString,
			},
			variable: map[string]interface{}{
				anyKey: anyString,
			},
			err: nil,
		},
		{
			data: map[string]interface{}{
				anyKey: anyString + anyString,
			},
			variable: map[string]interface{}{
				anyKey: anyString + anyString,
			},
			err: nil,
		},
		{
			data:     map[string]interface{}{},
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
			if d.variable != nil && len(d.variable) != 0 {
				t.Error("Expected result is not empty, got ", d.variable)
			}
			return
		}

		for k, v := range d.variable {
			if result[k] != v {
				t.Errorf("Expected data %v == %v, got %v", d.data, d.variable, result)
			}
		}
	}
}

func Test_ParseVariablesLikeSlice(t *testing.T) {
	anyString := "anyString"

	data := []struct {
		data     []interface{}
		variable map[string]interface{}
		err      error
	}{
		{
			data: []interface{}{
				anyString,
			},
			variable: map[string]interface{}{
				"_0": anyString,
			},
			err: nil,
		},
		{
			data: []interface{}{
				anyString + anyString,
			},
			variable: map[string]interface{}{
				"_0": anyString + anyString,
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

		for k, v := range d.variable {
			if result[k] != v {
				t.Errorf("Expected data %v == %v, got %v", d.data, d.variable, result)
			}
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
