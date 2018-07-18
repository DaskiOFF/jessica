package utils

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func AskQuestion(question string, answerIsRequired bool) (answer string) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Printf("%s ", question)
	for {
		answer, _ = reader.ReadString('\n')
		answer = strings.TrimSpace(answer)

		if !answerIsRequired || len(answer) > 0 {
			return
		}
	}
}

func AskQuestionWithAnswers(question string, answers []string) (answer string) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Printf("%s ", question)
	for {
		answer, _ = reader.ReadString('\n')
		answer = strings.TrimSpace(answer)

		if len(strings.TrimSpace(answer)) > 0 && sliceContains(answers, answer) {
			return
		}
	}
}
