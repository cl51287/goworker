package goworker

type Producer interface {
	SetAgent(Agent)

	Stop()

	Start()
}
