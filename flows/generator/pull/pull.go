package pull

import (
	"strings"

	"github.com/daskioff/jessica/utils/git"
	"github.com/daskioff/jessica/utils/print"
)

func Execute(args []string, absTemplatesFolderPath string) {
	if len(args) == 0 {
		print.PrintlnErrorMessage("Вы не указали URL git репозитория")
		return
	}

	url := args[0]
	args = args[1:]

	if !strings.HasPrefix(url, "http") {
		url = "https://" + url
	}

	if !strings.HasSuffix(url, ".git") {
		url = url + ".git"
	}

	branch := ""
	if len(args) > 0 {
		branch = args[0]
	}

	path := absTemplatesFolderPath
	err := git.Clone(url, branch, path)
	if err != nil {
		panic(err)
	}
}
