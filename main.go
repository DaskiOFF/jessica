package main

import (
	"os"

	"github.com/daskioff/jessica/router"
	"github.com/daskioff/jessica/utils/print"
)

func printHelp() {
	print.PrintlnInfoMessage(`
Jessica – персональный ассистент для вашего проекта

Использование:

	jessica <command> [action] [arguments]

Список команд:

	version     Выводит текущую версию
	setup       Конфигурация проекта
	readme      Создание необходимых файлов и шаблонов для генерации "README.md" файла
	struct      Создание и описание структуры проекта
	generator   Создание и описание структуры проекта

Используйте "jessica help <command>" для подробной информации о команде.
`)
}

func main() {
	argsWithProg := os.Args
	if len(argsWithProg) == 1 {
		printHelp()
		return
	}

	arguments := os.Args[1:]

	router := router.New()
	err := router.Handle(arguments)
	if err != nil {
		print.PrintlnErrorMessage(err.Error())
	}
}
