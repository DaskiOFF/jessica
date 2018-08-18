package hi

import (
	"fmt"
)

type HiFlow struct {
	version string
}

func (flow *HiFlow) Start(args []string) {
	fmt.Println("Привет, моя версия " + flow.version)
}

func (flow *HiFlow) Description() string {
	return "Тестовый метод проверки работоспособности"
	return `--------------------------------------------------------------------------------
Приветствие и вывод своей текущей версии
--------------------------------------------------------------------------------`
}

// ----------------------------------------------------------------------------
func New(version string) *HiFlow {
	flow := HiFlow{version}
	return &flow
}
