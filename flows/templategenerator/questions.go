package templategenerator

import "github.com/daskioff/jessica/utils"

type question struct {
	key        string
	text       string
	isRequired bool
}

func newQuestions(in []interface{}) []question {
	questions := []question{}

	for _, questionInterface := range in {
		questionMap := questionInterface.(map[interface{}]interface{})
		question := question{}

		question.key = questionMap["key"].(string)
		question.text = questionMap["text"].(string)

		isRequired := questionMap["required"]
		if isRequired == nil {
			question.isRequired = false
		} else {
			question.isRequired = isRequired.(bool)
		}

		questions = append(questions, question)
	}

	return questions
}

func askQuestions(questions []question) map[string]interface{} {
	answers := make(map[string]interface{}, 0)

	for _, question := range questions {
		answer := utils.AskQuestion(question.text, question.isRequired)
		answers[question.key] = answer
	}

	return answers
}
