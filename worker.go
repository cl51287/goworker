package goworker

// Worker is a interface
type Worker interface {
	Handle(task *Task) error
}
