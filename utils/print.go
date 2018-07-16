package utils

import "fmt"

// PrintlnInfoMessage Выводит информационное сообщение в лог
func PrintlnInfoMessage(message string) {
	fmt.Println(message)
}

// PrintlnSuccessMessage Выводит сообщение успеха в лог
func PrintlnSuccessMessage(message string) {
	fmt.Println(message + "  🎉")
}

// PrintlnErrorMessage Выводит сообщение ошибки в лог
func PrintlnErrorMessage(message string) {
	fmt.Println("❌  " + message + "  ❌")
}

// PrintlnAttentionMessage Выводит сообщение заслуживающее внимания в лог
func PrintlnAttentionMessage(message string) {
	fmt.Println("🔶  " + message + "  🔶")
}
