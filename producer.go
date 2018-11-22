package goworker

type Producer interface {
	SetProducerManager(*ProducerManager)

	Start() bool

	Stop() bool
}
