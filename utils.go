package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func isFileExist(filename string) bool {
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		return false
	}

	return true
}

func writeToFile(fileName string, text string) {
	d1 := []byte(text)
	err := ioutil.WriteFile(fileName, d1, os.ModePerm)
	if err != nil {
		fmt.Println("Error create file with name: fileName")
	}
}

func printlnSuccessMessage(message string) {
	fmt.Println(message + "  🎉")
}

func printlnErrorMessage(message string) {
	fmt.Println("❌  " + message + "  ❌")
}

func printlnAttentionMessage(message string) {
	fmt.Println("🔶  " + message + "  🔶")
}
