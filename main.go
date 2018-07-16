package main

import (
	"os"

	"github.com/daskioff/jessica/router"
	"github.com/daskioff/jessica/utils"
)

func main() {
	argsWithProg := os.Args
	if len(argsWithProg) == 1 {
		utils.PrintlnErrorMessage("Вы так и не сказали что делать")
		return
	}

	argsWithoutProg := os.Args[1:]

	router := router.NewRouter()
	err := router.Handle(argsWithoutProg)
	if err != nil {
		utils.PrintlnErrorMessage(err.Error())
	}
}
