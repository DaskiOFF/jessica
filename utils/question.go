package utils

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strconv"
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

		if len(answer) > 0 && sliceContains(answers, answer) {
			return
		}
	}
}

func AskQuestionWithBoolAnswer(question string) bool {
	reader := bufio.NewReader(os.Stdin)

	positiveAnswers := []string{"y", "yes", "true"}
	negativeAnswers := []string{"n", "no", "false"}

	question = question + "(" + strings.Join(positiveAnswers, "/") + "|" + strings.Join(negativeAnswers, "/") + ")"

	fmt.Printf("%s ", question)
	for {
		answer, _ := reader.ReadString('\n')
		answer = strings.TrimSpace(answer)
		answer = strings.ToLower(answer)

		if sliceContains(positiveAnswers, answer) {
			return true
		} else if sliceContains(negativeAnswers, answer) {
			return false
		}
	}
}

func AskQuestionWithChooseFileAnswer(question string, fileExt string) string {
	fmt.Println("question")
	files, err := ioutil.ReadDir("./")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(files)
	xcodeprojFiles := []string{}
	for _, f := range files {
		extension := filepath.Ext(f.Name())
		fmt.Println(f.Name() + "   " + extension)
		if extension == fileExt {
			xcodeprojFiles = append(xcodeprojFiles, f.Name())
		}
	}

	count := len(xcodeprojFiles)
	if count == 0 {
		return ""
	}
	if count == 1 {
		return xcodeprojFiles[0]
	}

	answers := []string{}
	variants := []string{}
	for i, fileName := range xcodeprojFiles {
		answers = append(answers, string(i))
		variants = append(variants, string(i)+". "+fileName)
	}

	question = question + "\n" + strings.Join(variants, "\n")
	answer := AskQuestionWithAnswers(question, answers)

	index, err := strconv.Atoi(answer)
	if err != nil {
		index = 0
	}

	return variants[index]
}

func AskQuestionWithChooseFolderAnswer(question string) string {
	files, err := ioutil.ReadDir("./")
	if err != nil {
		log.Fatal(err)
	}

	folders := []string{"."}
	for _, f := range files {
		if f.IsDir() {
			folders = append(folders, f.Name())
		}
	}

	count := len(folders)
	if count == 1 {
		return folders[0]
	}

	answers := []string{}
	variants := []string{}
	for i, folderName := range folders {
		answers = append(answers, string(i))
		variants = append(variants, string(i)+". "+folderName)
	}

	question = question + "\n" + strings.Join(variants, "\n")
	answer := AskQuestionWithAnswers(question, answers)

	index, err := strconv.Atoi(answer)
	if err != nil {
		index = 0
	}

	return variants[index]
}
