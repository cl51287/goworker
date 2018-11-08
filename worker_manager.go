package goworker

type WorkerManager struct {
	workers []*Worker
}

func (workerManager *WorkerManager) AddTask(task *Task) error {
	return nil
}

func (WorkerManager *WorkerManager) StartWorking() {
}
