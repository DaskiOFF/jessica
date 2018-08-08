package generator

import "github.com/daskioff/jessica/utils/question"

type Quest struct {
	key        string
	text       string
	isRequired bool
}

func NewQuestions(in []interface{}) []Quest {
	questions := []Quest{}

	for _, questionInterface := range in {
		questionMap := questionInterface.(map[interface{}]interface{})
		Quest := Quest{}

		Quest.key = questionMap["key"].(string)
		Quest.text = questionMap["text"].(string)

		isRequired := questionMap["required"]
		if isRequired == nil {
			Quest.isRequired = false
		} else {
			Quest.isRequired = isRequired.(bool)
		}

		questions = append(questions, Quest)
	}

	return questions
}

func AskQuestions(questions []Quest) map[string]interface{} {
	answers := make(map[string]interface{}, 0)

	for _, Quest := range questions {
		answer := question.AskQuestion(Quest.text, Quest.isRequired)
		answers[Quest.key] = answer
	}

	return answers
}
