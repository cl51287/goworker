package goworker

type Producer interface {
	Produce() []*Task
}
