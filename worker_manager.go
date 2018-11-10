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

func (workerManager *WorkerManager) SetTaskQueueLen(len int) {
	workerManager.taskQueueLen = len
}

func (workerManager *WorkerManager) AddTask(task *Task) {
	workerManager.taskQueue <- task
}

func (workerManager *WorkerManager) SetDispatcher(dispatcher Dispatcher) {
	workerManager.dispatcher = dispatcher
}

func (workerManager *WorkerManager) StartWorking() {
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
