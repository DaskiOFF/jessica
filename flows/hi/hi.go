package hi

import (
	"fmt"

	"github.com/daskioff/jessica/flows"
)

type HiFlow struct {
}

func (flow *HiFlow) Start(args []string) {
	fmt.Println("Привет, моя версия 1.3.2")
}

func (flow *HiFlow) Setup() {}

func (flow *HiFlow) Description() string {
	return "Тестовый метод проверки работоспособности"
}

// ----------------------------------------------------------------------------
func NewFlow() flows.Flow {
	flow := HiFlow{}
	return &flow
}
