package gentemplate

import "errors"

var (
	ErrQuestionNeedKey  = errors.New("Поле key у вопроса не может быть пустым")
	ErrQuestionNeedText = errors.New("Поле text у вопроса не может быть пустым")
)

// Question описывает вопрос для шаблона
type Question struct {
	Key        string
	Text       string
	IsRequired bool
}

func parseQuestions(data []interface{}) ([]Question, error) {
	if data == nil || len(data) == 0 {
		return []Question{}, nil
	}

	questions := []Question{}
	for _, questionData := range data {
		q, err := parseQuestion(questionData.(map[interface{}]interface{}))

		if err != nil {
			return nil, err
		}
		questions = append(questions, *q)
	}

	return questions, nil
}

func parseQuestion(data map[interface{}]interface{}) (*Question, error) {
	q := Question{
		Key:        "",
		Text:       "",
		IsRequired: false,
	}

	key := data["key"]
	if key == nil {
		return nil, ErrQuestionNeedKey
	}
	q.Key = key.(string)

	text := data["text"]
	if text == nil {
		return nil, ErrQuestionNeedText
	}
	q.Text = text.(string)

	isRequired := data["required"]
	if isRequired != nil {
		q.IsRequired = isRequired.(bool)
	}

	return &q, nil
}
