package gentemplate

import (
	"errors"
	"testing"
)

func Test_ParseQuestion(t *testing.T) {
	anyString := "anyString"

	data := []struct {
		data     map[interface{}]interface{}
		question *Question
		err      error
	}{
		{
			data: map[interface{}]interface{}{
				"key":      anyString,
				"text":     anyString,
				"required": false,
			},
			question: &Question{
				Key:        anyString,
				Text:       anyString,
				IsRequired: false,
			},
			err: nil,
		},
		{
			data: map[interface{}]interface{}{
				"key":      anyString,
				"text":     anyString,
				"required": true,
			},
			question: &Question{
				Key:        anyString,
				Text:       anyString,
				IsRequired: true,
			},
			err: nil,
		},
		{
			data: map[interface{}]interface{}{
				"key":  anyString,
				"text": anyString,
			},
			question: &Question{
				Key:        anyString,
				Text:       anyString,
				IsRequired: false,
			},
			err: nil,
		},
		{
			data: map[interface{}]interface{}{
				"text": anyString,
			},
			question: nil,
			err:      errors.New("Поле key у вопроса не может быть пустым"),
		},
		{
			data:     map[interface{}]interface{}{},
			question: nil,
			err:      errors.New("Поле key у вопроса не может быть пустым"),
		},
		{
			data: map[interface{}]interface{}{
				"key": anyString,
			},
			question: nil,
			err:      errors.New("Поле text у вопроса не может быть пустым"),
		},
	}

	for _, d := range data {
		result, err := parseQuestion(d.data)

		if d.err == nil && err != nil || d.err != nil && err == nil {
			t.Error("Expected err == d.err, got ", err.Error())
		}

		if d.question == nil {
			if result != nil {
				t.Errorf("Expected %v == nil", result)
			}
			return
		}
		if result.Key != d.question.Key || result.Text != d.question.Text || result.IsRequired != d.question.IsRequired {
			t.Errorf("Expected data %v == %v, got %v", d.data, d.question, result)
		}
	}
}

func Test_ParseQuestions(t *testing.T) {
	anyString := "anyString"

	data := []struct {
		data     map[interface{}]interface{}
		question *Question
		err      error
	}{
		{
			data: map[interface{}]interface{}{
				"key":      anyString,
				"text":     anyString,
				"required": false,
			},
			question: &Question{
				Key:        anyString,
				Text:       anyString,
				IsRequired: false,
			},
			err: nil,
		},
		{
			data: map[interface{}]interface{}{
				"key":      anyString,
				"text":     anyString,
				"required": true,
			},
			question: &Question{
				Key:        anyString,
				Text:       anyString,
				IsRequired: true,
			},
			err: nil,
		},
		{
			data: map[interface{}]interface{}{
				"key":  anyString,
				"text": anyString,
			},
			question: &Question{
				Key:        anyString,
				Text:       anyString,
				IsRequired: false,
			},
			err: nil,
		},
		{
			data: map[interface{}]interface{}{
				"text": anyString,
			},
			question: nil,
			err:      errors.New("Поле key у вопроса не может быть пустым"),
		},
		{
			data:     map[interface{}]interface{}{},
			question: nil,
			err:      errors.New("Поле key у вопроса не может быть пустым"),
		},
		{
			data: map[interface{}]interface{}{
				"key": anyString,
			},
			question: nil,
			err:      errors.New("Поле text у вопроса не может быть пустым"),
		},
	}

	result, err := parseQuestions([]interface{}{data[0].data, data[1].data})
	if len(result) != 2 || err != nil {
		t.Error("Expected length result == 2, error == nil, got, ", len(result), err)
	}

	result, err = parseQuestions([]interface{}{data[0].data, data[3].data})
	if len(result) != 0 || err == nil || (err != nil && err.Error() != data[3].err.Error()) {
		t.Error("Expected result == nil, error == error key, got, ", result, err)
	}

	result, err = parseQuestions([]interface{}{data[0].data, data[5].data})
	if len(result) != 0 || err == nil || (err != nil && err.Error() != data[5].err.Error()) {
		t.Error("Expected result == nil, error == error text, got, ", result, err)
	}
}
