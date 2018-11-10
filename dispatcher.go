package goworker

type Dispatcher interface {
	Dispatch([]Worker) Worker
}
