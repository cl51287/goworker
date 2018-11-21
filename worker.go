package goworker

// Worker is a interface
type Worker interface {
	SetAgent(Agent)

	Stop()

	Start()
}
