package templategenerator

import "github.com/daskioff/jessica/utils/question"

type quest struct {
	key        string
	text       string
	isRequired bool
}

func (flow *TemplateGeneratorFlow) newQuestions(in []interface{}) []quest {
	questions := []quest{}

	for _, questionInterface := range in {
		questionMap := questionInterface.(map[interface{}]interface{})
		quest := quest{}

		quest.key = questionMap["key"].(string)
		quest.text = questionMap["text"].(string)

		isRequired := questionMap["required"]
		if isRequired == nil {
			quest.isRequired = false
		} else {
			quest.isRequired = isRequired.(bool)
		}

		questions = append(questions, quest)
	}

	return questions
}

func (flow *TemplateGeneratorFlow) askQuestions(questions []quest) map[string]interface{} {
	answers := make(map[string]interface{}, 0)

	for _, quest := range questions {
		answer := question.AskQuestion(quest.text, quest.isRequired)
		answers[quest.key] = answer
	}

	return answers
}
