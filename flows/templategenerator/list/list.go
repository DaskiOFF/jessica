package list

import "github.com/daskioff/jessica/utils/print"

func Show(templates []string) {
	if len(templates) == 0 {
		print.PrintlnAttentionMessage("Шаблоны не найдены")
		return
	}

	list := ""
	for _, template := range templates {
		if len(list) == 0 {
			list = template
		} else {
			list = list + "\n" + template
		}
	}
	print.PrintlnInfoMessage(list)
}
