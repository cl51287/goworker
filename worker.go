package goworker

// Worker is a interface
type Worker interface {
	AddTask(*Task)

	Stop()

	Start()
}
