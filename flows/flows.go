package flows

// Flow Интерфейс любого flow, который может запустить router
type Flow interface {
	Start(args []string)

	Setup()

	Description() string
}
