package goworker

type WorkerManager struct {
	workers          []Worker
	workerTaskQueues []chan *Task
	taskQueue        chan *Task
	dispatcher       Dispatcher
	taskQueueLen     int
}

func NewWorkerManager(dispatcer Dispatcher, workers []Worker) *WorkerManager {
	return &WorkerManager{
		dispatcher:   dispatcer,
		workers:      workers,
		taskQueueLen: 100,
	}
}

func (wm *WorkerManager) AddWorker(workerName string, worker Worker) bool {

}

func (wm *WorkerManager) HasWorker(workerName string) bool {

}

func (wm *WorkerManager) DelWorker(workerName string) bool {

}

func (workerManager *WorkerManager) SetAgent(agent Agent) {
}

func (workerManager *WorkerManager) Run() {
	if workerManager.dispatcher == nil {
		panic("dispatcher not set")
	}

	workerManager.taskQueue = make(chan *Task, workerManager.taskQueueLen)

	for true {
		task := <-workerManager.taskQueue
		worker := workerManager.dispatcher.Dispatch(workerManager.workers)
		go worker.Handle(task)
	}
}
