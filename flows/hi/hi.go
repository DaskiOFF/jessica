package hi

import (
	"fmt"

	"github.com/daskioff/jessica/flows"
)

type HiFlow struct {
	version string
}

func (flow *HiFlow) Start(args []string) {
	fmt.Println("Привет, моя версия " + flow.version)
}

func (flow *HiFlow) Description() string {
	return "Тестовый метод проверки работоспособности"
}

// ----------------------------------------------------------------------------
func NewFlow(version string) flows.Flow {
	flow := HiFlow{version}
	return &flow
}
