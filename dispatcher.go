package goworker

type Dispatcher interface {
	dispatch([]Worker) Worker
}
