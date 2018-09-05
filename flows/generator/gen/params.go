package gen

import "strings"

// Params описывает аргументы переданные с действием генератора
type Params struct {
	TemplateName      string
	ModuleName        string
	NeedGenerateTests bool
	NeedGenerateMock  bool
	CustomKeys        map[string]interface{}
}

// NewParams возвращает объект Params заполненный данными из `args`
func NewParams(args []string) Params {
	p := Params{}

	p.TemplateName = args[0]

	p.ModuleName = ""
	splitPosition := 1
	if len(args) > 1 && !strings.HasPrefix(args[1], "--") && !strings.Contains(args[1], ":") {
		p.ModuleName = args[1]
		splitPosition = 2
	}

	p.NeedGenerateTests = true
	p.NeedGenerateMock = true
	p.CustomKeys = make(map[string]interface{})
	for _, arg := range args[splitPosition:] {
		if arg == "--notest" {
			p.NeedGenerateTests = false
		}
		if arg == "--nomock" {
			p.NeedGenerateMock = false
		}

		// custom keys
		splitResult := strings.Split(arg, ":")
		if len(splitResult) == 2 {
			k := strings.TrimSpace(splitResult[0])
			v := strings.TrimSpace(splitResult[1])
			p.CustomKeys[k] = v
		}
	}

	return p
}
